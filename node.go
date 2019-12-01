package karigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

// NewNode ...
func NewNode(journal Journal, src Source) *Node {
	node := &Node{
		log: journal,
		main: source{
			src: src,
		},

		schema: FirstSchema(),
		// funcs: map[string]Tx{},

		err:      make(chan error),
		shutdown: make(chan bool),

		logger: zerolog.Logger{},
	}

	return node
}

// Node ...
type Node struct {
	Name    string
	Domains []string

	// Run
	log  Journal
	main source

	// Schema
	schema *jsonapi.Schema
	// funcs  map[string]Tx

	// Channels
	err      chan error
	shutdown chan bool

	// Internal
	logger zerolog.Logger
	sync.Mutex
}

// Run ...
func (n *Node) Run() error {
	// n.Lock()
	// n.Unlock()
	// Handle events
	for {
		select {}
	}
}

// Handle ...
func (n *Node) Handle(r *Request) *jsonapi.Document {
	var (
		res jsonapi.Resource
		doc = &jsonapi.Document{}
		err error
	)

	defer func() {
		if p := recover(); p != nil {
			var err error

			switch e := p.(type) {
			case string:
				err = errors.New(e)
			case jsonapi.Error:
				err = e
			case error:
				err = e
			}

			r.Logger.
				Err(err).
				Str("abc", "123").
				Int("def", 42).
				Msg("Panic")
		}
	}()

	if r.Method == POST || r.Method == PATCH {
		r.Doc, err = jsonapi.UnmarshalDocument(r.Body, n.schema)
		if err != nil {
			r.Logger.Err(err).Msg("Could not unmarshal document")
		}
	}

	if r.Method == PATCH {
		frame := struct {
			Data json.RawMessage
		}{}

		err = json.Unmarshal(r.Body, &frame)
		if err == nil {
			res, err = jsonapi.UnmarshalPartialResource(frame.Data, n.schema)
			if err != nil {
				r.Logger.
					Err(err).
					Msg("Could not partially unmarshal reosurce")
			}
		}

		r.Doc.Data = res
	}

	if jaerr, ok := err.(jsonapi.Error); ok {
		doc.Data = jaerr
		return doc
	} else if err != nil {
		jaerr = jsonapi.NewErrInternalServerError()
		jaerr.Detail = err.Error()
		doc.Data = jaerr
		return doc
	}

	if r.Doc != nil && r.Doc.Data != nil {
		res, _ = r.Doc.Data.(jsonapi.Resource)
	}

	cp := &Checkpoint{
		node: n,
		ops:  []Op{},
	}

	// Check password is correct if request is writing (non-GET).
	if r.Method == POST || r.Method == PATCH || r.Method == DELETE {
		pwRes, _ := n.main.src.Resource(QueryRes{
			Set:    "0_sets",
			ID:     "password",
			Fields: []string{"value"},
		})
		if pwRes != nil {
			if hashed, _ := pwRes.Get("value").(string); hashed != "" {
				if pw, _ := r.Doc.Meta["password"].(string); pw != "" {
					err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pw))
					if err != nil {
						// jaerr := jsonapi.NewErrForbidden()
						// doc.Data = jaerr
						doc.Data = jsonapi.NewErrForbidden()
						return doc
					}
				}
			}
		}
	}

	// Hash password if it's being updated.
	if r.Method == POST || r.Method == PATCH {
		if r.URL.ResType == "0_meta" {
			if res, ok := r.Doc.Data.(jsonapi.Resource); ok {
				if pw, _ := res.Get("value").(string); pw != "" && res.GetID() == "password" {
					npw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
					if err != nil {
						jaerr := jsonapi.NewErrInternalServerError()
						jaerr.Detail = err.Error()
						doc.Data = jaerr

						return doc
					}

					res.Set("value", string(npw))
				}
			}
		}
	}

	tx := TxDefault
	ops := []Op{}
	// Prepare transaction
	switch r.Method {
	case GET:
	case POST:
		if res.GetID() == "" {
			res.SetID(uuid.New().String()[:8])
		}
		// TODO Do not hardcode the following conditions. It can
		// be handled in a much better way.
		switch res.GetType().Name {
		case "0_sets":
			res.SetID(res.Get("name").(string))
			ops = NewOpAddSet(res.GetID())
		case "0_attrs":
			ops = NewOpAddAttr(
				res.GetToOne("set"),
				res.Get("name").(string),
				res.Get("type").(string),
				res.Get("null").(bool),
			)
			res.SetID(ops[0].Value.(string))
		case "0_rels":
			ops = NewOpAddRel(
				res.GetToOne("from-set"),
				res.Get("from-name").(string),
				res.GetToOne("to-set"),
				res.Get("to-name").(string),
				res.Get("to-one").(bool),
				res.Get("from-one").(bool),
			)
			res.SetID(ops[0].Value.(string))
		default:
			ops = NewOpInsert(res)
		}

		found, _ := n.resource(0, QueryRes{
			Set:    res.GetType().Name,
			ID:     res.GetID(),
			Fields: []string{"id"},
		})

		if found != nil {
			cp.Fail(errors.New("id already used"))
		}
	case PATCH:
		ops = []Op{}

		for _, attr := range res.Attrs() {
			ops = append(ops, NewOpSet(
				r.URL.ResType,
				res.GetID(),
				attr.Name,
				res.Get(attr.Name),
			))
		}

		for _, rel := range res.Rels() {
			if rel.ToOne {
				ops = append(ops, NewOpSet(
					r.URL.ResType,
					res.GetID(),
					rel.FromName,
					res.GetToOne(rel.FromName),
				))
			} else {
				ops = append(ops, NewOpSet(
					r.URL.ResType,
					res.GetID(),
					rel.FromName,
					res.GetToMany(rel.FromName),
				))
			}
		}
	case DELETE:
		ops = []Op{NewOpSet(r.URL.ResType, r.URL.ResID, "id", "")}
	}

	cp.Apply(ops)

	if r.isSchemaChange() {
		// Handle schema change
		handleSchemaChange(n.schema, r, cp)
	} else {
		// Execute
		tx(cp)
	}

	for _, op := range ops {
		r.Logger.Debug().
			Str("op", op.String()).
			Msg("Operation")
	}

	if cp.err != nil {
		r.Logger.
			Err(cp.err).
			Msg("Transaction failed")

		// Rollback
		err = cp.rollback()
		if err != nil {
			panic(fmt.Errorf("could not rollback: %s", err))
		}

		// Handle error
		var jaErr jsonapi.Error

		switch cp.err {
		case ErrNotImplemented:
			jaErr = jsonapi.NewErrNotImplemented()
		default:
			jaErr = jsonapi.NewErrInternalServerError()
		}

		doc.Errors = []jsonapi.Error{jaErr}
	} else {
		// Commit the transaction entry
		err = n.log.Append(cp.ops.Bytes())
		if err != nil {
			panic(fmt.Errorf("could not append: %s", err))
		}

		err = cp.commit()
		if err != nil {
			panic(fmt.Errorf("could not commit: %s", err))
		}

		// Response payload
		switch r.Method {
		case GET:
			if !r.URL.IsCol {
				res := cp.Resource(NewQueryRes(r.URL))
				doc.Data = res
			} else {
				col := &jsonapi.SoftCollection{}
				typ := n.schema.GetType(r.URL.ResType)
				col.SetType(&typ)
				resources := cp.Collection(NewQueryCol(r.URL))
				for i := 0; i < resources.Len(); i++ {
					col.Add(resources.At(i))
				}
				doc.Data = jsonapi.Collection(col)
			}
		case POST, PATCH:
			qry := NewQueryRes(r.URL)
			qry.ID = res.GetID()
			res := cp.Resource(qry)
			doc.Data = res
		case DELETE:
		}
	}

	return doc
}

// resource ...
// TODO Validate the query?
func (n *Node) resource(_ uint, qry QueryRes) (jsonapi.Resource, error) {
	return n.main.src.Resource(qry)
}

// collection ...
// TODO Validate the query?
// TODO Complete the sorting rule
func (n *Node) collection(_ uint, qry QueryCol) (jsonapi.Collection, error) {
	return n.main.src.Collection(qry)
}

// apply ...
func (n *Node) apply(ops []Op) error {
	err := n.main.src.Apply(ops)
	if err != nil {
		return errors.New("karigo: an operation could not be executed")
	}

	return nil
}
