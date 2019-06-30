package karigo

import (
	"errors"
	"sync"

	"github.com/mfcochauxlaberge/jsonapi"
)

// NewNode ...
func NewNode(journal Journal, src Source) *Node {
	node := &Node{
		log: journal,
		main: source{
			src: src,
		},

		// funcs: map[string]Tx{},

		err: make(chan error),
	}

	return node
}

// Node ...
type Node struct {
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
	sync.Mutex
}

// Run ...
func (n *Node) Run() error {
	n.Lock()
	n.Unlock()

	// Handle events
	for {
		select {}
	}
}

// Handle ...
func (n *Node) Handle(r *Request) *jsonapi.Document {
	n.Lock()
	defer n.Unlock()

	// Transaction
	var tx Tx
	switch r.Method {
	case "GET":
		tx = TxGet
	case "POST":
		tx = TxCreate
	case "PATCH":
		tx = TxUpdate
	case "DELETE":
		tx = TxDelete
	default:
		tx = TxNotImplemented
	}

	doc := &jsonapi.Document{}

	// Execute
	cp := &Checkpoint{
		node: n,
		ops:  []Op{},
	}
	tx(cp)

	switch r.Method {
	case "GET":
		if r.URL.IsCol {
			cp.Collection(NewQueryCol(r.URL))
		} else {
			cp.Resource(NewQueryRes(r.URL))
		}
	case "POST", "PATCH":
		cp.Resource(NewQueryRes(r.URL))
	default:
		tx = TxNotImplemented
	}

	if cp.err != nil {
		var jaErr jsonapi.Error
		switch cp.err {
		case ErrNotImplemented:
			jaErr = jsonapi.NewErrNotImplemented()
		default:
			jaErr = jsonapi.NewErrInternalServerError()
		}
		doc.Errors = []jsonapi.Error{jaErr}
	} else {
		doc.Data = nil
	}

	return doc
}

// resource ...
func (n *Node) resource(v uint, qry QueryRes) (jsonapi.Resource, error) {
	// TODO Validate the query?

	return nil, errors.New("karigo: no source could handle the query")
}

// collection ...
func (n *Node) collection(v uint, qry QueryCol) ([]jsonapi.Resource, error) {
	// TODO Validate the query?
	// TODO Complete the sorting rule

	return n.main.src.Collection(qry)
}

// do ...
func (n *Node) do(ops []Op) error {
	return errors.New("karigo: an operation could not be executed")
}
