package karigo

// Operations
const (
	OpSet = iota
	OpAdd = iota
)

// Op ...
type Op struct {
	Key   Key // Set, ID, Field
	Op    int
	Value interface{}
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
