package karigo

import (
	"math"
)

// Source ...
type Source interface {
	Reset() error

	Resource(QueryRes) (map[string]interface{}, error)
	Collection(QueryCol) ([]map[string]interface{}, error)

	Begin() (SourceTx, error)
}

// source is a thin convenient wrapper for a Source.
type source struct {
	src      Source
	versions map[string]uint64
	// node     *Node
}

// version returns the lowest version from all the sets.
func (s *source) version() uint64 {
	mv := uint64(math.MaxUint64)
	for _, v := range s.versions {
		if mv == 0 {
			mv = v
		}
		if v < mv {
			mv = v
		}
	}
	return mv
}

// func (s *source) run() {
// 	for {
// 		time.Sleep(2 * time.Second)
// 	}
// }

// SourceTx ...
type SourceTx interface {
	Apply([]Op) error
	Rollback() error
	Commit() error
}
