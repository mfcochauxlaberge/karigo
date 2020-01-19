package karigo

import (
	"github.com/mfcochauxlaberge/jsonapi"
)

// A Source is the interface used by karigo to query data from a database and
// apply operations to mutate the data.
//
// Since the state of a karigo application is based on a deterministic log of
// events, all implementations must rigorously follow the rules to avoid
// corrupting the data. Given an identical list of operations, two sources must
// always produce the same output.
//
// Each of the interface's methods have comments to explain how to implement a
// Source. See the current implementations for more details.
type Source interface {
	// Reset wipes all data and brings the underlying database
	// to a clean state.
	//
	// A clean state represents the schema returned by FirstSchema.
	Reset() error

	// Resource returns the resource defined by QueryRes if found.
	//
	// In all other cases, an error is returned and the resource
	// is nil.
	Resource(QueryRes) (jsonapi.Resource, error)

	// Collection returns the collection of resources defined by
	// QueryCol.
	//
	// An empty collection and no error are returned even if no
	// resources fall under the query.
	//
	// In all other cases, an error is returned and the collection
	// is nil.
	Collection(QueryCol) (jsonapi.Collection, error)

	// Apply applies the operations to the underlying database.
	//
	// The operations must not be committed and the implementation
	// must be ready for a possible rollback.
	//
	// Those operations must also be considered for future calls
	// of the Resource and Collection methods.
	Apply([]Op) error

	// Commit commits all the operations to the underlying database.
	//
	// Either all operations are permanently saved for future requests
	// or none of them are and an error is returned.
	Commit() error

	// Rollback rollbacks all applied but uncommitted operations.
	Rollback() error
}

// source is a thin convenient wrapper for a Source.
type source struct {
	src Source
	// versions map[string]uint64
	// node     *Node
}
