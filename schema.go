package karigo

import (
	"errors"
	"fmt"

	"github.com/mfcochauxlaberge/jsonapi"
)

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
	Rels  []string `json:"rels" api:"rel,0_attrs,set"`
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
	Name          string `json:"name" api:"attr"`
	ToOne         bool   `json:"to-one" api:"attr"`
	Active        bool   `json:"active" api:"attr"`
	InverseName   string `json:"inverse-name" api:"attr"`
	InverseToOne  bool   `json:"inverse-to-one" api:"attr"`
	InverseActive bool   `json:"inverse-active" api:"attr"`

	// Relationships
	Set        string `json:"set" api:"rel,0_sets,rels"`
	InverseSet string `json:"inverse-set" api:"rel,0_rels"`
}

func handleSchemaChange(r *Request, cp *Checkpoint, s *jsonapi.Schema) {
	var (
		res jsonapi.Resource
		ops []Op
		err error
	)

	res, _ = r.Body.Data.(jsonapi.Resource)

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
			ops, err = activateSet(s, res)
			cp.Check(err)
			cp.Apply(ops)
		}

		if deactivate, ok := res.Get("active").(bool); !deactivate && ok {
			ops, err = deactivateSet(s, res)
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
	typ := jsonapi.Type{
		Name: res.Get("name").(string),
	}

	err := s.AddType(typ)
	if err != nil {
		return nil, err
	}

	id := typ.Name
	ops := []Op{NewOpAdd("0_sets", "", "id", id)}
	return ops, nil
}

func deleteSet(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
	return []Op{}, nil
}

func activateSet(s *jsonapi.Schema, res jsonapi.Resource) ([]Op, error) {
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
