package memory

import (
	"sync"

	"github.com/mfcochauxlaberge/karigo"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Source ...
type Source struct {
	ID       string
	Location string

	sets map[string]*jsonapi.SoftCollection

	sync.Mutex
}

// Reset ...
func (s *Source) Reset() error {
	s.Lock()
	defer s.Unlock()

	s.sets = map[string]*jsonapi.SoftCollection{}

	// 0_meta
	typ := &jsonapi.Type{
		Name: "0_meta",
	}
	typ.AddAttr(jsonapi.Attr{
		Name: "value",
		Type: jsonapi.AttrTypeString,
		Null: false,
	})

	s.sets["0_meta"] = &jsonapi.SoftCollection{}
	s.sets["0_meta"].SetType(typ)

	// 0_sets
	typ = &jsonapi.Type{
		Name: "0_sets",
	}
	typ.AddAttr(jsonapi.Attr{
		Name: "name",
		Type: jsonapi.AttrTypeString,
		Null: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name: "version",
		Type: jsonapi.AttrTypeUint,
		Null: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name: "active",
		Type: jsonapi.AttrTypeBool,
		Null: false,
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

	s.sets["0_sets"] = &jsonapi.SoftCollection{}
	s.sets["0_sets"].SetType(typ)
	s.sets["0_sets"].Add(makeSoftResource(
		typ,
		"0_meta",
		map[string]interface{}{
			"name":    "0_meta",
			"version": 0,
			"active":  true,
			"attrs": []string{
				"0_meta_value",
			},
		},
	))
	s.sets["0_sets"].Add(makeSoftResource(
		typ,
		"0_sets",
		map[string]interface{}{
			"name":    "0_sets",
			"version": 0,
			"active":  true,
			"attrs": []string{
				"0_sets_name",
				"0_sets_version",
				"0_sets_active",
			},
			"rels": []string{
				"0_sets_attrs",
				"0_sets_rels",
				"0_funcs_get_func",
				"0_funcs_create_func",
				"0_funcs_update_func",
				"0_funcs_delete_func",
			},
		},
	))
	s.sets["0_sets"].Add(makeSoftResource(
		typ,
		"0_attrs",
		map[string]interface{}{
			"name":    "0_attrs",
			"version": 0,
			"active":  true,
			"attrs": []string{
				"0_attrs_name",
				"0_attrs_type",
				"0_attrs_null",
				"0_attrs_active",
			},
			"rels": []string{
				"0_attrs_set",
			},
		},
	))
	s.sets["0_sets"].Add(makeSoftResource(
		typ,
		"0_rels",
		map[string]interface{}{
			"name":    "0_rels",
			"version": 0,
			"active":  true,
			"attrs": []string{
				"0_rels_name",
				"0_rels_to-one",
				"0_rels_active",
			},
			"rels": []string{
				"0_rels_set",
			},
		},
	))
	s.sets["0_sets"].Add(makeSoftResource(
		typ,
		"0_funcs",
		map[string]interface{}{
			"name":    "0_funcs",
			"version": 0,
			"active":  true,
			"attrs": []string{
				"0_funcs_func",
			},
			"rels": []string{
				"0_funcs_set",
			},
		},
	))

	// 0_attrs
	typ = &jsonapi.Type{
		Name: "0_attrs",
	}
	typ.AddAttr(jsonapi.Attr{
		Name: "name",
		Type: jsonapi.AttrTypeString,
		Null: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name: "type",
		Type: jsonapi.AttrTypeUint,
		Null: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name: "null",
		Type: jsonapi.AttrTypeBool,
		Null: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name: "active",
		Type: jsonapi.AttrTypeBool,
		Null: false,
	})
	typ.AddRel(jsonapi.Rel{
		Name:         "set",
		Type:         "0_sets",
		ToOne:        true,
		InverseName:  "attrs",
		InverseType:  "0_attrs",
		InverseToOne: false,
	})

	s.sets["0_attrs"] = &jsonapi.SoftCollection{}
	s.sets["0_attrs"].SetType(typ)
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_meta_value",
		map[string]interface{}{
			"name":   "value",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_meta",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_sets_name",
		map[string]interface{}{
			"name":   "name",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_sets",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_sets_version",
		map[string]interface{}{
			"name":   "version",
			"type":   "int",
			"null":   false,
			"active": true,
			"set":    "0_sets",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_sets_active",
		map[string]interface{}{
			"name":   "active",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_sets",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_attrs_name",
		map[string]interface{}{
			"name":   "name",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_attrs",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_attrs_type",
		map[string]interface{}{
			"name":   "type",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_attrs",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_attrs_null",
		map[string]interface{}{
			"name":   "null",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_attrs",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_attrs_active",
		map[string]interface{}{
			"name":   "active",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_attrs",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_rels_name",
		map[string]interface{}{
			"name":   "name",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_rels",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_rels_to-one",
		map[string]interface{}{
			"name":   "to-one",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_rels",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_rels_active",
		map[string]interface{}{
			"name":   "active",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_rels",
		},
	))
	s.sets["0_attrs"].Add(makeSoftResource(
		typ,
		"0_funcs_func",
		map[string]interface{}{
			"name":   "func",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_funcs",
		},
	))

	// 0_rels
	typ = &jsonapi.Type{
		Name: "0_rels",
	}
	typ.AddAttr(jsonapi.Attr{
		Name: "name",
		Type: jsonapi.AttrTypeString,
		Null: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name: "to-one",
		Type: jsonapi.AttrTypeBool,
		Null: false,
	})
	typ.AddAttr(jsonapi.Attr{
		Name: "active",
		Type: jsonapi.AttrTypeBool,
		Null: false,
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

	s.sets["0_rels"] = &jsonapi.SoftCollection{}
	s.sets["0_rels"].SetType(typ)
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_sets_attrs",
		map[string]interface{}{
			"name":    "attrs",
			"to-one":  false,
			"active":  true,
			"inverse": "0_attrs_set",
			"set":     "0_sets",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_sets_rels",
		map[string]interface{}{
			"name":    "rels",
			"to-one":  false,
			"active":  true,
			"inverse": "0_rels_set",
			"set":     "0_sets",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_sets_get_func",
		map[string]interface{}{
			"name":    "get_func",
			"to-one":  true,
			"active":  true,
			"inverse": "", // TODO Inverse?
			"set":     "0_funcs",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_sets_create_func",
		map[string]interface{}{
			"name":    "create_func",
			"to-one":  true,
			"active":  true,
			"inverse": "", // TODO Inverse?
			"set":     "0_funcs",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_sets_update_func",
		map[string]interface{}{
			"name":    "update_func",
			"to-one":  true,
			"active":  true,
			"inverse": "", // TODO Inverse?
			"set":     "0_funcs",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_sets_delete_func",
		map[string]interface{}{
			"name":    "delete_func",
			"to-one":  true,
			"active":  true,
			"inverse": "", // TODO Inverse?
			"set":     "0_funcs",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_attrs_set",
		map[string]interface{}{
			"name":    "set",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_attrs",
			"set":     "0_attrs",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_rels_inverse",
		map[string]interface{}{
			"name":    "inverse",
			"to-one":  true,
			"active":  true,
			"inverse": "0_rels_inverse",
			"set":     "0_rels",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_rels_set",
		map[string]interface{}{
			"name":    "set",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_rels",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_funcs_get_func",
		map[string]interface{}{
			"name":    "get_func",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_funcs_create_func",
		map[string]interface{}{
			"name":    "create_func",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_funcs_update_func",
		map[string]interface{}{
			"name":    "update_func",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))
	s.sets["0_rels"].Add(makeSoftResource(
		typ,
		"0_funcs_delete_func",
		map[string]interface{}{
			"name":    "delete_func",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))

	// 0_funcs
	typ = &jsonapi.Type{
		Name: "funcs",
	}
	typ.AddAttr(jsonapi.Attr{
		Name: "func",
		Type: jsonapi.AttrTypeString,
		Null: false,
	})

	s.sets["0_funcs"] = &jsonapi.SoftCollection{}
	s.sets["0_funcs"].SetType(typ)
	s.sets["0_funcs"].Add(makeSoftResource(
		typ,
		"_not_implemented",
		map[string]interface{}{
			"func": `func(snap *Snapshot) error {
				snap.Fail(ErrNotImplemented)
			}`,
		},
	))

	return nil
}

// Resource ...
func (s *Source) Resource(qry karigo.QueryRes) (jsonapi.Resource, error) {
	s.Lock()
	defer s.Unlock()

	// Get resource
	res := s.sets[qry.Set].Resource(qry.ID, qry.Fields)

	return res, nil
}

// Collection ...
func (s *Source) Collection(qry karigo.QueryCol) ([]jsonapi.Resource, error) {
	s.Lock()
	defer s.Unlock()

	// BelongsToFilter
	var ids []string
	if qry.BelongsToFilter.ID != "" {
		res := s.sets[qry.BelongsToFilter.Type].Resource(qry.BelongsToFilter.ID, []string{})
		ids = res.GetToMany(qry.BelongsToFilter.Name)
	}

	// Get all records from the given set
	recs := s.sets[qry.Set].Range(
		ids,
		nil,
		qry.Sort,
		qry.Fields,
		uint(qry.PageSize),
		uint(qry.PageNumber),
	)

	return recs, nil
}

// Apply ...
func (s *Source) Apply(ops []karigo.Op) error {
	s.Lock()
	defer s.Unlock()

	for _, op := range ops {
		switch op.Op {
		case karigo.OpSet:
			s.opSet(op.Key.Set, op.Key.ID, op.Key.Field, op.Value)
		}
	}

	return nil
}

func (s *Source) opSet(setname, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s = %v\n", setname, id, field, v)

	// Type change
	if setname == "0_sets" {
		if id != "" && field == "active" && v.(bool) {
			// New set
			s.sets[id] = &jsonapi.SoftCollection{}
			s.sets[id].SetType(&jsonapi.Type{
				Name: id,
			})
		}
	} else if setname == "0_attrs" {
		if id != "" && field == "active" && v.(bool) {
			// New attribute
			setID := s.sets["0_attrs"].Resource(id, nil).Get("set").(string)
			attrType, _ := jsonapi.GetAttrType(s.sets["0_attrs"].Resource(id, nil).Get("type").(string))
			s.sets[setID].GetType().AddAttr(jsonapi.Attr{
				Name: id,
				Type: attrType,
				Null: s.sets["0_attrs"].Resource(id, nil).Get("null").(bool),
			})
		}
	} else if setname == "0_rels" {
		if id != "" && field == "active" && v.(bool) {
			// New relationship
			setID := s.sets["0_rels"].Resource(id, nil).Get("set").(string)
			s.sets[setID].GetType().AddRel(jsonapi.Rel{
				Name:  id,
				Type:  s.sets["0_rels"].Resource(id, nil).Get("set").(string),
				ToOne: s.sets["0_rels"].Resource(id, nil).Get("to-one").(bool),
				// InverseName:  id,
				// InverseType:  s.sets["0_rels"].Resource(id,nil).Get("type").(string),
				// InverseToOne: s.sets["0_rels"].Resource(id,nil).Get("to-one").(bool),
			})
		}
	}

	if id != "" && field != "id" {
		// Set a field
		s.sets[setname].Resource(id, nil).Set(field, v)
	} else if id == "" && field == "id" {
		// Create a resource
		typ := s.sets[setname].GetType()
		s.sets[setname].Add(makeSoftResource(typ, v.(string), map[string]interface{}{}))
	} else if id != "" && field == "id" && v.(string) == "" {
		// Delete a resource
		s.sets[setname].Remove(id)
	} else {
		// Should not happen
		// TODO Should this code path be reported?
	}
}

func makeSoftResource(typ *jsonapi.Type, id string, vals map[string]interface{}) *jsonapi.SoftResource {
	sr := &jsonapi.SoftResource{}
	sr.SetType(typ)
	sr.SetID(id)

	for f, v := range vals {
		sr.Set(f, v)
	}

	return sr
}
