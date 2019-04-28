package karigo

// Operations
const (
	OpSet = iota
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

// NewOpAddSet ...
func NewOpAddSet(set string) []Op {
	return []Op{
		NewOpSet("0_sets", "", "id", set),
		NewOpSet("0_sets", set, "name", set),
		NewOpSet("0_sets", set, "version", 0),
		NewOpSet("0_sets", set, "created", true),
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
		NewOpSet("0_attrs", id, "created", true),
	}
}

// NewOpAddRel ...
func NewOpAddRel(set, name string, toOne bool) []Op {
	id := set + "_" + name
	return []Op{
		NewOpSet("0_rels", "", "id", id),
		NewOpSet("0_rels", id, "name", name),
		NewOpSet("0_rels", id, "to-one", toOne),
		NewOpSet("0_rels", id, "set", set),
		NewOpSet("0_rels", id, "created", true),
	}
}
