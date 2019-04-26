package memory

import (
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/mfcochauxlaberge/karigo"
)

// Source ...
type Source struct {
	ID       string
	Location string

	schema *jsonapi.Schema
	data   map[string]set

	oldSchema *jsonapi.Schema
	oldData   map[string]set

	sync.Mutex
}

// Reset ...
func (m *Source) Reset() error {
	m.Lock()
	defer m.Unlock()

	m.schema = &jsonapi.Schema{}

	// 0_meta
	typ := jsonapi.Type{
		Name: "0_meta",
	}
	typ.AddAttr(jsonapi.Attr{
		Name: "value",
		Type: jsonapi.AttrTypeString,
		Null: false,
	})
	m.schema.AddType(typ)

	m.data["0_meta"] = set{
		data: []record{},
	}

	// 0_sets
	typ = jsonapi.Type{
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
		Name: "created",
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
	m.schema.AddType(typ)

	m.data["0_sets"] = set{
		data: []record{
			record{
				id: "0_meta",
				vals: map[string]interface{}{
					"name":    "0_meta",
					"version": 0,
					"created": true,
					"attrs": []string{
						"0_meta_value",
					},
				},
			},
			record{
				id: "0_sets",
				vals: map[string]interface{}{
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
			},
			record{
				id: "0_attrs",
				vals: map[string]interface{}{
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
			},
			record{
				id: "0_rels",
				vals: map[string]interface{}{
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
			},
			record{
				id: "0_funcs",
				vals: map[string]interface{}{
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
			},
		},
	}

	// 0_attrs
	typ = jsonapi.Type{
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
		Name: "created",
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
	m.schema.AddType(typ)

	m.data["0_attrs"] = set{
		data: []record{
			record{
				id: "0_meta_value",
				vals: map[string]interface{}{
					"name":    "value",
					"type":    "string",
					"null":    false,
					"created": true,
					"set":     "0_meta",
				},
			},
			record{
				id: "0_sets_name",
				vals: map[string]interface{}{
					"name":    "name",
					"type":    "string",
					"null":    false,
					"created": true,
					"set":     "0_sets",
				},
			},
			record{
				id: "0_sets_version",
				vals: map[string]interface{}{
					"name":    "version",
					"type":    "int",
					"null":    false,
					"created": true,
					"set":     "0_sets",
				},
			},
			record{
				id: "0_sets_created",
				vals: map[string]interface{}{
					"name":    "created",
					"type":    "bool",
					"null":    false,
					"created": true,
					"set":     "0_sets",
				},
			},
			record{
				id: "0_attrs_name",
				vals: map[string]interface{}{
					"name":    "name",
					"type":    "string",
					"null":    false,
					"created": true,
					"set":     "0_attrs",
				},
			},
			record{
				id: "0_attrs_type",
				vals: map[string]interface{}{
					"name":    "type",
					"type":    "string",
					"null":    false,
					"created": true,
					"set":     "0_attrs",
				},
			},
			record{
				id: "0_attrs_null",
				vals: map[string]interface{}{
					"name":    "null",
					"type":    "bool",
					"null":    false,
					"created": true,
					"set":     "0_attrs",
				},
			},
			record{
				id: "0_attrs_created",
				vals: map[string]interface{}{
					"name":    "created",
					"type":    "bool",
					"null":    false,
					"created": true,
					"set":     "0_attrs",
				},
			},
			record{
				id: "0_rels_name",
				vals: map[string]interface{}{
					"name":    "name",
					"type":    "string",
					"null":    false,
					"created": true,
					"set":     "0_rels",
				},
			},
			record{
				id: "0_rels_to-one",
				vals: map[string]interface{}{
					"name":    "to-one",
					"type":    "bool",
					"null":    false,
					"created": true,
					"set":     "0_rels",
				},
			},
			record{
				id: "0_rels_created",
				vals: map[string]interface{}{
					"name":    "created",
					"type":    "bool",
					"null":    false,
					"created": true,
					"set":     "0_rels",
				},
			},
			record{
				id: "0_funcs_func",
				vals: map[string]interface{}{
					"name":    "func",
					"type":    "string",
					"null":    false,
					"created": true,
					"set":     "0_funcs",
				},
			},
		},
	}

	// 0_rels
	typ = jsonapi.Type{
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
		Name: "created",
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
	m.schema.AddType(typ)

	m.data["0_rels"] = set{
		data: []record{
			record{
				id: "0_sets_attrs",
				vals: map[string]interface{}{
					"name":    "attrs",
					"to-one":  false,
					"created": true,
					"inverse": "0_attrs_set",
					"set":     "0_sets",
				},
			},
			record{
				id: "0_sets_rels",
				vals: map[string]interface{}{
					"name":    "rels",
					"to-one":  false,
					"created": true,
					"inverse": "0_rels_set",
					"set":     "0_sets",
				},
			},
			record{
				id: "0_sets_get_func",
				vals: map[string]interface{}{
					"name":    "get_func",
					"to-one":  true,
					"created": true,
					"inverse": "", // Inverse?
					"set":     "0_funcs",
				},
			},
			record{
				id: "0_sets_create_func",
				vals: map[string]interface{}{
					"name":    "create_func",
					"to-one":  true,
					"created": true,
					"inverse": "", // Inverse?
					"set":     "0_funcs",
				},
			},
			record{
				id: "0_sets_update_func",
				vals: map[string]interface{}{
					"name":    "update_func",
					"to-one":  true,
					"created": true,
					"inverse": "", // Inverse?
					"set":     "0_funcs",
				},
			},
			record{
				id: "0_sets_delete_func",
				vals: map[string]interface{}{
					"name":    "delete_func",
					"to-one":  true,
					"created": true,
					"inverse": "", // Inverse?
					"set":     "0_funcs",
				},
			},
			record{
				id: "0_attrs_set",
				vals: map[string]interface{}{
					"name":    "set",
					"to-one":  true,
					"created": true,
					"inverse": "0_sets_attrs",
					"set":     "0_attrs",
				},
			},
			record{
				id: "0_rels_inverse",
				vals: map[string]interface{}{
					"name":    "inverse",
					"to-one":  true,
					"created": true,
					"inverse": "0_rels_inverse",
					"set":     "0_rels",
				},
			},
			record{
				id: "0_rels_set",
				vals: map[string]interface{}{
					"name":    "set",
					"to-one":  true,
					"created": true,
					"inverse": "0_sets_rels",
					"set":     "0_rels",
				},
			},
			record{
				id: "0_funcs_get_func",
				vals: map[string]interface{}{
					"name":    "get_func",
					"to-one":  true,
					"created": true,
					"inverse": "0_sets_rels",
					"set":     "0_funcs",
				},
			},
			record{
				id: "0_funcs_create_func",
				vals: map[string]interface{}{
					"name":    "create_func",
					"to-one":  true,
					"created": true,
					"inverse": "0_sets_rels",
					"set":     "0_funcs",
				},
			},
			record{
				id: "0_funcs_update_func",
				vals: map[string]interface{}{
					"name":    "update_func",
					"to-one":  true,
					"created": true,
					"inverse": "0_sets_rels",
					"set":     "0_funcs",
				},
			},
			record{
				id: "0_funcs_delete_func",
				vals: map[string]interface{}{
					"name":    "delete_func",
					"to-one":  true,
					"created": true,
					"inverse": "0_sets_rels",
					"set":     "0_funcs",
				},
			},
		},
	}

	// 0_funcs
	typ = jsonapi.Type{
		Name: "funcs",
	}
	typ.AddAttr(jsonapi.Attr{
		Name: "func",
		Type: jsonapi.AttrTypeString,
		Null: false,
	})
	m.schema.AddType(typ)

	// TODO Add missing functions
	m.data["0_funcs"] = set{
		data: []record{
			record{
				id: "0_meta",
				vals: map[string]interface{}{
					"func": `func(snap *Snapshot) error {
						snap.Fail(ErrNotImplemented)
					}`,
				},
			},
			record{
				id: "0_sets",
				vals: map[string]interface{}{
					"func": `func(snap *Snapshot) error {
						snap.Fail(ErrNotImplemented)
					}`,
				},
			},
			record{
				id: "0_attrs",
				vals: map[string]interface{}{
					"func": `func(snap *Snapshot) error {
						snap.Fail(ErrNotImplemented)
					}`,
				},
			},
			record{
				id: "0_rels",
				vals: map[string]interface{}{
					"func": `func(snap *Snapshot) error {
						snap.Fail(ErrNotImplemented)
					}`,
				},
			},
			record{
				id: "0_get-funcs",
				vals: map[string]interface{}{
					"func": `func(snap *Snapshot) error {
						snap.Fail(ErrNotImplemented)
					}`,
				},
			},
			record{
				id: "0_create-funcs",
				vals: map[string]interface{}{
					"func": `func(snap *Snapshot) error {
						snap.Fail(ErrNotImplemented)
					}`,
				},
			},
			record{
				id: "0_update-funcs",
				vals: map[string]interface{}{
					"func": `func(snap *Snapshot) error {
						snap.Fail(ErrNotImplemented)
					}`,
				},
			},
			record{
				id: "0_delete-funcs",
				vals: map[string]interface{}{
					"func": `func(snap *Snapshot) error {
						snap.Fail(ErrNotImplemented)
					}`,
				},
			},
		},
	}

	errs := m.schema.Check()
	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

// Resource ...
func (m *Source) Resource(qry karigo.QueryRes) (jsonapi.Resource, error) {
	m.Lock()
	defer m.Unlock()

	// Get resource
	var rec record
	for i := range m.data[qry.Set].data {
		if m.data[qry.Set].data[i].id == qry.ID {
			rec = m.data[qry.Set].data[i]
		}
	}

	typ, _ := m.schema.GetType(qry.Set)
	res := jsonapi.NewSoftResource(typ, rec.vals)

	// Filter fields
	for field := range rec.vals {
		for _, f := range qry.Fields {
			if field == f {
				res.Set(field, rec.vals[field])
				break
			}
		}
	}

	return res, nil
}

// Collection ...
func (m *Source) Collection(qry karigo.QueryCol) ([]jsonapi.Resource, error) {
	m.Lock()
	defer m.Unlock()

	// Get all records from the given set
	recs := m.data[qry.Set]

	// BelongsToFilter
	if qry.BelongsToFilter.ID != "" {
		resqry := karigo.QueryRes{
			Set:    qry.BelongsToFilter.Type,
			ID:     qry.BelongsToFilter.ID,
			Fields: []string{qry.BelongsToFilter.Name},
		}
		res, err := m.Resource(resqry)
		if err != nil {
			return nil, err
		}
		kept := set{}
		ids := res.GetToMany(qry.BelongsToFilter.Name)
		for i := range recs.data {
			keep := false
			for i := range ids {
				if recs.data[i].id == ids[i] {
					keep = true
					break
				}
			}
			if keep {
				kept.data = append(kept.data, recs.data[i])
			}
		}
		recs = kept
	}

	// TODO Filter

	// Sort
	recs.sort = qry.Sort
	sort.Sort(&recs)

	// Pagination
	if qry.PageSize == 0 {
		recs = set{}
	} else {
		skip := qry.PageNumber * qry.PageSize
		if skip >= len(recs.data) {
			recs = set{}
		} else {
			page := set{}
			for i := skip; i < len(recs.data) || i < qry.PageSize; i++ {
				page.data = append(page.data, recs.data[i])
			}
			recs = page
		}
	}

	// Fields
	for i := range recs.data {
		for k := range recs.data[i].vals {
			found := false
			for _, f := range qry.Fields {
				if k == f {
					found = true
					break
				}
			}
			if !found {
				delete(recs.data[i].vals, k)
			}
		}
	}

	return nil, nil
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
	// 	m.data[set][id][field] = v
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

type record struct {
	schema *jsonapi.Schema
	id     string
	vals   map[string]interface{}
}

type set struct {
	data []record
	sort []string
}

// Len ...
func (s *set) Len() int { return len(s.data) }

// Swap ...
func (s *set) Swap(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] }

// Less ...
func (s *set) Less(i, j int) bool {
	less := false

	for _, r := range s.sort {
		inverse := false
		if strings.HasPrefix(r, "-") {
			r = r[1:]
			inverse = true
		}

		switch v := s.data[i].vals[r].(type) {
		case string:
			if v == s.data[j].vals[r].(string) {
				continue
			}
			if inverse {
				return v > s.data[j].vals[r].(string)
			}
			return v < s.data[j].vals[r].(string)
		case int:
			if v == s.data[j].vals[r].(int) {
				continue
			}
			if inverse {
				return v > s.data[j].vals[r].(int)
			}
			return v < s.data[j].vals[r].(int)
		case int8:
			if v == s.data[j].vals[r].(int8) {
				continue
			}
			if inverse {
				return v > s.data[j].vals[r].(int8)
			}
			return v < s.data[j].vals[r].(int8)
		case int16:
			if v == s.data[j].vals[r].(int16) {
				continue
			}
			if inverse {
				return v > s.data[j].vals[r].(int16)
			}
			return v < s.data[j].vals[r].(int16)
		case int32:
			if v == s.data[j].vals[r].(int32) {
				continue
			}
			if inverse {
				return v > s.data[j].vals[r].(int32)
			}
			return v < s.data[j].vals[r].(int32)
		case int64:
			if v == s.data[j].vals[r].(int64) {
				continue
			}
			if inverse {
				return v > s.data[j].vals[r].(int64)
			}
			return v < s.data[j].vals[r].(int64)
		case uint:
			if v == s.data[j].vals[r].(uint) {
				continue
			}
			if inverse {
				return v > s.data[j].vals[r].(uint)
			}
			return v < s.data[j].vals[r].(uint)
		case uint8:
			if v == s.data[j].vals[r].(uint8) {
				continue
			}
			if inverse {
				return v > s.data[j].vals[r].(uint8)
			}
			return v < s.data[j].vals[r].(uint8)
		case uint16:
			if v == s.data[j].vals[r].(uint16) {
				continue
			}
			if inverse {
				return v > s.data[j].vals[r].(uint16)
			}
			return v < s.data[j].vals[r].(uint16)
		case uint32:
			if v == s.data[j].vals[r].(uint32) {
				continue
			}
			if inverse {
				return v > s.data[j].vals[r].(uint32)
			}
			return v < s.data[j].vals[r].(uint32)
		case bool:
			if v == s.data[j].vals[r].(bool) {
				continue
			}
			if inverse {
				return v
			}
			return !v
		case time.Time:
			if v.Equal(s.data[j].vals[r].(time.Time)) {
				continue
			}
			if inverse {
				return v.After(s.data[j].vals[r].(time.Time))
			}
			return v.Before(s.data[j].vals[r].(time.Time))
		case *string:
			if *v == *(s.data[j].vals[r].(*string)) {
				continue
			}
			if inverse {
				return *v > *(s.data[j].vals[r].(*string))
			}
			return *v < *(s.data[j].vals[r].(*string))
		case *int:
			if *v == *(s.data[j].vals[r].(*int)) {
				continue
			}
			if inverse {
				return *v > *(s.data[j].vals[r].(*int))
			}
			return *v < *(s.data[j].vals[r].(*int))
		case *int8:
			if *v == *(s.data[j].vals[r].(*int8)) {
				continue
			}
			if inverse {
				return *v > *(s.data[j].vals[r].(*int8))
			}
			return *v < *(s.data[j].vals[r].(*int8))
		case *int16:
			if *v == *(s.data[j].vals[r].(*int16)) {
				continue
			}
			if inverse {
				return *v > *(s.data[j].vals[r].(*int16))
			}
			return *v < *(s.data[j].vals[r].(*int16))
		case *int32:
			if *v == *(s.data[j].vals[r].(*int32)) {
				continue
			}
			if inverse {
				return *v > *(s.data[j].vals[r].(*int32))
			}
			return *v < *(s.data[j].vals[r].(*int32))
		case *int64:
			if *v == *(s.data[j].vals[r].(*int64)) {
				continue
			}
			if inverse {
				return *v > *(s.data[j].vals[r].(*int64))
			}
			return *v < *(s.data[j].vals[r].(*int64))
		case *uint:
			if *v == *(s.data[j].vals[r].(*uint)) {
				continue
			}
			if inverse {
				return *v > *(s.data[j].vals[r].(*uint))
			}
			return *v < *(s.data[j].vals[r].(*uint))
		case *uint8:
			if *v == *(s.data[j].vals[r].(*uint8)) {
				continue
			}
			if inverse {
				return *v > *(s.data[j].vals[r].(*uint8))
			}
			return *v < *(s.data[j].vals[r].(*uint8))
		case *uint16:
			if *v == *(s.data[j].vals[r].(*uint16)) {
				continue
			}
			if inverse {
				return *v > *(s.data[j].vals[r].(*uint16))
			}
			return *v < *(s.data[j].vals[r].(*uint16))
		case *uint32:
			if *v == *(s.data[j].vals[r].(*uint32)) {
				continue
			}
			if inverse {
				return *v > *(s.data[j].vals[r].(*uint32))
			}
			return *v < *(s.data[j].vals[r].(*uint32))
		case *bool:
			if *v == *(s.data[j].vals[r].(*bool)) {
				continue
			}
			if inverse {
				return *v
			}
			return !*v
		case *time.Time:
			if v.Equal(*(s.data[j].vals[r].(*time.Time))) {
				continue
			}
			if inverse {
				return v.After(*(s.data[j].vals[r].(*time.Time)))
			}
			return v.Before(*(s.data[j].vals[r].(*time.Time)))
		}
	}

	return less
}
