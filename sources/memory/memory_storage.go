package memory

import (
	"github.com/mfcochauxlaberge/jsonapi"
)

type db struct {
	schema *jsonapi.Schema
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

func (d *db) withRecords(recs ...record) *db {
	for _, rec := range recs {
		rec.schema = d.schema
		d.insert(rec)
	}
	return d
}

func (d *db) insert(rec record) {
	if rec.id != "" {
		d.recs[rec.id] = rec
	}
}

type record struct {
	schema *jsonapi.Schema
	id     string
	vals   map[string]interface{}
}

// func newRecord() *db {
// 	return &db{
// 		recs: map[string]record{},
// 	}
// }

// type val struct {
// 	typ int
// 	val []byte
// }

// func newVal() *db {
// 	return &db{
// 		recs: map[string]record{},
// 	}
// }
