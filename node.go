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
	var (
		res jsonapi.Resource
		id  string
	)
	if r.Body != nil && r.Body.Data != nil {
		res, _ = r.Body.Data.(jsonapi.Resource)
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
		ops = []Op{NewOpSet(r.URL.ResType, "", "id", id)}
	case PATCH:
		n.logger.Debug("PATCH request")
		ops = []Op{NewOpSet(r.URL.ResType, r.URL.ResID, "id", id)}
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
					rel.Name,
					res.GetToOne(rel.Name),
				))
			} else {
				ops = append(ops, NewOpSet(
					r.URL.ResType,
					id,
					rel.Name,
					res.GetToMany(rel.Name),
				))
			}
		}
	case DELETE:
		n.logger.Debug("DELETE request")
		ops = []Op{NewOpSet(r.URL.ResType, r.URL.ResID, "id", "")}
	}

	doc := &jsonapi.Document{}

	// // Handle schema change
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

	cp := &Checkpoint{
		node: n,
		ops:  []Op{},
	}
	if r.isSchemaChange() {
		// Handle schema change
	} else {
		// Execution
		tx(cp, ops)
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
