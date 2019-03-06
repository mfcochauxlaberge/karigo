package sources

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/mfcochauxlaberge/karigo"

	"github.com/mitchellh/copystructure"
)

// Memory ...
type Memory struct {
	ID       string
	Location string

	data   map[string]map[string]map[string]interface{}
	fields map[string][]string

	oldData   map[string]map[string]map[string]interface{}
	oldFields map[string][]string

	sync.Mutex
}

// Reset ...
func (m *Memory) Reset() error {
	m.Lock()
	defer m.Unlock()

	m.data = map[string]map[string]map[string]interface{}{}

	m.data["0_meta"] = map[string]map[string]interface{}{}

	m.data["0_sets"] = map[string]map[string]interface{}{
		"0_meta": map[string]interface{}{
			"name":    "0_meta",
			"version": 0,
			"created": true,
			"attrs": []string{
				"0_meta_value",
			},
		},
		"0_sets": map[string]interface{}{
			"name":    "0_sets",
			"version": 0,
			"created": true,
			"attrs": []string{
				"0_sets_name",
				"0_sets_version",
			},
			"rels": []string{
				"0_sets_fields",
			},
		},
		"0_attrs": map[string]interface{}{
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
		"0_rels": map[string]interface{}{
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
		"0_get-funcs": map[string]interface{}{
			"name":    "0_get-funcs",
			"version": 0,
			"created": true,
			"attrs": []string{
				"0_get-funcs_func",
			},
			"rels": []string{
				"0_get-funcs_set",
			},
		},
		"0_create-funcs": map[string]interface{}{
			"name":    "0_create-funcs",
			"version": 0,
			"created": true,
			"attrs": []string{
				"0_create-funcs_func",
			},
			"rels": []string{
				"0_create-funcs_set",
			},
		},
		"0_update-funcs": map[string]interface{}{
			"name":    "0_update-funcs",
			"version": 0,
			"created": true,
			"attrs": []string{
				"0_update-funcs_func",
			},
			"rels": []string{
				"0_update-funcs_set",
			},
		},
		"0_delete-funcs": map[string]interface{}{
			"name":    "0_delete-funcs",
			"version": 0,
			"created": true,
			"attrs": []string{
				"0_delete-funcs_func",
			},
			"rels": []string{
				"0_delete-funcs_set",
			},
		},
	}

	m.data["0_attrs"] = map[string]map[string]interface{}{
		"0_meta_value": map[string]interface{}{
			"name":    "value",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_meta",
		},
		"0_sets_version": map[string]interface{}{
			"name":    "version",
			"type":    "int",
			"null":    false,
			"created": true,
			"set":     "0_sets",
		},
		"0_sets_created": map[string]interface{}{
			"name":    "created",
			"type":    "bool",
			"null":    false,
			"created": true,
			"set":     "0_sets",
		},
		"0_attrs_name": map[string]interface{}{
			"name":    "name",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_attrs",
		},
		"0_attrs_type": map[string]interface{}{
			"name":    "type",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_attrs",
		},
		"0_attrs_null": map[string]interface{}{
			"name":    "null",
			"type":    "bool",
			"null":    false,
			"created": true,
			"set":     "0_attrs",
		},
		"0_attrs_created": map[string]interface{}{
			"name":    "created",
			"type":    "bool",
			"null":    false,
			"created": true,
			"set":     "0_attrs",
		},
		"0_rels_name": map[string]interface{}{
			"name":    "name",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_rels",
		},
		"0_rels_to-one": map[string]interface{}{
			"name":    "to-one",
			"type":    "bool",
			"null":    false,
			"created": true,
			"set":     "0_rels",
		},
		"0_rels_created": map[string]interface{}{
			"name":    "created",
			"type":    "bool",
			"null":    false,
			"created": true,
			"set":     "0_rels",
		},
		"0_get-funcs_func": map[string]interface{}{
			"name":    "func",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_get-funcs",
		},
		"0_create-funcs_func": map[string]interface{}{
			"name":    "func",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_create-funcs",
		},
		"0_update-funcs_func": map[string]interface{}{
			"name":    "func",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_update-funcs",
		},
		"0_delete-funcs_func": map[string]interface{}{
			"name":    "func",
			"type":    "string",
			"null":    false,
			"created": true,
			"set":     "0_delete-funcs",
		},
	}

	m.data["0_rels"] = map[string]map[string]interface{}{
		"0_sets_attrs": map[string]interface{}{
			"name":    "attrs",
			"to-one":  false,
			"created": true,
			"inverse": "0_attrs_set",
			"set":     "0_sets",
		},
		"0_sets_rels": map[string]interface{}{
			"name":    "rels",
			"to-one":  false,
			"created": true,
			"inverse": "0_rels_set",
			"set":     "0_sets",
		},
		"0_attrs_set": map[string]interface{}{
			"name":    "set",
			"to-one":  true,
			"created": true,
			"inverse": "0_sets_attrs",
			"set":     "0_attrs",
		},
		"0_rels_inverse": map[string]interface{}{
			"name":    "inverse",
			"to-one":  true,
			"created": true,
			"inverse": "0_rels_inverse",
			"set":     "0_rels",
		},
		"0_rels_set": map[string]interface{}{
			"name":    "set",
			"to-one":  true,
			"created": true,
			"inverse": "0_sets_rels",
			"set":     "0_rels",
		},
	}

	m.data["0_get-funcs"] = map[string]map[string]interface{}{}

	m.data["0_create-funcs"] = map[string]map[string]interface{}{
		"0_meta": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_sets": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_attrs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_rels": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_get-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_create-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_update-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_delete-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		}}

	m.data["0_update-funcs"] = map[string]map[string]interface{}{
		"0_meta": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_sets": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_attrs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_rels": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_get-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_create-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_update-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_delete-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		}}

	m.data["0_delete-funcs"] = map[string]map[string]interface{}{
		"0_meta": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_sets": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_attrs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_rels": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_get-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_create-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_update-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		},
		"0_delete-funcs": map[string]interface{}{
			"func": `func(snap *Snapshot) error {
	snap.Fail(ErrNotImplemented)
}`,
		}}

	// Fields
	m.fields = map[string][]string{}
	m.fields["0_meta"] = []string{"id"}
	m.fields["0_sets"] = []string{"id", "version", "fields"}
	m.fields["0_attrs"] = []string{"id", "name", "type", "null", "set"}
	m.fields["0_get-funcs"] = []string{"id", "func"}
	m.fields["0_create-funcs"] = []string{"id", "func"}
	m.fields["0_update-funcs"] = []string{"id", "func"}
	m.fields["0_delete-funcs"] = []string{"id", "func"}
	// Make sure the fields are sorted
	for k := range m.fields {
		sort.Strings(m.fields[k])
	}

	return nil
}

