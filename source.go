package karigo

import (
	"github.com/mfcochauxlaberge/karigo/query"

	"github.com/mfcochauxlaberge/jsonapi"
)

// A Source is the interface used by karigo to query data from a database and
// apply operations to mutate the data.
//
// Each of the interface's methods have comments to explain how to implement a
// Source. See the current implementations for more details.
type Source interface {
	Service

	// Reset wipes all data and brings the underlying database
	// to a clean state.
	//
	// A clean state represents the given schema.
	Reset(*jsonapi.Schema) error

	// NewTx returns a new Tx object.
	NewTx() (query.Tx, error)
}

// source is a thin convenient wrapper for a Source.
type source struct {
	src   Source
	alive bool
	// versions map[string]uint64
	// node     *Node
}
