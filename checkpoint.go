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
	tx    SourceTx
	ops   []Op
	undo  []Op
	ready bool

	err error
}

// Resource ...
func (s *Checkpoint) Resource(qry QueryRes) jsonapi.Resource {
	if s.err != nil {
		return nil
	}

	// allowed := false
	// for n := range s.locks {
	// 	if n == qry.Set {
	// 		allowed = true
	// 		break
	// 	}
	// }
	// if !allowed {
	// 	s.Fail(errors.New("karigo: can't get resource fron unlocked set"))
	// 	return nil
	// }

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

	// allowed := false
	// for n := range s.locks {
	// 	if n == qry.Set {
	// 		allowed = true
	// 		break
	// 	}
	// }
	// if !allowed {
	// 	s.Fail(errors.New("karigo: can't get collection fron unlocked set"))
	// 	return nil
	// }

	col, err := s.node.collection(s.version, qry)
	if err != nil {
		s.Fail(err)
		return nil
	}

	return col
}

// RLock ...
func (s *Checkpoint) RLock(set string) {
	// TODO
	// if s.err != nil {
	// 	return
	// }

	// if s.ready {
	// 	s.Fail(errors.New("karigo: checkpoint is ready and cannot lock anymore"))
	// 	return
	// }
	// if _, ok := s.locks[set]; ok {
	// 	s.Fail(errors.New("karigo: set already locked"))
	// 	return
	// }

	// // err := s.node.RLock(set)
	// // if err != nil {
	// // 	s.Fail(err)
	// // 	return
	// // }
	// s.locks[set] = false
}

// WLock ...
func (s *Checkpoint) WLock(set string) {
	// TODO
	// if s.err != nil {
	// 	return
	// }

	// if s.ready {
	// 	s.Fail(errors.New("karigo: checkpoint is ready and cannot lock anymore"))
	// 	return
	// }
	// if _, ok := s.locks[set]; ok {
	// 	s.Fail(errors.New("karigo: set already locked"))
	// 	return
	// }

	// // err := s.node.WLock(set)
	// // if err != nil {
	// // 	s.Fail(err)
	// // 	return
	// // }
	// s.locks[set] = true
}

// Unlock ...
func (s *Checkpoint) Unlock(set string) {
	// TODO
	// if s.err != nil {
	// 	return
	// }

	// for i := range s.ops {
	// 	if s.ops[i].Key.Set == set {
	// 		s.Fail(errors.New("karigo: cannot unlock set mentioned in pending operation"))
	// 		return
	// 	}
	// }

	// var err error
	// // TODO Unlock with node or something
	// // if w, ok := s.locks[set]; ok {
	// // 	if !w {
	// // 		err = s.node.RUnlock(set)
	// // 	} else {
	// // 		err = s.node.WUnlock(set)
	// // 	}
	// // }
	// if err != nil {
	// 	s.Fail(err)
	// 	return
	// }
	// delete(s.locks, set)
}

// Ready ...
func (s *Checkpoint) Ready() {
	// TODO
	// if s.err != nil {
	// 	return
	// }

	// if !s.ready {
	// 	// s.node.snapLock.Unlock()
	// }
	// s.ready = true
}

// Add ...
func (s *Checkpoint) Add(op Op) {
	// TODO
	// if s.err != nil {
	// 	return
	// }

	// // Make sure there is a write-lock for the set
	// for n, w := range s.locks {
	// 	if n == op.Key.Set && w {
	// 		s.ops = append(s.ops, op)
	// 		return
	// 	}
	// }

	// s.Fail(errors.New("karigo: cannot operate on write-unlocked set"))
	// return
}

// Flush ...
func (s *Checkpoint) Flush() {
	// TODO
	// if s.err != nil {
	// 	return
	// }

	// // err := s.node.Apply(s.ops)
	// // if err != nil {
	// // 	s.Fail(err)
	// // 	return
	// // }
}

// Commit ...
func (s *Checkpoint) Commit() {
	// TODO
	// if s.err != nil {
	// 	return
	// }

	// s.Ready()
	// s.Flush()

	// for n := range s.locks {
	// 	s.Unlock(n)
	// }
}

// Fail ...
func (s *Checkpoint) Fail(err error) {
	// TODO
	// if !s.ready {
	// 	s.Ready()
	// }
	// s.err = err
}
