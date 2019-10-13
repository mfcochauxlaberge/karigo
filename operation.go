package karigo

import (
	"encoding/json"
	"fmt"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Operations
const (
	OpSet = '='
	OpAdd = '&'
)

// Op ...
type Op struct {
	Key   Key // Set, ID, Field
	Op    byte
	Value interface{}
}

type Entry []Op

// Bytes ...
func (e *Entry) Bytes() []byte {
	b, err := json.Marshal(e)
	if err != nil {
		panic(fmt.Errorf("can't get Entry bytes: %s", err))
	}
	return b
}

// NewOpSet ...
func NewOpSet(set, id, field string, v interface{}) Op {
	return Op{
		Key: Key{
			Set:   set,
			ID:    id,
			Field: field,
		},
		Op:    OpSet,
		Value: v,
	}
}

// NewOpAdd ...
func NewOpAdd(set, id, field string, v interface{}) Op {
	return Op{
		Key: Key{
			Set:   set,
			ID:    id,
			Field: field,
		},
		Op:    OpAdd,
		Value: v,
	}
}

// NewOpInsert ...
func NewOpInsert(res jsonapi.Resource) []Op {
	set := res.GetType().Name
	id := res.GetID()
	ops := []Op{}

	// New resource
	ops = append(ops, NewOpSet(set, "", "id", id))

	// Attributes
	for _, attr := range res.Attrs() {
		ops = append(ops, NewOpSet(set, id, attr.Name, res.Get(attr.Name)))
	}

	// Relationships
	for _, rel := range res.Rels() {
		var op Op
		if rel.ToOne {
			op = NewOpSet(set, id, rel.FromName, res.GetToOne(rel.FromName))
		} else {
			op = NewOpSet(set, id, rel.FromName, res.GetToMany(rel.FromName))
		}
		ops = append(ops, op)
	}

	return ops
}

// NewOpAddSet ...
func NewOpAddSet(set string) []Op {
	return []Op{
		NewOpSet("0_sets", "", "id", set),
		NewOpSet("0_sets", set, "name", set),
		NewOpSet("0_sets", set, "version", 0),
		NewOpSet("0_sets", set, "active", true),
	}
}

// NewOpAddAttr ...
func NewOpAddAttr(set, name, typ string, null bool) []Op {
	id := set + "_" + name
	return []Op{
		NewOpSet("0_attrs", "", "id", id),
		NewOpSet("0_attrs", id, "name", name),
		NewOpSet("0_attrs", id, "type", typ),
		NewOpSet("0_attrs", id, "null", null),
		NewOpSet("0_attrs", id, "set", set),
		NewOpSet("0_attrs", id, "active", true),
		NewOpAdd("0_sets", set, "attrs", id),
	}
}

// NewOpAddRel ...
func NewOpAddRel(fromSet, fromName, toSet, toName string, toOne, fromOne bool) []Op {
	id := fromSet + "_" + fromName
	return []Op{
		NewOpSet("0_rels", "", "id", id),
		NewOpSet("0_rels", id, "from-name", fromName),
		NewOpSet("0_rels", id, "from-set", fromSet),
		NewOpSet("0_rels", id, "to-one", toOne),
		NewOpSet("0_rels", id, "to-set", toSet),
		NewOpSet("0_rels", id, "to-name", toName),
		NewOpSet("0_rels", id, "from-one", fromOne),
		NewOpSet("0_rels", id, "active", true),
		NewOpAdd("0_sets", fromSet, "rels", id),
	}
}
