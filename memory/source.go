package memory

import (
	"reflect"
	"sync"

	"github.com/mfcochauxlaberge/karigo"

	"github.com/mfcochauxlaberge/jsonapi"
)

var _ karigo.Source = (*Source)(nil)

// Source ...
type Source struct {
	sets map[string]*jsonapi.SoftCollection

	sync.Mutex
}

// Reset ...
func (s *Source) Reset() error {
	s.Lock()
	defer s.Unlock()

	types := map[string]*jsonapi.Type{}
	for _, typ := range karigo.FirstSchema().Types {
		typ := typ
		types[typ.Name] = &typ
	}

	s.sets = map[string]*jsonapi.SoftCollection{}

	// 0_meta
	s.sets["0_meta"] = &jsonapi.SoftCollection{}
	s.sets["0_meta"].SetType(types["0_meta"])

	// 0_sets
	s.sets["0_sets"] = &jsonapi.SoftCollection{}
	s.sets["0_sets"].SetType(types["0_sets"])

	s.sets["0_sets"].Add(makeSoftResource(
		types["0_sets"],
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
		types["0_sets"],
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
			},
		},
	))
	s.sets["0_sets"].Add(makeSoftResource(
		types["0_sets"],
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
		types["0_sets"],
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

	// 0_attrs
	s.sets["0_attrs"] = &jsonapi.SoftCollection{}
	s.sets["0_attrs"].SetType(types["0_attrs"])

	s.sets["0_attrs"].Add(makeSoftResource(
		types["0_attrs"],
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
		types["0_attrs"],
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
		types["0_attrs"],
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
		types["0_attrs"],
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
		types["0_attrs"],
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
		types["0_attrs"],
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
		types["0_attrs"],
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
		types["0_attrs"],
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
		types["0_attrs"],
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
		types["0_attrs"],
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
		types["0_attrs"],
		"0_rels_active",
		map[string]interface{}{
			"name":   "active",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_rels",
		},
	))

	// 0_rels
	s.sets["0_rels"] = &jsonapi.SoftCollection{}
	s.sets["0_rels"].SetType(types["0_rels"])

	s.sets["0_rels"].Add(makeSoftResource(
		types["0_rels"],
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
		types["0_rels"],
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
		types["0_rels"],
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
		types["0_rels"],
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
		types["0_rels"],
		"0_rels_set",
		map[string]interface{}{
			"name":    "set",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_rels",
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
		case karigo.OpAdd:
			s.opAdd(op.Key.Set, op.Key.ID, op.Key.Field, op.Value)
		}
	}

	return nil
}

func (s *Source) opSet(set, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s = %v\n", set, id, field, v)

	// Type change
	if set == "0_sets" {
		if id != "" && field == "active" && v.(bool) {
			// New set
			s.sets[id] = &jsonapi.SoftCollection{}
			s.sets[id].SetType(&jsonapi.Type{
				Name: id,
			})
		}
	} else if set == "0_attrs" {
		if id != "" && field == "active" && v.(bool) {
			// New attribute
			setID := s.sets["0_attrs"].Resource(id, nil).Get("set").(string)
			attrName := s.sets["0_attrs"].Resource(id, nil).Get("name").(string)
			attrType, _ := jsonapi.GetAttrType(s.sets["0_attrs"].Resource(id, nil).Get("type").(string))
			s.sets[setID].GetType().AddAttr(jsonapi.Attr{
				Name:     attrName,
				Type:     attrType,
				Nullable: s.sets["0_attrs"].Resource(id, nil).Get("null").(bool),
			})
		}
	} else if set == "0_rels" {
		if id != "" && field == "active" && v.(bool) {
			// New relationship
			setID := s.sets["0_rels"].Resource(id, nil).Get("set").(string)
			relName := s.sets["0_rels"].Resource(id, nil).Get("name").(string)
			s.sets[setID].GetType().AddRel(jsonapi.Rel{
				Name:  relName,
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
		typ := s.sets[set].GetType()
		for _, attr := range typ.Attrs {
			if attr.Name == field {
				s.sets[set].Resource(id, nil).Set(field, v)
			}
		}
		for _, rel := range typ.Rels {
			if rel.Name == field {
				if rel.ToOne {
					s.sets[set].Resource(id, nil).SetToOne(field, v.(string))
				} else {
					s.sets[set].Resource(id, nil).SetToMany(field, v.([]string))
				}
			}
		}
	} else if id == "" && field == "id" {
		// Create a resource
		typ := s.sets[set].GetType()
		s.sets[set].Add(makeSoftResource(typ, v.(string), map[string]interface{}{}))
	} else if id != "" && field == "id" && v.(string) == "" {
		// Delete a resource
		s.sets[set].Remove(id)
	} else {
		// Should not happen
		// TODO Should this code path be reported?
	}
}

func (s *Source) opAdd(set, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s += %v\n", set, id, field, v)

	curr := reflect.ValueOf(s.sets[set].Resource(id, nil).Get(field))
	curr = reflect.Append(curr, reflect.ValueOf(v))

	typ := s.sets[set].GetType()
	for _, attr := range typ.Attrs {
		if attr.Name == field {
			s.sets[set].Resource(id, nil).Set(field, v)
		}
	}
	for _, rel := range typ.Rels {
		if rel.Name == field {
			if rel.ToOne {
				s.sets[set].Resource(id, nil).SetToOne(field, curr.Interface().(string))
			} else {
				s.sets[set].Resource(id, nil).SetToMany(field, curr.Interface().([]string))
			}
		}
	}
}

func makeSoftResource(typ *jsonapi.Type, id string, vals map[string]interface{}) *jsonapi.SoftResource {
	sr := &jsonapi.SoftResource{}
	sr.SetType(typ)
	sr.SetID(id)

	for f, v := range vals {
		for _, attr := range typ.Attrs {
			if attr.Name == f {
				sr.Set(f, v)
			}
		}
		for _, rel := range typ.Rels {
			if rel.Name == f {
				if rel.ToOne {
					sr.SetToOne(f, v.(string))
				} else {
					sr.SetToMany(f, v.([]string))
				}
			}
		}
	}

	return sr
}
