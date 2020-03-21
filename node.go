package karigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/mfcochauxlaberge/karigo/query"

	"github.com/google/uuid"
	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

// NewNode ...
func NewNode(config Config) *Node {
	node := &Node{
		schema: FirstSchema(),
		// funcs: map[string]Action{},

		err:      make(chan error),
		shutdown: make(chan bool),

		logger: zerolog.Logger{},
	}

	node.Hosts = config.Hosts
	node.Journal = config.Journal
	node.Sources = config.Sources

	_ = node.connect()

	return node
}

// Node ...
type Node struct {
	Name string

	Hosts   []string
	Journal map[string]string
	Sources map[string]map[string]string

	// Run
	journal journal
	main    source

	// Schema
	schema *jsonapi.Schema
	funcs  map[string]Action

	// Channels
	err      chan error
	shutdown chan bool

	// Internal
	logger zerolog.Logger
	sync.Mutex
}

// Handle ...
func (n *Node) Handle(r *Request) *jsonapi.Document {
	if !n.journal.alive || !n.main.alive {
		if !n.connect() {
			panic("cannot connect to necessary services")
		}
	}

	if n.funcs == nil {
		n.funcs = map[string]Action{
			"POST 0_sets":    ActionPostSet,
			"POST 0_attrs":   ActionPostSet,
			"POST 0_rels":    ActionPostSet,
			"PATCH 0_sets":   ActionPatchSet,
			"PATCH 0_attrs":  ActionPatchSet,
			"PATCH 0_rels":   ActionPatchSet,
			"DELETE 0_sets":  ActionDeleteSet,
			"DELETE 0_attrs": ActionDeleteSet,
			"DELETE 0_rels":  ActionDeleteSet,
		}
	}

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
		ops:  []query.Op{},
	}

	// Check password is correct if request is writing (non-GET).
	if r.Method == POST || r.Method == PATCH || r.Method == DELETE {
		pwRes, _ := cp.tx.Resource(query.Res{
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

	execute := n.funcs[r.Method+" "+r.URL.ResType]
	if execute == nil {
		execute = ActionDefault
	}

	ops := []query.Op{}
	// Prepare action
	switch r.Method {
	case GET:
	case POST:
		if res.GetID() == "" {
			res.SetID(uuid.New().String()[:8])
		}

		ops = query.NewOpCreateRes(res)

		found, _ := cp.tx.Resource(query.Res{
			Set:    res.GetType().Name,
			ID:     res.GetID(),
			Fields: []string{"id"},
		})

		if found != nil {
			cp.Fail(errors.New("id already used"))
		}
	case PATCH:
		ops = []query.Op{}

		for _, attr := range res.Attrs() {
			ops = append(ops, query.NewOpSet(
				r.URL.ResType,
				res.GetID(),
				attr.Name,
				res.Get(attr.Name),
			))
		}

		for _, rel := range res.Rels() {
			if rel.ToOne {
				ops = append(ops, query.NewOpSet(
					r.URL.ResType,
					res.GetID(),
					rel.FromName,
					res.GetToOne(rel.FromName),
				))
			} else {
				ops = append(ops, query.NewOpSet(
					r.URL.ResType,
					res.GetID(),
					rel.FromName,
					res.GetToMany(rel.FromName),
				))
			}
		}
	case DELETE:
		ops = []query.Op{query.NewOpSet(r.URL.ResType, r.URL.ResID, "id", "")}
	}

	cp.Apply(ops)

	// Execute
	execute(cp)

	if r.isSchemaChange() {
		err = updateSchema(n.schema, cp.tx)
		if err != nil {
			panic(fmt.Errorf("could not update schema: %s", err))
		}
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
		if len(ops) > 0 {
			enc, err := query.Encode(0, cp.ops)
			if err != nil {
				panic(fmt.Errorf("could not encode ops: %s", err))
			}

			// Commit the entry
			err = n.journal.jrnl.Append(enc)
			if err != nil {
				n.journal.alive = false
				panic(fmt.Errorf("could not append: %s", err))
			}

			err = cp.commit()
			if err != nil {
				n.main.alive = false
				panic(fmt.Errorf("could not commit: %s", err))
			}
		}

		// Response payload
		switch r.Method {
		case GET:
			if !r.URL.IsCol {
				res := cp.Resource(query.NewRes(r.URL))
				doc.Data = res
			} else {
				col := &jsonapi.SoftCollection{}
				typ := n.schema.GetType(r.URL.ResType)
				col.SetType(&typ)
				resources := cp.Collection(query.NewCol(r.URL))
				for i := 0; i < resources.Len(); i++ {
					col.Add(resources.At(i))
				}
				doc.Data = jsonapi.Collection(col)
			}
		case POST, PATCH:
			qry := query.NewRes(r.URL)
			qry.ID = res.GetID()
			res := cp.Resource(qry)
			doc.Data = res
		case DELETE:
		}
	}

	return doc
}

func (n *Node) connect() bool {
	n.Lock()
	defer n.Unlock()

	if !n.main.alive {
		src, err := newSource(n.Sources["main"])
		if err != nil {
			return false
		}

		n.main.src = src
		n.main.alive = true
	}

	if !n.journal.alive {
		jrnl, err := newJournal(n.Journal)
		if err != nil {
			return false
		}

		_ = jrnl.Reset()

		n.journal.jrnl = jrnl
		n.journal.alive = true
	}

	return true
}
