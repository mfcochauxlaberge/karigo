package memory

import (
	"reflect"
	"sync"

	"github.com/mfcochauxlaberge/karigo/query"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Tx ...
type Tx struct {
	src  *Source
	sets map[string]*jsonapi.SoftCollection

	sync.Mutex
}

// Resource ...
func (t *Tx) Resource(qry query.Res) (jsonapi.Resource, error) {
	t.Lock()
	defer t.Unlock()

	// Get resource
	res := t.sets[qry.Set].Resource(qry.ID, qry.Fields)

	return res, nil
}

// Collection ...
func (t *Tx) Collection(qry query.Col) (jsonapi.Collection, error) {
	t.Lock()
	defer t.Unlock()

	// BelongsToFilter
	var ids []string

	if qry.BelongsToFilter.ID != "" {
		res := t.sets[qry.BelongsToFilter.Type].Resource(qry.BelongsToFilter.ID, []string{})
		ids = res.GetToMany(qry.BelongsToFilter.Name)
	}

	// Get all records from the given set
	recs := jsonapi.Range(
		t.sets[qry.Set],
		ids,
		nil,
		qry.Sort,
		qry.PageSize,
		qry.PageNumber,
	)

	return recs, nil
}

// Apply ...
func (t *Tx) Apply(ops []query.Op) error {
	t.Lock()
	defer t.Unlock()

	for _, op := range ops {
		switch op.Op {
		case query.OpSet:
			t.opSet(op.Key.Set, op.Key.ID, op.Key.Field, op.Value)
		case query.OpInsert:
			t.opInsert(op.Key.Set, op.Key.ID, op.Key.Field, op.Value)
		}
	}

	return nil
}

// Commit ...
func (t *Tx) Commit() error {
	return nil
}

// Rollback ...
func (t *Tx) Rollback() error {
	return nil
}

func (t *Tx) opSet(set, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s = %v\n", set, id, field, v)
	// Type change
	switch set {
	case "0_sets":
		if id != "" && field == "created" && v.(bool) {
			// New set
			t.sets[id] = &jsonapi.SoftCollection{}
			t.sets[id].SetType(&jsonapi.Type{
				Name: id,
			})
		}
	case "0_attrs":
		if id != "" && field == "created" && v.(bool) {
			// New attribute
			attr := t.sets["0_attrs"].Resource(id, nil)
			typ, null := jsonapi.GetAttrType(attr.Get("type").(string))

			_ = t.sets[attr.GetToOne("set")].Type.AddAttr(
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
			rel := t.sets["0_rels"].Resource(id, nil)

			_ = t.sets[rel.GetToOne("from-set")].Type.AddRel(
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
		typ := t.sets[set].Type
		for _, attr := range typ.Attrs {
			if attr.Name == field {
				t.sets[set].Resource(id, nil).Set(field, v)
			}
		}

		for _, rel := range typ.Rels {
			if rel.FromName == field {
				if rel.ToOne {
					t.sets[set].Resource(id, nil).SetToOne(field, v.(string))
				} else {
					t.sets[set].Resource(id, nil).SetToMany(field, v.([]string))
				}
			}
		}
	case id == "" && field == "id":
		// Create a resource
		typ := t.sets[set].Type
		t.sets[set].Add(makeSoftResource(typ, v.(string), map[string]interface{}{}))
	case id != "" && field == "id" && v.(string) == "":
		// Delete a resource
		t.sets[set].Remove(id)
	}
}

func (t *Tx) opInsert(set, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s += %v\n", set, id, field, v)
	curr := reflect.ValueOf(t.sets[set].Resource(id, nil).GetToMany(field))
	curr = reflect.Append(curr, reflect.ValueOf(v))

	typ := t.sets[set].Type
	for _, attr := range typ.Attrs {
		if attr.Name == field {
			t.sets[set].Resource(id, nil).Set(field, v)
		}
	}

	for _, rel := range typ.Rels {
		if rel.FromName == field {
			if rel.ToOne {
				t.sets[set].Resource(id, nil).SetToOne(field, curr.Interface().(string))
			} else {
				t.sets[set].Resource(id, nil).SetToMany(field, curr.Interface().([]string))
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
