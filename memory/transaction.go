package memory

import (
	"reflect"
	"sync"

	"github.com/mfcochauxlaberge/karigo"

	"github.com/mfcochauxlaberge/jsonapi"
)

var _ karigo.Source = (*Source)(nil)

// Tx ...
type Tx struct {
	src  *Source
	sets map[string]*jsonapi.SoftCollection

	sync.Mutex
}

// Resource ...
func (s *Tx) Resource(qry karigo.QueryRes) (jsonapi.Resource, error) {
	s.Lock()
	defer s.Unlock()

	// Get resource
	res := s.sets[qry.Set].Resource(qry.ID, qry.Fields)

	return res, nil
}

// Collection ...
func (s *Tx) Collection(qry karigo.QueryCol) (jsonapi.Collection, error) {
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
func (s *Tx) Apply(ops []karigo.Op) error {
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
func (s *Tx) Commit() error {
	return nil
}

// Rollback ...
func (s *Tx) Rollback() error {
	return nil
}

func (s *Tx) opSet(set, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s = %v\n", set, id, field, v)
	// Type change
	switch set {
	case "0_sets":
		if id != "" && field == "created" && v.(bool) {
			// New set
			s.sets[id] = &jsonapi.SoftCollection{}
			s.sets[id].SetType(&jsonapi.Type{
				Name: id,
			})
		}
	case "0_attrs":
		if id != "" && field == "created" && v.(bool) {
			// New attribute
			attr := s.sets["0_attrs"].Resource(id, nil)
			typ, null := jsonapi.GetAttrType(attr.Get("type").(string))

			_ = s.sets[attr.GetToOne("set")].Type.AddAttr(
				jsonapi.Attr{
					Name:     attr.Get("name").(string),
					Type:     typ,
					Nullable: null,
				},
			)
		}
	case "0_rels":
		if id != "" && field == "created" && v.(bool) {
			// New relationship
			rel := s.sets["0_rels"].Resource(id, nil)

			_ = s.sets[rel.GetToOne("from-set")].Type.AddRel(
				jsonapi.Rel{
					FromType: rel.GetToOne("from-set"),
					FromName: rel.Get("from-name").(string),
					ToOne:    rel.Get("to-one").(bool),
					ToType:   rel.GetToOne("to-set"),
					ToName:   rel.Get("to-name").(string),
					FromOne:  rel.Get("from-one").(bool),
				},
			)
		}
	}

	switch {
	case id != "" && field != "id":
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
	case id == "" && field == "id":
		// Create a resource
		typ := s.sets[set].Type
		s.sets[set].Add(makeSoftResource(typ, v.(string), map[string]interface{}{}))
	case id != "" && field == "id" && v.(string) == "":
		// Delete a resource
		s.sets[set].Remove(id)
	}
}

func (s *Tx) opAdd(set, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s += %v\n", set, id, field, v)
	curr := reflect.ValueOf(s.sets[set].Resource(id, nil).GetToMany(field))
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
