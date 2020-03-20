package karigo

import (
	"errors"

	"github.com/mfcochauxlaberge/karigo/query"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Checkpoint ...
type Checkpoint struct {
	Res jsonapi.Resource
	Inc map[string]jsonapi.Resource

	tx   query.Tx
	node *Node

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
