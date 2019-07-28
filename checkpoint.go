package karigo

import (
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
		s.Fail(err)
		return nil
	}

	return res
}

// Collection ...
func (s *Checkpoint) Collection(qry QueryCol) []jsonapi.Resource {
	if s.err != nil {
		return nil
	}

	col, err := s.node.collection(s.version, qry)
	if err != nil {
		s.Fail(err)
		return nil
	}

	return col
}

// Apply ...
func (s *Checkpoint) Apply(ops []Op) error {
	return nil
}

// Fail ...
func (s *Checkpoint) Fail(err error) {
	if s.err == nil {
		s.err = err
	}
}
