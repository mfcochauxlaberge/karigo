package karigo

import (
	"errors"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Snapshot ...
type Snapshot struct {
	Res jsonapi.Resource
	Inc map[string]jsonapi.Resource

	node *Node

	version uint
	locks   map[string]bool // false for read, true for write
	tx      SourceTx
	ops     []Op
	undo    []Op
	ready   bool

	err error
}

// Resource ...
func (s *Snapshot) Resource(qry QueryRes) jsonapi.Resource {
	if s.err != nil {
		return nil
	}

	allowed := false
	for n := range s.locks {
		if n == qry.Set {
			allowed = true
			break
		}
	}
	if !allowed {
		s.Fail(errors.New("karigo: can't get resource fron unlocked set"))
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
func (s *Snapshot) Collection(qry QueryCol) []jsonapi.Resource {
	if s.err != nil {
		return nil
	}

	allowed := false
	for n := range s.locks {
		if n == qry.Set {
			allowed = true
			break
		}
	}
	if !allowed {
		s.Fail(errors.New("karigo: can't get collection fron unlocked set"))
		return nil
	}

	col, err := s.node.collection(s.version, qry)
	if err != nil {
		s.Fail(err)
		return nil
	}

	return col
}

// RLock ...
func (s *Snapshot) RLock(set string) {
	if s.err != nil {
		return
	}

	if s.ready {
		s.Fail(errors.New("karigo: snapshot is ready and cannot lock anymore"))
		return
	}
	if _, ok := s.locks[set]; ok {
		s.Fail(errors.New("karigo: set already locked"))
		return
	}

	// err := s.node.RLock(set)
	// if err != nil {
	// 	s.Fail(err)
	// 	return
	// }
	s.locks[set] = false
}

// WLock ...
func (s *Snapshot) WLock(set string) {
	if s.err != nil {
		return
	}

	if s.ready {
		s.Fail(errors.New("karigo: snapshot is ready and cannot lock anymore"))
		return
	}
	if _, ok := s.locks[set]; ok {
		s.Fail(errors.New("karigo: set already locked"))
		return
	}

	// err := s.node.WLock(set)
	// if err != nil {
	// 	s.Fail(err)
	// 	return
	// }
	s.locks[set] = true
}

// Unlock ...
func (s *Snapshot) Unlock(set string) {
	if s.err != nil {
		return
	}

	for i := range s.ops {
		if s.ops[i].Key.Set == set {
			s.Fail(errors.New("karigo: cannot unlock set mentioned in pending operation"))
			return
		}
	}

	var err error
	// TODO Unlock with node or something
	// if w, ok := s.locks[set]; ok {
	// 	if !w {
	// 		err = s.node.RUnlock(set)
	// 	} else {
	// 		err = s.node.WUnlock(set)
	// 	}
	// }
	if err != nil {
		s.Fail(err)
		return
	}
	delete(s.locks, set)
}

// Ready ...
func (s *Snapshot) Ready() {
	if s.err != nil {
		return
	}

	if !s.ready {
		// s.node.snapLock.Unlock()
	}
	s.ready = true
}

// Add ...
func (s *Snapshot) Add(op Op) {
	if s.err != nil {
		return
	}

	// Make sure there is a write-lock for the set
	for n, w := range s.locks {
		if n == op.Key.Set && w {
			s.ops = append(s.ops, op)
			return
		}
	}

	s.Fail(errors.New("karigo: cannot operate on write-unlocked set"))
	return
}

// Flush ...
func (s *Snapshot) Flush() {
	if s.err != nil {
		return
	}

	// err := s.node.Apply(s.ops)
	// if err != nil {
	// 	s.Fail(err)
	// 	return
	// }
}

// Commit ...
func (s *Snapshot) Commit() {
	if s.err != nil {
		return
	}

	s.Ready()
	s.Flush()

	for n := range s.locks {
		s.Unlock(n)
	}
}

// Fail ...
func (s *Snapshot) Fail(err error) {
	if !s.ready {
		s.Ready()
	}
	s.err = err
}