// Resource ...
func (m *Memory) Resource(qry karigo.QueryRes) (map[string]interface{}, error) {
	m.Lock()

	var ok bool

	// Get set
	var set map[string]map[string]interface{}
	if set, ok = m.data[qry.Set]; !ok {
		return nil, karigo.ErrUnexpected
	}

	// Get resource
	var res map[string]interface{}
	if res, ok = set[qry.ID]; !ok {
		return nil, karigo.ErrNotFound
	}

	// Filter fields
	for field := range res {
		found := false
		for _, f := range qry.Fields {
			if field == f {
				found = true
			}
		}
		if !found {
			delete(res, field)
		}
	}

	m.Unlock()
	return res, nil
}

// Collection ...
func (m *Memory) Collection(qry karigo.QueryCol) ([]map[string]interface{}, error) {
	m.Lock()
	defer m.Unlock()

	// Get all records from the given set
	var recs []map[string]interface{}
	if _, ok := m.data[qry.Set]; ok {
		recs = make([]map[string]interface{}, 0, len(m.data[qry.Set]))
		for id := range m.data[qry.Set] {
			recs = append(recs, m.data[qry.Set][id])
		}
	}

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
		kept := make([]map[string]interface{}, 0, len(recs))
		if ids, ok := res[qry.BelongsToFilter.Name].([]string); ok {
			for i := range recs {
				if id, ok := recs[i]["id"].(string); ok {
					keep := false
					for i := range ids {
						if id == ids[i] {
							keep = true
							break
						}
					}
					if keep {
						kept = append(kept)
					}
				} else {
					return nil, errors.New("karigo: field id is not a string")
				}
			}
			recs = kept
		}
	}

	// TODO Filter

	// Sort
	ss := &sortableSet{
		recs:  recs,
		sorts: qry.Sort,
	}
	for id := range m.data[qry.Set] {
		ss.recs = append(ss.recs, m.data[qry.Set][id])
	}
	sort.Sort(ss)

	// Pagination
	if qry.PageSize == 0 {
		recs = []map[string]interface{}{}
	} else {
		skip := qry.PageNumber * qry.PageSize
		if skip >= len(recs) {
			recs = []map[string]interface{}{}
		} else {
			page := make([]map[string]interface{}, 0, qry.PageSize)
			for i := skip; i < len(recs) || i < qry.PageSize; i++ {
				page = append(page, recs[i])
			}
			recs = page
		}
	}

	// Fields
	for i := range ss.recs {
		for k := range ss.recs[i] {
			found := false
			for _, f := range qry.Fields {
				if k == f {
					found = true
				}
			}
			if !found {
				delete(ss.recs[i], k)
			}
		}
	}

	return nil, nil
}

