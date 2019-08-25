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
func (s *Checkpoint) Resource(qry QueryRes) jsonapi.Resource {
	if s.err != nil {
		return nil
	}

	res, err := s.node.resource(s.version, qry)
	if err != nil {
		s.Check(err)
		return nil
	}

	return res
}

// Collection ...
func (s *Checkpoint) Collection(qry QueryCol) jsonapi.Collection {
	if s.err != nil {
		return nil
	}

	col, err := s.node.collection(s.version, qry)
	if err != nil {
		s.Check(err)
		return nil
	}

	return col
}

// Apply ...
func (s *Checkpoint) Apply(ops []Op) {}

// Check ...
func (s *Checkpoint) Check(err error) {
	if err != nil && s.err == nil {
		s.err = err
	}
}

// Fail ...
func (s *Checkpoint) Fail(err error) {
	if err == nil {
		err = errors.New("an error occured")
	}
	s.err = err
}
