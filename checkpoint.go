package karigo

import (
	"errors"

	"github.com/mfcochauxlaberge/karigo/query"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Checkpoint ...
type Checkpoint struct {
	Res jsonapi.Resource

	tx     query.Tx
	schema *jsonapi.Schema

	ops []query.Op

	err error
}

// Resource ...
func (c *Checkpoint) Resource(qry query.Res) jsonapi.Resource {
	if c.err != nil {
		return nil
	}

	if c.Res == nil {
		c.Res = resourceOverOps{
			typ:    "TODO",
			id:     "$0",
			ops:    c.ops,
			schema: c.schema,
		}
	}

	res, err := c.tx.Resource(qry)
	if err != nil {
		c.Check(err)
		return nil
	}

	return res
}

// Collection ...
func (c *Checkpoint) Collection(qry query.Col) jsonapi.Collection {
	if c.err != nil {
		return nil
	}

	if c.Res == nil {
		c.Res = resourceOverOps{
			typ:    "TODO",
			id:     "$0",
			ops:    c.ops,
			schema: c.schema,
		}
	}

	col, err := c.tx.Collection(qry)
	if err != nil {
		c.Check(err)
		return nil
	}

	return col
}

// Apply ...
func (c *Checkpoint) Apply(ops []query.Op) {
	if c.err == nil {
		c.Check(c.tx.Apply(ops))
	}

	if c.err == nil {
		c.ops = append(c.ops, ops...)
	}
}

// Check ...
func (c *Checkpoint) Check(err error) {
	if err != nil && c.err == nil {
		c.err = err
	}
}

// Fail ...
func (c *Checkpoint) Fail(err error) {
	if err == nil {
		err = errors.New("an error occurred")
	}

	c.err = err
}

// commit ...
func (c *Checkpoint) commit() error {
	return c.tx.Commit()
}

// rollback ...
func (c *Checkpoint) rollback() error {
	return c.tx.Rollback()
}

// resourceOverOps ...
type resourceOverOps struct {
	typ    string
	id     string
	ops    []query.Op
	schema *jsonapi.Schema
}

func (r resourceOverOps) New() jsonapi.Resource {
	typ := r.GetType()

	sr := &jsonapi.SoftResource{}
	sr.SetType(&typ)

	return sr
}

func (r resourceOverOps) Copy() jsonapi.Resource {
	typ := r.GetType()

	sr := &jsonapi.SoftResource{}
	sr.SetType(&typ)

	// Copy the ID
	sr.SetID(r.GetID())

	// Copy all fields
	for key := range r.Attrs() {
		sr.Set(key, r.Get(key))
	}

	for key, rel := range r.Rels() {
		if rel.ToOne {
			sr.SetToOne(key, r.GetToOne(key))
		} else {
			sr.SetToMany(key, r.GetToMany(key))
		}
	}

	return sr
}

func (r resourceOverOps) Attrs() map[string]jsonapi.Attr {
	attrs := map[string]jsonapi.Attr{}

	for _, op := range r.ops {
		if op.Key.Set == r.typ && op.Key.ID == r.id {
			if attr, ok := r.schema.GetType(r.typ).Attrs[op.Key.Field]; ok {
				attrs[op.Key.Field] = attr
			}
		}
	}

	return attrs
}

func (r resourceOverOps) Rels() map[string]jsonapi.Rel {
	rels := map[string]jsonapi.Rel{}

	for _, op := range r.ops {
		if op.Key.Set == r.typ && op.Key.ID == r.id {
			if rel, ok := r.schema.GetType(r.typ).Rels[op.Key.Field]; ok {
				rels[op.Key.Field] = rel
			}
		}
	}

	return rels
}

func (r resourceOverOps) Attr(key string) jsonapi.Attr {
	for _, op := range r.ops {
		if op.Key.Set == r.typ && op.Key.ID == r.id {
			if attr, ok := r.schema.GetType(r.typ).Attrs[op.Key.Field]; ok {
				return attr
			}
		}
	}

	return jsonapi.Attr{}
}

func (r resourceOverOps) Rel(key string) jsonapi.Rel {
	for _, op := range r.ops {
		if op.Key.Set == r.typ && op.Key.ID == r.id {
			if rel, ok := r.schema.GetType(r.typ).Rels[op.Key.Field]; ok {
				return rel
			}
		}
	}

	return jsonapi.Rel{}
}

func (r resourceOverOps) GetID() string {
	return r.id
}

func (r resourceOverOps) GetType() jsonapi.Type {
	typ := jsonapi.Type{
		Name:  r.typ,
		Attrs: r.Attrs(),
		Rels:  r.Rels(),
	}

	return typ
}

func (r resourceOverOps) Get(key string) interface{} {
	for k := range r.Attrs() {
		if k == key {
			for _, op := range r.ops {
				if op.Key.Set == r.typ && op.Key.ID == r.id && op.Key.Field == k {
					return op.Value
				}
			}
		}
	}

	return nil
}

func (r resourceOverOps) GetToOne(key string) string {
	for k := range r.Rels() {
		if k == key {
			for _, op := range r.ops {
				if op.Key.Set == r.typ && op.Key.ID == r.id && op.Key.Field == k {
					return op.Value.(string)
				}
			}
		}
	}

	return ""
}

func (r resourceOverOps) GetToMany(key string) []string {
	for k := range r.Rels() {
		if k == key {
			for _, op := range r.ops {
				if op.Key.Set == r.typ && op.Key.ID == r.id && op.Key.Field == k {
					return op.Value.([]string)
				}
			}
		}
	}

	return nil
}

func (r resourceOverOps) SetID(id string) {
	for i := range r.ops {
		if r.ops[i].Key.Set == r.typ && r.ops[i].Key.ID == r.id {
			r.ops[i].Key.ID = id
		}
	}

	r.id = id
}

func (r resourceOverOps) Set(key string, val interface{}) {
	for k := range r.Attrs() {
		if k == key {
			for i := range r.ops {
				if r.ops[i].Key.Set == r.typ && r.ops[i].Key.ID == r.id && r.ops[i].Key.Field == k {
					r.ops[i].Value = val
				}
			}
		}
	}
}

func (r resourceOverOps) SetToOne(key string, rel string) {
	for k := range r.Rels() {
		if k == key {
			for i := range r.ops {
				if r.ops[i].Key.Set == r.typ && r.ops[i].Key.ID == r.id && r.ops[i].Key.Field == k {
					r.ops[i].Value = rel
				}
			}
		}
	}
}

func (r resourceOverOps) SetToMany(key string, rels []string) {
	for k := range r.Attrs() {
		if k == key {
			for i := range r.ops {
				if r.ops[i].Key.Set == r.typ && r.ops[i].Key.ID == r.id && r.ops[i].Key.Field == k {
					r.ops[i].Value = rels
				}
			}
		}
	}
}
