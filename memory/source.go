package memory

import (
	"sync"

	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/memory/internal/set"
)

// Source ...
type Source struct {
	ID       string
	Location string

	// schema *jsonapi.Schema
	data map[string]*set.Set

	// oldSchema *jsonapi.Schema
	// oldData map[string]set.Set

	sync.Mutex
}

// Reset ...
func (m *Source) Reset() error {
	m.Lock()
	defer m.Unlock()

	m.data = map[string]*set.Set{}

	// m.schema = &jsonapi.Schema{}

	// 0_meta
	// typ := jsonapi.Type{
	// 	Name: "0_meta",
	// }
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "value",
	// 	Type: jsonapi.AttrTypeString,
	// 	Null: false,
	// })
	// m.schema.AddType(typ)

	m.data["0_meta"] = &set.Set{}

	// 0_sets
	// typ = jsonapi.Type{
	// 	Name: "0_sets",
	// }
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "name",
	// 	Type: jsonapi.AttrTypeString,
	// 	Null: false,
	// })
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "version",
	// 	Type: jsonapi.AttrTypeUint,
	// 	Null: false,
	// })
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "active",
	// 	Type: jsonapi.AttrTypeBool,
	// 	Null: false,
	// })
	// typ.AddRel(jsonapi.Rel{
	// 	Name:         "attrs",
	// 	Type:         "0_attrs",
	// 	ToOne:        false,
	// 	InverseName:  "set",
	// 	InverseType:  "0_sets",
	// 	InverseToOne: true,
	// })
	// typ.AddRel(jsonapi.Rel{
	// 	Name:         "rels",
	// 	Type:         "0_rels",
	// 	ToOne:        false,
	// 	InverseName:  "set",
	// 	InverseType:  "0_sets",
	// 	InverseToOne: true,
	// })
	// m.schema.AddType(typ)

	m.data["0_sets"] = &set.Set{}
	m.data["0_sets"].Add(set.NewRecord(
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
	m.data["0_sets"].Add(set.NewRecord(
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
	m.data["0_sets"].Add(set.NewRecord(
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
	m.data["0_sets"].Add(set.NewRecord(
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
	m.data["0_sets"].Add(set.NewRecord(
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

	// // 0_attrs
	// typ = jsonapi.Type{
	// 	Name: "0_attrs",
	// }
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "name",
	// 	Type: jsonapi.AttrTypeString,
	// 	Null: false,
	// })
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "type",
	// 	Type: jsonapi.AttrTypeUint,
	// 	Null: false,
	// })
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "null",
	// 	Type: jsonapi.AttrTypeBool,
	// 	Null: false,
	// })
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "active",
	// 	Type: jsonapi.AttrTypeBool,
	// 	Null: false,
	// })
	// typ.AddRel(jsonapi.Rel{
	// 	Name:         "set",
	// 	Type:         "0_sets",
	// 	ToOne:        true,
	// 	InverseName:  "attrs",
	// 	InverseType:  "0_attrs",
	// 	InverseToOne: false,
	// })
	// m.schema.AddType(typ)

	m.data["0_attrs"] = &set.Set{}
	m.data["0_sets"].Add(set.NewRecord(
		"0_meta_value",
		map[string]interface{}{
			"name":   "value",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_meta",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_name",
		map[string]interface{}{
			"name":   "name",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_sets",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_version",
		map[string]interface{}{
			"name":   "version",
			"type":   "int",
			"null":   false,
			"active": true,
			"set":    "0_sets",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_active",
		map[string]interface{}{
			"name":   "active",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_sets",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_attrs_name",
		map[string]interface{}{
			"name":   "name",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_attrs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_attrs_type",
		map[string]interface{}{
			"name":   "type",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_attrs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_attrs_null",
		map[string]interface{}{
			"name":   "null",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_attrs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_attrs_active",
		map[string]interface{}{
			"name":   "active",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_attrs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_rels_name",
		map[string]interface{}{
			"name":   "name",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_rels",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_rels_to-one",
		map[string]interface{}{
			"name":   "to-one",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_rels",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_rels_active",
		map[string]interface{}{
			"name":   "active",
			"type":   "bool",
			"null":   false,
			"active": true,
			"set":    "0_rels",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_funcs_func",
		map[string]interface{}{
			"name":   "func",
			"type":   "string",
			"null":   false,
			"active": true,
			"set":    "0_funcs",
		},
	))

	// // 0_rels
	// typ = jsonapi.Type{
	// 	Name: "0_rels",
	// }
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "name",
	// 	Type: jsonapi.AttrTypeString,
	// 	Null: false,
	// })
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "to-one",
	// 	Type: jsonapi.AttrTypeBool,
	// 	Null: false,
	// })
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "active",
	// 	Type: jsonapi.AttrTypeBool,
	// 	Null: false,
	// })
	// typ.AddRel(jsonapi.Rel{
	// 	Name:         "inverse",
	// 	Type:         "0_rels",
	// 	ToOne:        true,
	// 	InverseName:  "inverse",
	// 	InverseType:  "0_rels",
	// 	InverseToOne: true,
	// })
	// typ.AddRel(jsonapi.Rel{
	// 	Name:         "set",
	// 	Type:         "0_sets",
	// 	ToOne:        true,
	// 	InverseName:  "rels",
	// 	InverseType:  "0_rels",
	// 	InverseToOne: false,
	// })
	// m.schema.AddType(typ)

	m.data["0_rels"] = &set.Set{}
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_attrs",
		map[string]interface{}{
			"name":    "attrs",
			"to-one":  false,
			"active":  true,
			"inverse": "0_attrs_set",
			"set":     "0_sets",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_rels",
		map[string]interface{}{
			"name":    "rels",
			"to-one":  false,
			"active":  true,
			"inverse": "0_rels_set",
			"set":     "0_sets",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_get_func",
		map[string]interface{}{
			"name":    "get_func",
			"to-one":  true,
			"active":  true,
			"inverse": "", // TODO Inverse?
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_create_func",
		map[string]interface{}{
			"name":    "create_func",
			"to-one":  true,
			"active":  true,
			"inverse": "", // TODO Inverse?
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_update_func",
		map[string]interface{}{
			"name":    "update_func",
			"to-one":  true,
			"active":  true,
			"inverse": "", // TODO Inverse?
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_delete_func",
		map[string]interface{}{
			"name":    "delete_func",
			"to-one":  true,
			"active":  true,
			"inverse": "", // TODO Inverse?
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_attrs_set",
		map[string]interface{}{
			"name":    "set",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_attrs",
			"set":     "0_attrs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_rels_inverse",
		map[string]interface{}{
			"name":    "inverse",
			"to-one":  true,
			"active":  true,
			"inverse": "0_rels_inverse",
			"set":     "0_rels",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_rels_set",
		map[string]interface{}{
			"name":    "set",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_rels",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_funcs_get_func",
		map[string]interface{}{
			"name":    "get_func",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_funcs_create_func",
		map[string]interface{}{
			"name":    "create_func",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_funcs_update_func",
		map[string]interface{}{
			"name":    "update_func",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_funcs_delete_func",
		map[string]interface{}{
			"name":    "delete_func",
			"to-one":  true,
			"active":  true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))

	// // 0_funcs
	// typ = jsonapi.Type{
	// 	Name: "funcs",
	// }
	// typ.AddAttr(jsonapi.Attr{
	// 	Name: "func",
	// 	Type: jsonapi.AttrTypeString,
	// 	Null: false,
	// })
	// m.schema.AddType(typ)

	m.data["0_funcs"] = &set.Set{}
	m.data["0_sets"].Add(set.NewRecord(
		"_not_implemented",
		map[string]interface{}{
			"func": `func(snap *Snapshot) error {
				snap.Fail(ErrNotImplemented)
			}`,
		},
	))

	// errs := m.schema.Check()
	// if len(errs) > 0 {
	// 	return errs[0]
	// }

	return nil
}

// Resource ...
func (m *Source) Resource(qry karigo.QueryRes) (jsonapi.Resource, error) {
	m.Lock()
	defer m.Unlock()

	// Get resource
	res := m.data[qry.Set].Resource(qry.ID, qry.Fields)

	return res, nil
}

// Collection ...
func (m *Source) Collection(qry karigo.QueryCol) ([]jsonapi.Resource, error) {
	m.Lock()
	defer m.Unlock()

	// BelongsToFilter
	var ids []string
	if qry.BelongsToFilter.ID != "" {
		res := m.data[qry.BelongsToFilter.Type].Resource(qry.BelongsToFilter.ID, []string{})
		ids = res.GetToMany(qry.BelongsToFilter.Name)
	}

	// Get all records from the given set
	recs := m.data[qry.Set].Collection(
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
func (m *Source) Apply(ops []karigo.Op) error {
	m.Lock()
	defer m.Unlock()

	for _, op := range ops {
		switch op.Op {
		case karigo.OpSet:
			m.opSet(op.Key.Set, op.Key.ID, op.Key.Field, op.Value)
		}
	}

	return nil
}

func (m *Source) opSet(setname, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s (%v)\n", setname, id, field, v)

	if id != "" && field != "id" {
		// Set a field
		m.data[setname].Set(id, field, v)
	} else if id == "" && field == "id" {
		// Create a resource

		// Before, check whether it's a new set because then it
		// requires a new entry in m.data.
		if setname == "0_sets" {
			m.data[v.(string)] = &set.Set{}
		}

		m.data[setname].Add(set.NewRecord(v.(string), map[string]interface{}{}))
	} else if id != "" && field == "id" {
		// Delete a resource

		if v.(string) == "" {
			m.data[setname].Del(id)
		}
	} else {
		// Should not happen
		// TODO Should this code path be reported?
	}
}
