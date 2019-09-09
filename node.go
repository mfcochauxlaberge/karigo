package karigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
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

		logger: logrus.New(),
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
	logger *logrus.Logger
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
		id  string
		err error
	)

	if r.Method == POST || r.Method == PATCH {
		r.Doc, err = jsonapi.Unmarshal(r.Body, n.schema)
		if err != nil {
			panic(err)
		}
	}

	if r.Doc != nil && r.Doc.Data != nil {
		res, _ = r.Doc.Data.(jsonapi.Resource)
	}

	cp := &Checkpoint{
		node: n,
		ops:  []Op{},
	}

	n.logger.Debugf("Node %s received a request", n.Name)

	tx := TxDefault
	ops := []Op{}
	// Prepare transaction
	switch r.Method {
	case GET:
		n.logger.Debug("GET request")
	case POST:
		n.logger.Debug("POST request")
		id = uuid.NewV4().String()[:8]
		// TODO Do not hardcode the following condition.
		if res.GetType().Name == "0_meta" {
			id = res.GetID()
		}
		found, _ := n.resource(0, QueryRes{
			Set:    res.GetType().Name,
			ID:     id,
			Fields: []string{"id"},
		})
		if found != nil {
			cp.Fail(errors.New("id already used"))
		}
		ops = NewOpInsert(res)
	case PATCH:
		n.logger.Debug("PATCH request")
		id = res.GetID()
		ops = []Op{}
		for _, attr := range res.Attrs() {
			ops = append(ops, NewOpSet(
				r.URL.ResType,
				id,
				attr.Name,
				res.Get(attr.Name),
			))
		}
		for _, rel := range res.Rels() {
			if rel.ToOne {
				ops = append(ops, NewOpSet(
					r.URL.ResType,
					id,
					rel.FromName,
					res.GetToOne(rel.FromName),
				))
			} else {
				ops = append(ops, NewOpSet(
					r.URL.ResType,
					id,
					rel.FromName,
					res.GetToMany(rel.FromName),
				))
			}
		}
	case DELETE:
		n.logger.Debug("DELETE request")
		ops = []Op{NewOpSet(r.URL.ResType, r.URL.ResID, "id", "")}
	}
	cp.ops = ops

	doc := &jsonapi.Document{}

	if r.isSchemaChange() {
		// Handle schema change
		handleSchemaChange(r, cp, n.schema)
	} else {
		// Execute
		tx(cp)
	}

	// Add to journal
	if cp.err == nil {
		entry, err := json.Marshal(Entry(ops))
		if err != nil {
			cp.Fail(fmt.Errorf("karigo: could not marshal entry: %s", err))
		}
		err = n.log.Append(entry)
	}

	if cp.err != nil {
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
			qry.ID = id
			res := cp.Resource(qry)
			doc.Data = res
		case DELETE:
		}
	}

	return doc
}

// resource ...
func (n *Node) resource(_ uint, qry QueryRes) (jsonapi.Resource, error) {
	// TODO Validate the query?

	return n.main.src.Resource(qry)
}

// collection ...
func (n *Node) collection(_ uint, qry QueryCol) (jsonapi.Collection, error) {
	// TODO Validate the query?
	// TODO Complete the sorting rule

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
