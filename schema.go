package karigo

import (
	"github.com/mfcochauxlaberge/jsonapi"
)

// FirstSchema ...
func FirstSchema() *jsonapi.Schema {
	schema := &jsonapi.Schema{}

	typ, err := jsonapi.BuildType(meta{})
	if err != nil {
		panic(err)
	}
	err = schema.AddType(typ)
	if err != nil {
		panic(err)
	}

	typ, err = jsonapi.BuildType(set{})
	if err != nil {
		panic(err)
	}
	err = schema.AddType(typ)
	if err != nil {
		panic(err)
	}

	typ, err = jsonapi.BuildType(attr{})
	if err != nil {
		panic(err)
	}
	err = schema.AddType(typ)
	if err != nil {
		panic(err)
	}

	typ, err = jsonapi.BuildType(rel{})
	if err != nil {
		panic(err)
	}
	err = schema.AddType(typ)
	if err != nil {
		panic(err)
	}

	typ, err = jsonapi.BuildType(log{})
	if err != nil {
		panic(err)
	}
	err = schema.AddType(typ)
	if err != nil {
		panic(err)
	}

	typ, err = jsonapi.BuildType(op{})
	if err != nil {
		panic(err)
	}
	err = schema.AddType(typ)
	if err != nil {
		panic(err)
	}

	return schema
}

// meta ...
type meta struct {
	ID string `json:"id" api:"0_meta"`

	// Attributes
	Value string `json:"value" api:"attr"`
}

// set ...
type set struct {
	ID string `json:"id" api:"0_sets"`

	// Attributes
	Name    string `json:"name" api:"attr"`
	Version uint   `json:"version" api:"attr"`
	Active  bool   `json:"active" api:"attr"`

	// Relationships
	Attrs []string `json:"attrs" api:"rel,0_attrs,set"`
	Rels  []string `json:"rels" api:"rel,0_rels,from-set"`
}

// attr ...
type attr struct {
	ID string `json:"id" api:"0_attrs"`

	// Attributes
	Name   string `json:"name" api:"attr"`
	Type   string `json:"type" api:"attr"`
	Null   bool   `json:"null" api:"attr"`
	Active bool   `json:"active" api:"attr"`

	// Relationships
	Set string `json:"set" api:"rel,0_sets,attrs"`
}

// rel ...
type rel struct {
	ID string `json:"id" api:"0_rels"`

	// Attributes
	FromName string `json:"from-name" api:"attr"`
	ToOne    bool   `json:"to-one" api:"attr"`
	ToName   string `json:"to-name" api:"attr"`
	FromOne  bool   `json:"from-one" api:"attr"`
	Active   bool   `json:"active" api:"attr"`

	// Relationships
	FromSet string `json:"from-set" api:"rel,0_sets,rels"`
	ToSet   string `json:"to-set" api:"rel,0_sets"`
}

// log ...
type log struct {
	ID string `json:"id" api:"0_log"`

	// Relationships
	Ops []string `json:"ops" api:"rel,0_ops,version"`
}

// op ...
type op struct {
	ID string `json:"id" api:"0_ops"`

	// Attributes
	Key   string `json:"set" api:"attr"`
	Op    string `json:"op" api:"attr"`
	Value string `json:"value" api:"attr"`

	// Relationships
	Version string `json:"version" api:"rel,0_log,ops"`
}

func handleSchemaChange(s *jsonapi.Schema, r *Request, cp *Checkpoint) {
	var (
		res jsonapi.Resource
		err error
	)

	res, _ = r.Doc.Data.(jsonapi.Resource)

	if r.Method == "PATCH" {
		// Can only be for activating or deactivating
		// a set, attribute, or relationship.
		if active, ok := res.Get("active").(bool); ok {
			if active {
				switch r.URL.ResType {
				case "0_sets":
					err = activateSet(s, res.GetID())
				case "0_attrs":
					err = activateAttr(s, res)
				case "0_rels":
					err = activateRel(s, res)
				}
			} else {
				switch r.URL.ResType {
				case "0_sets":
					deactivateSet(s, res)
				case "0_attrs":
					deactivateAttr(s, res)
				case "0_rels":
					deactivateRel(s, res)
				}
			}
		}

		cp.Check(err)
	}
}

func activateSet(s *jsonapi.Schema, name string) error {
	err := s.AddType(jsonapi.Type{
		Name: name,
	})
	return err
}

func deactivateSet(s *jsonapi.Schema, res jsonapi.Resource) {
	s.RemoveType(res.GetID())
}

func activateAttr(s *jsonapi.Schema, res jsonapi.Resource) error {
	err := s.AddAttr(res.GetToOne("set"), jsonapi.Attr{
		Name:     res.Get("name").(string),
		Type:     res.Get("type").(int),
		Nullable: res.Get("nullable").(bool),
	})
	return err
}

func deactivateAttr(s *jsonapi.Schema, res jsonapi.Resource) {
	s.RemoveAttr(res.GetToOne("set"), res.Get("Name").(string))
}

func activateRel(s *jsonapi.Schema, res jsonapi.Resource) error {
	rel := jsonapi.Rel{
		FromType: res.GetToOne("from-set"),
		FromName: res.Get("from-name").(string),
		ToOne:    res.Get("to-one").(bool),
		ToType:   res.GetToOne("to-type"),
		ToName:   res.Get("to-name").(string),
		FromOne:  res.Get("from-one").(bool),
	}
	rel.Normalize()
	err := s.AddRel(res.GetToOne("set"), rel)
	return err
}

func deactivateRel(s *jsonapi.Schema, res jsonapi.Resource) {
	s.RemoveRel(res.Get("from-type").(string), res.Get("name").(string))
}
