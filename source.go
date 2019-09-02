package karigo

import (
	"github.com/mfcochauxlaberge/jsonapi"
)

// Source ...
type Source interface {
	Reset() error

	Resource(QueryRes) (jsonapi.Resource, error)
	Collection(QueryCol) (jsonapi.Collection, error)
	Apply([]Op) error
}

// DirectSource ...
// type DirectSource interface {
// 	Source

// 	Apply(ops []Op) error
// }

// source is a thin convenient wrapper for a Source.
type source struct {
	src Source
	// versions map[string]uint64
	// node     *Node
}

// // version returns the lowest version from all the sets.
// func (s *source) version() uint64 {
// 	mv := uint64(math.MaxUint64)
// 	for _, v := range s.versions {
// 		if v < mv {
// 			mv = v
// 		}
// 	}
// 	return mv
// }

// // SourceTx ...
// type SourceTx interface {
// 	Apply([]Op) error
// 	Rollback() error
// 	Commit() error
// }