// Apply ...
func (m *Memory) Apply(ops []karigo.Op) error {
	m.Lock()
	defer m.Unlock()

	// TODO Create a mstx, apply the ops, commit it

	return nil
}

// Begin ...
func (m *Memory) Begin() (karigo.SourceTx, error) {
	m.Lock()
	defer m.Unlock()

	oldData, err := copystructure.Copy(m.data)
	if err != nil {
		return nil, err
	}
	m.oldData = oldData.(map[string]map[string]map[string]interface{})

	oldFields, err := copystructure.Copy(m.fields)
	if err != nil {
		return nil, err
	}
	m.oldFields = oldFields.(map[string][]string)

	return &mstx{
		ms: m,
	}, nil
}

type mstx struct {
	ms   *Memory
	undo []karigo.Op

	sync.Mutex
}

// Apply ...
func (m *mstx) Apply(ops []karigo.Op) error {
	m.Lock()
	defer m.Unlock()

	for _, op := range ops {
		switch op.Op {
		case karigo.OpSet:
			m.opSet(op.Key.Set, op.Key.ID, op.Key.Field, op.Value)
		default:
			return karigo.ErrUnexpected
		}
	}

	return nil
}

// Rollback ...
func (m *mstx) Rollback() error {
	m.Lock()
	defer m.Unlock()

	m.ms.data = m.ms.oldData
	m.ms.fields = m.ms.oldFields
	m.ms.oldData = map[string]map[string]map[string]interface{}{}
	m.ms.oldFields = map[string][]string{}

	return nil
}

// Commit ...
func (m *mstx) Commit() error {
	m.Lock()
	defer m.Unlock()

	m.ms.oldData = map[string]map[string]map[string]interface{}{}
	m.ms.oldFields = map[string][]string{}

	return nil
}

