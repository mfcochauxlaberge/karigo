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
	// 	Name: "created",
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
			"created": true,
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
			"created": true,
			"attrs": []string{
				"0_sets_name",
				"0_sets_version",
				"0_sets_created",
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
			"created": true,
			"attrs": []string{
				"0_attrs_name",
				"0_attrs_type",
				"0_attrs_null",
				"0_attrs_created",
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
			"created": true,
			"attrs": []string{
				"0_rels_name",
				"0_rels_to-one",
				"0_rels_created",
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
			"created": true,
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
	// 	Name: "created",
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
			"name":    "value",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_meta",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_name",
		map[string]interface{}{
			"name":    "name",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_sets",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_version",
		map[string]interface{}{
			"name":    "version",
			"type":    "int",
			"null":    false,
			"created": true,
			"set":     "0_sets",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_created",
		map[string]interface{}{
			"name":    "created",
			"type":    "bool",
			"null":    false,
			"created": true,
			"set":     "0_sets",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_attrs_name",
		map[string]interface{}{
			"name":    "name",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_attrs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_attrs_type",
		map[string]interface{}{
			"name":    "type",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_attrs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_attrs_null",
		map[string]interface{}{
			"name":    "null",
			"type":    "bool",
			"null":    false,
			"created": true,
			"set":     "0_attrs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_attrs_created",
		map[string]interface{}{
			"name":    "created",
			"type":    "bool",
			"null":    false,
			"created": true,
			"set":     "0_attrs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_rels_name",
		map[string]interface{}{
			"name":    "name",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_rels",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_rels_to-one",
		map[string]interface{}{
			"name":    "to-one",
			"type":    "bool",
			"null":    false,
			"created": true,
			"set":     "0_rels",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_rels_created",
		map[string]interface{}{
			"name":    "created",
			"type":    "bool",
			"null":    false,
			"created": true,
			"set":     "0_rels",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_funcs_func",
		map[string]interface{}{
			"name":    "func",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_funcs",
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
	// 	Name: "created",
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
			"created": true,
			"inverse": "0_attrs_set",
			"set":     "0_sets",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_rels",
		map[string]interface{}{
			"name":    "rels",
			"to-one":  false,
			"created": true,
			"inverse": "0_rels_set",
			"set":     "0_sets",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_get_func",
		map[string]interface{}{
			"name":    "get_func",
			"to-one":  true,
			"created": true,
			"inverse": "", // Inverse?
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_create_func",
		map[string]interface{}{
			"name":    "create_func",
			"to-one":  true,
			"created": true,
			"inverse": "", // Inverse?
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_update_func",
		map[string]interface{}{
			"name":    "update_func",
			"to-one":  true,
			"created": true,
			"inverse": "", // Inverse?
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_sets_delete_func",
		map[string]interface{}{
			"name":    "delete_func",
			"to-one":  true,
			"created": true,
			"inverse": "", // Inverse?
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_attrs_set",
		map[string]interface{}{
			"name":    "set",
			"to-one":  true,
			"created": true,
			"inverse": "0_sets_attrs",
			"set":     "0_attrs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_rels_inverse",
		map[string]interface{}{
			"name":    "inverse",
			"to-one":  true,
			"created": true,
			"inverse": "0_rels_inverse",
			"set":     "0_rels",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_rels_set",
		map[string]interface{}{
			"name":    "set",
			"to-one":  true,
			"created": true,
			"inverse": "0_sets_rels",
			"set":     "0_rels",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_funcs_get_func",
		map[string]interface{}{
			"name":    "get_func",
			"to-one":  true,
			"created": true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_funcs_create_func",
		map[string]interface{}{
			"name":    "create_func",
			"to-one":  true,
			"created": true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_funcs_update_func",
		map[string]interface{}{
			"name":    "update_func",
			"to-one":  true,
			"created": true,
			"inverse": "0_sets_rels",
			"set":     "0_funcs",
		},
	))
	m.data["0_sets"].Add(set.NewRecord(
		"0_funcs_delete_func",
		map[string]interface{}{
			"name":    "delete_func",
			"to-one":  true,
			"created": true,
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

// Begin ...
func (m *Source) Begin() (karigo.SourceTx, error) {
	// m.Lock()
	// defer m.sUnlock()

	return nil, nil
}

// Apply ...
func (m *Source) Apply(ops []karigo.Op) error {
	m.Lock()
	defer m.Unlock()

	// TODO Create a mtx, apply the ops, commit it

	return nil
}

func (m *Source) opSet(set, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s (%v)\n", set, id, field, v)

	// if id != "" && field != "id" {
	// 	m.data[set].data[id][field] = v
	// 	m.data[set].set(id, field, v)
	// }

	// if id == "" && field == "id" {
	// 	m.data[set][v.(string)] = map[string]interface{}{}
	// 	// fmt.Printf("New entry inserted.\n")
	// } else if strings.HasPrefix(set, "0_") && field == "created" {
	// 	// If a set, attribute, or relationship is marked as created, create it.
	// 	switch field {
	// 	case "created":
	// 		switch set {
	// 		case "0_sets":
	// 			name := m.data["0_sets"][id]["name"].(string)
	// 			m.data[name] = map[string]map[string]interface{}{}
	// 		case "0_attrs":
	// 			name := m.data["0_attrs"][id]["name"].(string)
	// 			typ := m.data["0_attrs"][id]["type"].(string)
	// 			set := m.data["0_attrs"][id]["set"].(string)
	// 			for id2 := range m.data[set] {
	// 				fmt.Printf("Created: %s %s %s\n", set, id2, name)
	// 				m.data[set][id2][name] = zeroVal(typ)
	// 			}
	// 		case "0_rels":
	// 			name := m.data["0_rels"][id]["name"].(string)
	// 			toOne := m.data["0_rels"][id]["to-one"].(bool)
	// 			set := m.data["0_rels"][id]["set"].(string)
	// 			for id2 := range m.data[set] {
	// 				if toOne {
	// 					m.data[set][id2][name] = ""
	// 				} else {
	// 					m.data[set][id2][name] = []string{}
	// 				}
	// 			}
	// 		}
	// 	}
	// 	// fmt.Printf("created=true, new thing created.\n")
	// } else {
	// 	// if _, ok := m.data[set]; !ok {
	// 	// 	fmt.Printf("Set %s does not exist.\n", set)
	// 	// }
	// 	// if _, ok := m.data[set][id]; !ok {
	// 	// 	fmt.Printf("ID %s does not exist.\n", id)
	// 	// }
	// 	// if _, ok := m.data[set][id][field]; !ok {
	// 	// 	fmt.Printf("Field %s does not exist.\n", field)
	// 	// }
	// 	m.data[set][id][field] = v
	// }
}
