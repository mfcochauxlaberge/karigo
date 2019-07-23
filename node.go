package karigo

import (
	"errors"
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

	var id string

	// Transaction
	tx := TxNothing
	switch r.Method {
	case "GET":
		n.logger.Debug("Executiong GET transaction")
	case "POST":
		id = uuid.NewV4().String()
		n.apply([]Op{NewOpSet(r.URL.ResType, "", "id", id)})
		n.logger.Debug("Executiong POST transaction")
	case "PATCH":
		n.logger.Debug("Executiong PATCH transaction")
	case "DELETE":
		n.logger.Debug("Executiong DELETE transaction")
	}

	doc := &jsonapi.Document{}

	// Execution
	cp := &Checkpoint{
		node: n,
		ops:  []Op{},
	}
	tx(cp)

	// Response payload
	switch r.Method {
	case "GET":
		if !r.URL.IsCol {
			res := cp.Resource(NewQueryRes(r.URL))
			doc.Data = jsonapi.Resource(res)
		} else {
			col := &jsonapi.SoftCollection{}
			typ := n.schema.GetType(r.URL.ResType)
			col.SetType(&typ)
			resources := cp.Collection(NewQueryCol(r.URL))
			for _, res := range resources {
				col.Add(res)
			}
			doc.Data = jsonapi.Collection(col)
		}
	case "POST", "PATCH":
		qry := NewQueryRes(r.URL)
		qry.ID = id
		res := cp.Resource(qry)
		doc.Data = jsonapi.Resource(res)
	case "DELETE":
		// cp.
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
	}

	return doc
}

// resource ...
func (n *Node) resource(v uint, qry QueryRes) (jsonapi.Resource, error) {
	// TODO Validate the query?

	return n.main.src.Resource(qry)
}

// collection ...
func (n *Node) collection(v uint, qry QueryCol) ([]jsonapi.Resource, error) {
	// TODO Validate the query?
	// TODO Complete the sorting rule

	return n.main.src.Collection(qry)
}

// do ...
func (n *Node) apply(ops []Op) error {
	err := n.main.src.Apply(ops)
	if err != nil {
		return errors.New("karigo: an operation could not be executed")
	}
	return nil
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
