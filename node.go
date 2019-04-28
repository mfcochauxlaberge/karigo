package karigo

import (
	"errors"
	"net/http"
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

		funcs: map[string]Tx{},

		requests: make(chan *http.Request),
		err:      make(chan error),
	}

	return node
}

// Node ...
type Node struct {
	// Run
	log  Journal
	main source

	// Funcs
	funcs map[string]Tx

	// Channels
	requests chan *http.Request
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
func (n *Node) Handle(r *http.Request) *jsonapi.Document {
	n.Lock()
	defer n.Unlock()

	// Tx
	var tx Tx
	tx = n.funcs[""]
	if tx == nil {
		tx = TxNotFound
	}

	doc := &jsonapi.Document{}

	cp := &Checkpoint{
		node: n,
		ops:  []Op{},
	}
	tx(cp)

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
