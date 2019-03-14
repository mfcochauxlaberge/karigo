package memory

import (
	"github.com/mfcochauxlaberge/jsonapi"
)

type db struct {
	schema jsonapi.Schema
	recs   map[string]record
}

func newDB() *db {
	return &db{
		recs: map[string]record{},
	}
}

func (d *db) withAttrs(attrs ...jsonapi.Attr) *db {
	for _, attr := range attrs {
		d.schema.AddAttr("", attr) // TODO Type
	}
	return d
}

func (d *db) withRels(rels ...jsonapi.Rel) *db {
	for _, rel := range rels {
		d.schema.AddRel("", rel) // TODO Type
	}
	return d
}

// func (d *db) withRels(rec record) {

// }

func (d *db) insert(rec record) {

}

type record struct {
	schema *jsonapi.Schema
	vals   map[string]val
}

func newRecord() *db {
	return &db{
		recs: map[string]record{},
	}
}

type val struct {
	typ int
	val []byte
}

func newVal() *db {
	return &db{
		recs: map[string]record{},
	}
}
