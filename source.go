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

// source is a thin convenient wrapper for a Source.
type source struct {
	src Source
	// versions map[string]uint64
	// node     *Node
}
