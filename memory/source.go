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
		ctyp := jsonapi.CopyType(typ)
		types[ctyp.Name] = &ctyp
	}

	s.sets = map[string]*jsonapi.SoftCollection{}

	// Sets
	for _, typ := range types {
		s.sets[typ.Name] = &jsonapi.SoftCollection{}
		s.sets[typ.Name].SetType(typ)
	}

	// Types and attributes
	for _, typ := range types {
		typ := jsonapi.CopyType(*typ)

		attrIDs := []string{}
		relIDs := []string{}
		for _, field := range typ.Fields() {
			if attr, ok := typ.Attrs[field]; ok {
				attrIDs = append(attrIDs, typ.Name+"_"+attr.Name)
			} else if rel, ok := typ.Rels[field]; ok {
				if rel.FromType+rel.FromName ==
					rel.Inverse().FromType+rel.Inverse().FromName {
					relIDs = append(relIDs, typ.Name+"_"+rel.FromName)
				}
			}
		}

		s.sets["0_sets"].Add(makeSoftResource(
			types["0_sets"],
			typ.Name,
			map[string]interface{}{
				"name":    typ.Name,
				"version": 0,
				"active":  true,
				"attrs":   attrIDs,
				"rels":    relIDs,
			},
		))

		// 0_attrs
		for _, attr := range typ.Attrs {
			s.sets["0_attrs"].Add(makeSoftResource(
				types["0_attrs"],
				typ.Name+"_"+attr.Name,
				map[string]interface{}{
					"name":   attr.Name,
					"type":   jsonapi.GetAttrTypeString(attr.Type, false),
					"null":   attr.Nullable,
					"active": true,
					"set":    typ.Name,
				},
			))
		}
	}

	// Relationships
	for _, rel := range karigo.FirstSchema().Rels() {
		s.sets["0_rels"].Add(makeSoftResource(
			types["0_rels"],
			rel.FromType+"_"+rel.FromName,
			map[string]interface{}{
				"from-name": rel.FromName,
				"to-one":    rel.ToOne,
				"to-name":   rel.ToName,
				"from-one":  rel.FromOne,
				"active":    true,
				"from-set":  rel.FromType,
				"to-set":    rel.ToType,
			},
		))
	}

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
func (s *Source) Collection(qry karigo.QueryCol) (jsonapi.Collection, error) {
	s.Lock()
	defer s.Unlock()

	// BelongsToFilter
	var ids []string
	if qry.BelongsToFilter.ID != "" {
		res := s.sets[qry.BelongsToFilter.Type].Resource(qry.BelongsToFilter.ID, []string{})
		ids = res.GetToMany(qry.BelongsToFilter.Name)
	}

	// Get all records from the given set
	recs := jsonapi.Range(
		s.sets[qry.Set],
		ids,
		nil,
		qry.Sort,
		qry.PageSize,
		qry.PageNumber,
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

// Commit ...
func (s *Source) Commit() error {
	return nil
}

// Rollback ...
func (s *Source) Rollback() error {
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
			setID := s.sets["0_attrs"].Resource(id, nil).GetToOne("set")
			attrName := s.sets["0_attrs"].Resource(id, nil).Get("name").(string)
			attrType, _ := jsonapi.GetAttrType(
				s.sets["0_attrs"].Resource(id, nil).Get("type").(string),
			)
			_ = s.sets[setID].Type.AddAttr(jsonapi.Attr{
				Name:     attrName,
				Type:     attrType,
				Nullable: s.sets["0_attrs"].Resource(id, nil).Get("null").(bool),
			})
		}
	} else if set == "0_rels" {
		if id != "" && field == "active" && v.(bool) {
			// New relationship
			setID := s.sets["0_rels"].Resource(id, nil).GetToOne("from-set")
			relName := s.sets["0_rels"].Resource(id, nil).Get("from-name").(string)
			_ = s.sets[setID].Type.AddRel(jsonapi.Rel{
				// FromType:  s.sets["0_rels"].Resource(id,nil).Get("type").(string),
				FromName: relName,
				ToOne:    s.sets["0_rels"].Resource(id, nil).Get("to-one").(bool),
				ToType:   s.sets["0_rels"].Resource(id, nil).GetToOne("from-set"),
				// ToName:  id,
				// FromOne: s.sets["0_rels"].Resource(id,nil).Get("to-one").(bool),
			})
		}
	}

	if id != "" && field != "id" {
		// Set a field
		typ := s.sets[set].Type
		for _, attr := range typ.Attrs {
			if attr.Name == field {
				s.sets[set].Resource(id, nil).Set(field, v)
			}
		}
		for _, rel := range typ.Rels {
			if rel.FromName == field {
				if rel.ToOne {
					s.sets[set].Resource(id, nil).SetToOne(field, v.(string))
				} else {
					s.sets[set].Resource(id, nil).SetToMany(field, v.([]string))
				}
			}
		}
	} else if id == "" && field == "id" {
		// Create a resource
		typ := s.sets[set].Type
		s.sets[set].Add(makeSoftResource(typ, v.(string), map[string]interface{}{}))
	} else if id != "" && field == "id" && v.(string) == "" {
		// Delete a resource
		s.sets[set].Remove(id)
		// } else {
		// Should not happen
		// TODO Should this code path be reported?
	}
}

func (s *Source) opAdd(set, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s += %v\n", set, id, field, v)

	curr := reflect.ValueOf(s.sets[set].Resource(id, nil).Get(field))
	curr = reflect.Append(curr, reflect.ValueOf(v))

	typ := s.sets[set].Type
	for _, attr := range typ.Attrs {
		if attr.Name == field {
			s.sets[set].Resource(id, nil).Set(field, v)
		}
	}
	for _, rel := range typ.Rels {
		if rel.FromName == field {
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
			if rel.FromName == f {
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
