package karigo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Operations
const (
	// OpSet means that the key must be set
	// to the value.
	OpSet = '='

	// OpAdd means that the key must be
	// incremented by the value. If the new
	// value is above its maximum, the key
	// must be set to that maximum.
	OpAdd = '+'

	// OpSubtract means that the key must
	// be decremented by the value. If the
	// new value is under its minimum, the
	// key must be set to that minimum.
	OpSubtract = '-'

	// OpInsert means that the value must
	// be added to the set represented by
	// the key if it is not present.
	OpInsert = '>'

	// OpRemove means that the value must
	// be removed from the set represented
	// by the key if it is present.
	OpRemove = '<'
)

// ParseOp parses s and returns an Op object.
//
// The concrete type of Op.Value will
func ParseOp(s string, t int) (Op, error) {
	parts := strings.Split(s, " ")

	if len(parts) != 3 {
		return Op{}, errors.New("invalid op")
	}

	key := strings.Split(parts[0], ".")

	if len(key) != 3 {
		return Op{}, errors.New("invalid key")
	}

	return Op{
		Key: Key{
			Set:   key[0],
			ID:    key[1],
			Field: key[2],
		},
		Op:    parts[1][0],
		Value: parts[2],
	}, nil
}

// Op ...
type Op struct {
	Key   Key // Set, ID, Field
	Op    byte
	Value interface{}
}

// String ...
func (o Op) String() string {
	id := o.Key.ID
	if id == "" {
		id = "_"
	}

	return fmt.Sprintf(
		"%s.%s.%s %v %v",
		o.Key.Set,
		id,
		o.Key.Field,
		string(o.Op),
		o.Value,
	)
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

// NewOpSubtract ...
func NewOpSubtract(set, id, field string, v interface{}) Op {
	return Op{
		Key: Key{
			Set:   set,
			ID:    id,
			Field: field,
		},
		Op:    OpSubtract,
		Value: v,
	}
}

// NewOpInsert ...
func NewOpInsert(set, id, field string, v interface{}) Op {
	return Op{
		Key: Key{
			Set:   set,
			ID:    id,
			Field: field,
		},
		Op:    OpInsert,
		Value: v,
	}
}

// NewOpRemove ...
func NewOpRemove(set, id, field string, v interface{}) Op {
	return Op{
		Key: Key{
			Set:   set,
			ID:    id,
			Field: field,
		},
		Op:    OpRemove,
		Value: v,
	}
}

// NewOpCreateRes ...
func NewOpCreateRes(res jsonapi.Resource) []Op {
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

// NewOpCreateSet ...
func NewOpCreateSet(set string) []Op {
	return []Op{
		NewOpSet("0_sets", "", "id", set),
		NewOpSet("0_sets", set, "name", set),
		NewOpSet("0_sets", set, "version", 0),
		NewOpSet("0_sets", set, "created", true),
		NewOpSet("0_sets", set, "active", false),
	}
}

// NewOpDeleteSet ...
func NewOpDeleteSet(id string) []Op {
	return []Op{
		NewOpSet("0_sets", id, "id", ""),
	}
}

// NewOpActivateSet ...
func NewOpActivateSet(set string) []Op {
	return []Op{
		NewOpSet("0_sets", set, "active", true),
	}
}

// NewOpDeactivateSet ...
func NewOpDeactivateSet(set string) []Op {
	return []Op{
		NewOpSet("0_sets", set, "active", false),
	}
}

// NewOpCreateAttr ...
func NewOpCreateAttr(set, name, typ string, null bool) []Op {
	id := set + "_" + name

	return []Op{
		NewOpSet("0_attrs", "", "id", id),
		NewOpSet("0_attrs", id, "name", name),
		NewOpSet("0_attrs", id, "type", typ),
		NewOpSet("0_attrs", id, "null", null),
		NewOpSet("0_attrs", id, "set", set),
		NewOpSet("0_attrs", id, "created", true),
		NewOpSet("0_attrs", id, "active", false),
		NewOpInsert("0_sets", set, "attrs", id),
	}
}

// NewOpDeleteAttr ...
func NewOpDeleteAttr(set, name string) []Op {
	id := set + "_" + name

	return []Op{
		NewOpSet("0_attrs", id, "id", ""),
	}
}

// NewOpActivateAttr ...
func NewOpActivateAttr(set, name string) []Op {
	id := set + "_" + name

	return []Op{
		NewOpSet("0_attrs", id, "active", true),
	}
}

// NewOpDeactivateAttr ...
func NewOpDeactivateAttr(set, name string) []Op {
	id := set + "_" + name

	return []Op{
		NewOpSet("0_attrs", id, "active", false),
	}
}

// NewOpCreateRel ...
func NewOpCreateRel(fromSet, fromName, toSet, toName string, toOne, fromOne bool) []Op {
	id := fromSet + "_" + fromName

	if toName != "" {
		id2 := toSet + "_" + toName
		if id < id2 {
			id = id + "_" + id2
		} else {
			id = id2 + "_" + id
		}
	}

	return []Op{
		NewOpSet("0_rels", "", "id", id),
		NewOpSet("0_rels", id, "from-name", fromName),
		NewOpSet("0_rels", id, "from-set", fromSet),
		NewOpSet("0_rels", id, "to-one", toOne),
		NewOpSet("0_rels", id, "to-set", toSet),
		NewOpSet("0_rels", id, "to-name", toName),
		NewOpSet("0_rels", id, "from-one", fromOne),
		NewOpSet("0_rels", id, "created", true),
		NewOpSet("0_rels", id, "active", false),
		NewOpInsert("0_sets", fromSet, "rels", id),
	}
}

// NewOpDeleteRel ...
func NewOpDeleteRel(id string) []Op {
	return []Op{
		NewOpSet("0_attrs", id, "id", ""),
	}
}

// NewOpActivateRel ...
func NewOpActivateRel(id string) []Op {
	return []Op{
		NewOpSet("0_rels", id, "active", true),
	}
}

// NewOpDeactivateRel ...
func NewOpDeactivateRel(id string) []Op {
	return []Op{
		NewOpSet("0_rels", id, "active", false),
	}
}
