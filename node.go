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
	n.Lock()
	n.Unlock()

	// Handle events
	for {
		select {}
	}
}

// Handle ...
func (n *Node) Handle(r *Request) *jsonapi.Document {
	var id string

	// Check for schema change
	// if r.Method == "POST" {
	// 	switch r.URL.ResType {
	// 	case "0_sets":
	// 		res := r.Body.Data.(jsonapi.Resource)
	// 		var (
	// 			name = res.Get("name")
	// 		)

	// 	case "0_attrs":
	// 	case "0_rels":
	// 	}
	// } else if r.Method == "PATCH" {
	// 	// if r.URL
	// }

	// Transaction
	tx := TxNothing
	switch r.Method {
	case "GET":
		n.logger.Debug("GET request")
	case "POST":
		id = uuid.NewV4().String()
		_ = n.apply([]Op{NewOpSet(r.URL.ResType, "", "id", id)})
		n.logger.Debug("POST request")
	case "PATCH":
		n.logger.Debug("PATCH request")
	case "DELETE":
		n.logger.Debug("DELETE request")
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
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "value",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})
	_ = schema.AddType(*typ)

	// Sets
	typ = &jsonapi.Type{
		Name: "0_sets",
	}
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "name",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "version",
		Type:     jsonapi.AttrTypeUint,
		Nullable: false,
	})
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "active",
		Type:     jsonapi.AttrTypeBool,
		Nullable: false,
	})
	_ = typ.AddRel(jsonapi.Rel{
		Name:         "attrs",
		Type:         "0_attrs",
		ToOne:        false,
		InverseName:  "set",
		InverseType:  "0_sets",
		InverseToOne: true,
	})
	_ = typ.AddRel(jsonapi.Rel{
		Name:         "rels",
		Type:         "0_rels",
		ToOne:        false,
		InverseName:  "set",
		InverseType:  "0_sets",
		InverseToOne: true,
	})
	_ = schema.AddType(*typ)

	// Attrs
	typ = &jsonapi.Type{
		Name: "0_attrs",
	}
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "name",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "type",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "null",
		Type:     jsonapi.AttrTypeBool,
		Nullable: false,
	})
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "active",
		Type:     jsonapi.AttrTypeBool,
		Nullable: false,
	})
	_ = typ.AddRel(jsonapi.Rel{
		Name:         "set",
		Type:         "0_sets",
		ToOne:        true,
		InverseName:  "attrs",
		InverseType:  "0_attrs",
		InverseToOne: false,
	})
	_ = schema.AddType(*typ)

	// Rels
	typ = &jsonapi.Type{
		Name: "0_rels",
	}
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "name",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "to-one",
		Type:     jsonapi.AttrTypeBool,
		Nullable: false,
	})
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "active",
		Type:     jsonapi.AttrTypeBool,
		Nullable: false,
	})
	_ = typ.AddRel(jsonapi.Rel{
		Name:         "inverse",
		Type:         "0_rels",
		ToOne:        true,
		InverseName:  "inverse",
		InverseType:  "0_rels",
		InverseToOne: true,
	})
	_ = typ.AddRel(jsonapi.Rel{
		Name:         "set",
		Type:         "0_sets",
		ToOne:        true,
		InverseName:  "rels",
		InverseType:  "0_rels",
		InverseToOne: false,
	})
	_ = schema.AddType(*typ)

	return schema
}