func (m *mstx) opSet(set, id, field string, v interface{}) {
	// fmt.Printf("set, id, field = %s, %s, %s (%v)\n", set, id, field, v)

	if id != "" && field != "id" {
		m.ms.data[set][id][field] = v
	}

	if id == "" && field == "id" {
		m.ms.data[set][v.(string)] = map[string]interface{}{}
		// fmt.Printf("New entry inserted.\n")
	} else if strings.HasPrefix(set, "0_") && field == "created" {
		// If a set, attribute, or relationship is marked as created, create it.
		switch field {
		case "created":
			switch set {
			case "0_sets":
				name := m.ms.data["0_sets"][id]["name"].(string)
				m.ms.data[name] = map[string]map[string]interface{}{}
			case "0_attrs":
				name := m.ms.data["0_attrs"][id]["name"].(string)
				typ := m.ms.data["0_attrs"][id]["type"].(string)
				set := m.ms.data["0_attrs"][id]["set"].(string)
				for id2 := range m.ms.data[set] {
					fmt.Printf("Created: %s %s %s\n", set, id2, name)
					m.ms.data[set][id2][name] = zeroVal(typ)
				}
			case "0_rels":
				name := m.ms.data["0_rels"][id]["name"].(string)
				toOne := m.ms.data["0_rels"][id]["to-one"].(bool)
				set := m.ms.data["0_rels"][id]["set"].(string)
				for id2 := range m.ms.data[set] {
					if toOne {
						m.ms.data[set][id2][name] = ""
					} else {
						m.ms.data[set][id2][name] = []string{}
					}
				}
			}
		}
		// fmt.Printf("created=true, new thing created.\n")
	} else {
		// if _, ok := m.ms.data[set]; !ok {
		// 	fmt.Printf("Set %s does not exist.\n", set)
		// }
		// if _, ok := m.ms.data[set][id]; !ok {
		// 	fmt.Printf("ID %s does not exist.\n", id)
		// }
		// if _, ok := m.ms.data[set][id][field]; !ok {
		// 	fmt.Printf("Field %s does not exist.\n", field)
		// }
		m.ms.data[set][id][field] = v
	}
}

// sortableSet ...
type sortableSet struct {
	recs  []map[string]interface{}
	sorts []string
}

// Len ...
func (s *sortableSet) Len() int { return len(s.recs) }

// Swap ...
func (s *sortableSet) Swap(i, j int) { s.recs[i], s.recs[j] = s.recs[j], s.recs[i] }

// Less ...
func (s *sortableSet) Less(i, j int) bool {
	less := false

	for _, r := range s.sorts {
		inverse := false
		if strings.HasPrefix(r, "-") {
			r = r[1:]
			inverse = true
		}

		switch v := s.recs[i][r].(type) {
		case string:
			if v == s.recs[j][r].(string) {
				continue
			}
			if inverse {
				return v > s.recs[j][r].(string)
			}
			return v < s.recs[j][r].(string)
		case int:
			if v == s.recs[j][r].(int) {
				continue
			}
			if inverse {
				return v > s.recs[j][r].(int)
			}
			return v < s.recs[j][r].(int)
		case bool:
			if v == s.recs[j][r].(bool) {
				continue
			}
			if inverse {
				return v
			}
			return !v
		case time.Time:
			if v.Equal(s.recs[j][r].(time.Time)) {
				continue
			}
			if inverse {
				return v.After(s.recs[j][r].(time.Time))
			}
			return v.Before(s.recs[j][r].(time.Time))
		}
	}

	return less
}

func zeroVal(typ string) interface{} {
	switch typ {
	case "string":
		return string("")
	case "int":
		return int(0)
	case "int8":
		return int8(0)
	case "int16":
		return int16(0)
	case "int32":
		return int32(0)
	case "int64":
		return int64(0)
	case "uint":
		return uint(0)
	case "uint8":
		return uint8(0)
	case "uint16":
		return uint16(0)
	case "uint32":
		return uint32(0)
	case "bool":
		return bool(false)
	case "time":
		return time.Time{}
	case "*string":
		v := string("")
		return &v
	case "*int":
		v := int(0)
		return &v
	case "*int8":
		v := int8(0)
		return &v
	case "*int16":
		v := int16(0)
		return &v
	case "*int32":
		v := int32(0)
		return &v
	case "*int64":
		v := int64(0)
		return &v
	case "*uint":
		v := uint(0)
		return &v
	case "*uint8":
		v := uint8(0)
		return &v
	case "*uint16":
		v := uint16(0)
		return &v
	case "*uint32":
		v := uint32(0)
		return &v
	case "*bool":
		v := bool(false)
		return &v
	case "*time":
		return &time.Time{}
	default:
		return nil
	}
}
