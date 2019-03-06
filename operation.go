package karigo

// Operations
const (
	OpSet = iota
	// OpRem
	// OpAdd
	// OpReset
	// OpInc
	// OpDec
)

// Op ...
type Op struct {
	Key   Key // Set, ID, Field
	Op    int
	Value interface{}
}

// OpAddSet ...
func OpAddSet(set string) []Op {
	return []Op{
		Op{
			Key: Key{
				Set:   "0_sets",
				ID:    "",
				Field: "id",
			},
			Op:    OpSet,
			Value: set,
		},
		Op{
			Key: Key{
				Set:   "0_sets",
				ID:    set,
				Field: "name",
			},
			Op:    OpSet,
			Value: set,
		},
		Op{
			Key: Key{
				Set:   "0_sets",
				ID:    set,
				Field: "version",
			},
			Op:    OpSet,
			Value: 0,
		},
		Op{
			Key: Key{
				Set:   "0_sets",
				ID:    set,
				Field: "created",
			},
			Op:    OpSet,
			Value: true,
		},
	}
}

// OpAddAttr ...
func OpAddAttr(set, name, typ string, null bool) []Op {
	id := set + "_" + name
	return []Op{
		Op{
			Key: Key{
				Set:   "0_attrs",
				ID:    "",
				Field: "id",
			},
			Op:    OpSet,
			Value: id,
		},
		Op{
			Key: Key{
				Set:   "0_attrs",
				ID:    id,
				Field: "name",
			},
			Op:    OpSet,
			Value: name,
		},
		Op{
			Key: Key{
				Set:   "0_attrs",
				ID:    id,
				Field: "type",
			},
			Op:    OpSet,
			Value: typ,
		},
		Op{
			Key: Key{
				Set:   "0_attrs",
				ID:    id,
				Field: "null",
			},
			Op:    OpSet,
			Value: null,
		},
		Op{
			Key: Key{
				Set:   "0_attrs",
				ID:    id,
				Field: "set",
			},
			Op:    OpSet,
			Value: set,
		},
		Op{
			Key: Key{
				Set:   "0_attrs",
				ID:    id,
				Field: "created",
			},
			Op:    OpSet,
			Value: true,
		},
	}
}

// OpAddRel ...
func OpAddRel(set, name string, toOne bool) []Op {
	id := set + "_" + name
	return []Op{
		Op{
			Key: Key{
				Set:   "0_rels",
				ID:    "",
				Field: "id",
			},
			Op:    OpSet,
			Value: id,
		},
		Op{
			Key: Key{
				Set:   "0_rels",
				ID:    id,
				Field: "name",
			},
			Op:    OpSet,
			Value: name,
		},
		Op{
			Key: Key{
				Set:   "0_rels",
				ID:    id,
				Field: "to-one",
			},
			Op:    OpSet,
			Value: toOne,
		},
		Op{
			Key: Key{
				Set:   "0_rels",
				ID:    id,
				Field: "set",
			},
			Op:    OpSet,
			Value: set,
		},
		Op{
			Key: Key{
				Set:   "0_rels",
				ID:    id,
				Field: "created",
			},
			Op:    OpSet,
			Value: true,
		},
	}
}
