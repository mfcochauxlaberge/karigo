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
		// funcs: map[string]Action{},

		err:      make(chan error),
		shutdown: make(chan bool),

		logger: zerolog.Logger{},
	}

	return node
}

// Node ...
type Node struct {
	Name string

	// Run
	log  Journal
	main source

	// Schema
	schema *jsonapi.Schema
	// funcs  map[string]Action

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
				Info().
				Str("err", err.Error()).
				Msg("Panic")
		}
	}()

	if len(r.Body) > 0 {
		r.Doc, err = jsonapi.UnmarshalDocument(r.Body, n.schema)
		if err != nil {
			r.Logger.
				Debug().
				Str("error", err.Error()).
				Msg("Could not unmarshal document")
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
					Debug().
					Str("error", err.Error()).
					Msg("Could not partially unmarshal resource")
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

	tx, _ := n.main.src.NewTx()

	cp := &Checkpoint{
		tx:   tx,
		node: n,
		ops:  []Op{},
	}

	// Check password is correct if request is writing (non-GET).
	if r.Method == POST || r.Method == PATCH || r.Method == DELETE {
		pwRes, _ := cp.tx.Resource(QueryRes{
			Set:    "0_meta",
			ID:     "password",
			Fields: []string{"value"},
		})
		if pwRes != nil {
			if hash, _ := pwRes.Get("value").(string); hash != "" {
				if r.Doc.Meta == nil {
					// Temporary. Maybe the Meta field should
					// never be nil?
					r.Doc.Meta = map[string]interface{}{}
				}

				pw, _ := r.Doc.Meta["password"].(string)

				err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
				if err != nil {
					// jaerr := jsonapi.NewErrForbidden()
					// doc.Data = jaerr
					doc.Errors = []jsonapi.Error{jsonapi.NewErrForbidden()}

					return doc
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

	execute := ActionDefault
	ops := []Op{}
	// Prepare action
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

		found, _ := cp.tx.Resource(QueryRes{
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
		execute(cp)
	}

	for _, op := range cp.ops {
		r.Logger.Debug().
			Str("op", op.String()).
			Msg("Operation")
	}

	if cp.err != nil {
		r.Logger.
			Debug().
			Str("error", cp.err.Error()).
			Msg("Action failed")

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
		// Commit the entry
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

// AddSource adds a source to the node.
//
// name is currently ignored and only one source can be added. Adding a second
// source simply overrides the first one.
//
// TODO Add support for multiple sources.
func (n *Node) AddSource(name string, s Source) {
	n.main = source{
		src: s,
	}
}

// RegisterJournal ...
func (n *Node) RegisterJournal(j Journal) {
	n.log = j
}
