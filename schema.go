package karigo

import (
	"errors"
	"fmt"

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
		ops []Op
		err error
	)

	res, _ = r.Doc.Data.(jsonapi.Resource)

	err = validateSchemaChange(res)
	if err != nil {
		cp.Check(err)
	}

	if r.Method == "POST" {
		if res.GetType().Name == "0_sets" {
			// Add set
			ops, err = addSet(s, res)
			cp.Check(err)
			cp.Apply(ops)
		} else if res.GetType().Name == "0_attrs" {
			// Add attribute
			ops, err = addAttr(s, res)
			cp.Check(err)
			cp.Apply(ops)
		} else if res.GetType().Name == "0_rels" {
			// Add relationship
			ops, err = addRel(s, res)
			cp.Check(err)
			cp.Apply(ops)
		}
	} else if r.Method == "PATCH" {
		// Can only be for activating or deactivating
		// a set, attribute, or relationship.
		if activate, ok := res.Get("active").(bool); activate && ok {
			switch r.URL.ResType {
			case "0_sets":
				ops, err = activateSet(s, res)
			case "0_attributes":
				ops, err = activateAttr(s, res)
			case "0_relationships":
				ops, err = activateRel(s, res)
			}
			cp.Check(err)
			cp.Apply(ops)
		}

		if deactivate, ok := res.Get("active").(bool); !deactivate && ok {
			switch r.URL.ResType {
			case "0_sets":
				ops, err = deactivateSet(s, res)
			case "0_attributes":
				ops, err = deactivateAttr(s, res)
			case "0_relationships":
				ops, err = deactivateRel(s, res)
			}
			cp.Check(err)
			cp.Apply(ops)
		}
	} else if r.Method == "DELETE" {
		currRes := cp.Resource(QueryRes{
			Set:    res.GetType().Name,
			ID:     res.GetID(),
			Fields: []string{"active"},
		})

		if currRes.GetID() == "" {
			cp.Fail(errors.New("schema element does not exist"))
		}

		if active, _ := currRes.Get("active").(bool); !active {
			// Only possible is active is false
			ops, err = deleteSet(s, res)
			cp.Check(err)
			cp.Apply(ops)
		} else {
			cp.Fail(errors.New("schema element is still active"))
		}
	}
}

func validateSchemaChange(res jsonapi.Resource) error {
	if res.Get("name") == "" {
		return fmt.Errorf("karigo: %q is not a valid JSON:API name", res.Get("name"))
	}
	return nil
}

func addSet(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	id := res.Get("name").(string)
	current := s.GetType(id)
	if current.Name != "" {
		return nil, fmt.Errorf("type %q already exists", id)
	}

	return NewOpAddSet(res.Get("name").(string)), nil
}

func deleteSet(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	name := res.Get("name").(string)
	current := s.GetType(name)
	if current.Name == "" {
		return nil, fmt.Errorf("type %q does not exist", name)
	}

	s.RemoveType(res.GetID())

	ops := []Op{
		NewOpSet(
			res.GetType().Name,
			res.GetID(),
			"id",
			"",
		),
	}
	return ops, nil
}

func activateSet(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	typ := jsonapi.Type{
		Name: res.Get("name").(string),
	}

	err := s.AddType(typ)
	if err != nil {
		return nil, err
	}

	return []Op{}, nil
}

func deactivateSet(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	return []Op{}, nil
}

func addAttr(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	return []Op{}, nil
}

func deleteAttr(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	return []Op{}, nil
}

func activateAttr(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	return []Op{}, nil
}

func deactivateAttr(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	return []Op{}, nil
}

func addRel(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	return []Op{}, nil
}

func deleteRel(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	return []Op{}, nil
}

func activateRel(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	return []Op{}, nil
}

func deactivateRel(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	return []Op{}, nil
}
