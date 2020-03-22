package karigo

import (
	"errors"

	"github.com/mfcochauxlaberge/karigo/query"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Checkpoint ...
type Checkpoint struct {
	ResID string
	Res   jsonapi.Resource

	tx query.Tx

	ops []query.Op

	err error
}

// Resource ...
func (c *Checkpoint) Resource(qry query.Res) jsonapi.Resource {
	if c.err != nil {
		return nil
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

	col, err := c.tx.Collection(qry)
	if err != nil {
		c.Check(err)
		return nil
	}

	return col
}

// OpsRes ...
func (c *Checkpoint) OpsRes(typ, id string) jsonapi.Resource {
	opsc := make([]query.Op)
	return resourceOverOps{typ: typ, id: id, ops: c.ops}
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
	typ string
	id  string
	ops []query.Op
}

func (r resourceOverOps) New() jsonapi.Resource {
	return nil
}

func (r resourceOverOps) Copy() jsonapi.Resource {
	return nil
}

func (r resourceOverOps) Attrs() map[string]jsonapi.Attr {
	return nil
}

func (r resourceOverOps) Rels() map[string]jsonapi.Rel {
	return nil
}

func (r resourceOverOps) Attr(key string) jsonapi.Attr {
	return jsonapi.Attr{}
}

func (r resourceOverOps) Rel(key string) jsonapi.Rel {
	return jsonapi.Rel{}
}

func (r resourceOverOps) GetID() string {
	return ""
}

func (r resourceOverOps) GetType() jsonapi.Type {
	return jsonapi.Type{}
}

func (r resourceOverOps) Get(key string) interface{} {
	return nil
}

func (r resourceOverOps) GetToOne(key string) string {
	return ""
}

func (r resourceOverOps) GetToMany(key string) []string {
	return nil
}

func (r resourceOverOps) SetID(id string) {}

func (r resourceOverOps) Set(key string, val interface{}) {}

func (r resourceOverOps) SetToOne(key string, rel string) {}

func (r resourceOverOps) SetToMany(key string, rels []string) {}
