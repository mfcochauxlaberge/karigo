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

	// Relationships
	Version string `json:"version" api:"rel,0_log,ops"`
}

func handleSchemaChange(r *Request, cp *Checkpoint, s *jsonapi.Schema) {}
