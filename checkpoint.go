package karigo

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mfcochauxlaberge/karigo/query"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Checkpoint ...
type Checkpoint struct {
	tx     query.Tx
	schema *jsonapi.Schema
	node   *Node

	ops []query.Op
	ids map[string]string

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

// Val finds and returns the value the given field is set to in the current ops.
//
// The second return argument reports whether a set operation was found. Other
// operations (add, subtract, insert, etc) are ignored since a fix value cannot
// be calculated.
func (c *Checkpoint) Val(set, id, field string) (interface{}, bool) {
	for _, op := range c.ops {
		if op.Key.Set == set && op.Key.ID == id && op.Key.Field == field {
			return op.Value, true
		}
	}

	return nil, false
}

// NumChange finds and returns the change applied to the given field in the
// current ops.
//
// A change is represented by OpAdd or OpSubtract. It is assumed that the field
// is of numerical type (int, int8, uint, etc).
func (c *Checkpoint) NumChange(set, id, field string) int {
	for _, op := range c.ops {
		if op.Key.Set == set && op.Key.ID == id && op.Key.Field == field {
			i, _ := strconv.ParseInt(fmt.Sprintf("%d", op.Value), 10, 64)
			return int(i)
		}
	}

	return 0
}

// RelsChange finds and returns the inserted and removed IDs for the given
// relationship.
//
// The first return argument is the added IDs, and the second one is the removed
// IDs.
//
// A change is represented by OpInset or OpRemove. It is assumed that the field
// is of type []byte or []string for to-many relationships.
func (c *Checkpoint) RelsChange(set, id, field string) ([]string, []string) {
	var insertions, removals []string

	for _, op := range c.ops {
		if op.Key.Set == set && op.Key.ID == id && op.Key.Field == field {
			if op.Op == query.OpInsert {
				insertions = op.Value.([]string)
			} else {
				removals = op.Value.([]string)
			}
		}
	}

	return insertions, removals
}

// SetID ...
func (c *Checkpoint) SetID(placeholder, id string) interface{} {
	if c.ids == nil {
		c.ids = map[string]string{}
	}

	c.ids[placeholder] = id

	return nil
}

// commit ...
func (c *Checkpoint) commit() error {
	return c.tx.Commit()
}

// rollback ...
func (c *Checkpoint) rollback() error {
	return c.tx.Rollback()
}
