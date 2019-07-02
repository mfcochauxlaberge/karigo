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

		schema: FirstSchema(),
		// funcs: map[string]Tx{},

		err:      make(chan error),
		shutdown: make(chan bool),
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

// FirstSchema ...
func FirstSchema() *jsonapi.Schema {
	schema := &jsonapi.Schema{}

	// Meta
	typ := &jsonapi.Type{
		Name: "0_meta",
	}
	typ.AddAttr(jsonapi.Attr{
		Name:     "value",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})
	schema.AddType(*typ)

	// Sets
	typ = &jsonapi.Type{
		Name: "0_sets",
	}
	typ.AddAttr(jsonapi.Attr{
		Name:     "name",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name:     "version",
		Type:     jsonapi.AttrTypeUint,
		Nullable: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name:     "active",
		Type:     jsonapi.AttrTypeBool,
		Nullable: false,
	})
	typ.AddRel(jsonapi.Rel{
		Name:         "attrs",
		Type:         "0_attrs",
		ToOne:        false,
		InverseName:  "set",
		InverseType:  "0_sets",
		InverseToOne: true,
	})
	typ.AddRel(jsonapi.Rel{
		Name:         "rels",
		Type:         "0_rels",
		ToOne:        false,
		InverseName:  "set",
		InverseType:  "0_sets",
		InverseToOne: true,
	})
	schema.AddType(*typ)

	// Attrs
	typ = &jsonapi.Type{
		Name: "0_attrs",
	}
	typ.AddAttr(jsonapi.Attr{
		Name:     "name",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name:     "type",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name:     "null",
		Type:     jsonapi.AttrTypeBool,
		Nullable: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name:     "active",
		Type:     jsonapi.AttrTypeBool,
		Nullable: false,
	})
	typ.AddRel(jsonapi.Rel{
		Name:         "set",
		Type:         "0_sets",
		ToOne:        true,
		InverseName:  "attrs",
		InverseType:  "0_attrs",
		InverseToOne: false,
	})
	schema.AddType(*typ)

	// Rels
	typ = &jsonapi.Type{
		Name: "0_rels",
	}
	typ.AddAttr(jsonapi.Attr{
		Name:     "name",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name:     "to-one",
		Type:     jsonapi.AttrTypeBool,
		Nullable: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name:     "active",
		Type:     jsonapi.AttrTypeBool,
		Nullable: false,
	})
	typ.AddRel(jsonapi.Rel{
		Name:         "inverse",
		Type:         "0_rels",
		ToOne:        true,
		InverseName:  "inverse",
		InverseType:  "0_rels",
		InverseToOne: true,
	})
	typ.AddRel(jsonapi.Rel{
		Name:         "set",
		Type:         "0_sets",
		ToOne:        true,
		InverseName:  "rels",
		InverseType:  "0_rels",
		InverseToOne: false,
	})
	schema.AddType(*typ)

	return schema
}
