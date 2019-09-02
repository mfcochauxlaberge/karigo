package karigo

import (
	"errors"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Checkpoint ...
type Checkpoint struct {
	Res jsonapi.Resource
	Inc map[string]jsonapi.Resource

	node *Node

	version uint
	// locks   map[string]bool // false for read, true for write
	// tx    SourceTx
	ops []Op
	// undo  []Op
	// ready bool

	err error
}

// Resource ...
func (c *Checkpoint) Resource(qry QueryRes) jsonapi.Resource {
	if c.err != nil {
		return nil
	}

	res, err := c.node.resource(c.version, qry)
	if err != nil {
		c.Check(err)
		return nil
	}

	return res
}

// Collection ...
func (c *Checkpoint) Collection(qry QueryCol) jsonapi.Collection {
	if c.err != nil {
		return nil
	}

	col, err := c.node.collection(c.version, qry)
	if err != nil {
		c.Check(err)
		return nil
	}

	return col
}

// Apply ...
func (c *Checkpoint) Apply(ops []Op) {
	c.Check(c.node.apply(ops))
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
