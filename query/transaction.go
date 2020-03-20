package query

import (
	"github.com/mfcochauxlaberge/jsonapi"
)

// A Tx defines the interface of a transaction and is meant to be used by
// implementations of the Source interface.
//
// Since the state of a karigo application is based on a deterministic log of
// events, all implementations must rigorously follow the rules to avoid
// corrupting the data. Given an identical list of operations, two
// implementations must always produce the same output.
//
// Each of the interface's methods have comments that explain how to properly
// implement Tx. See the current implementations for more details.
type Tx interface {
	// Resource returns the resource defined by Res if found.
	//
	// In all other cases, an error is returned and the resource
	// is nil.
	Resource(Res) (jsonapi.Resource, error)

	// Collection returns the collection of resources defined by Col.
	//
	// An empty collection and no error are returned even if no
	// resources fall under the query.
	//
	// In all other cases, an error is returned and the collection
	// is nil.
	Collection(Col) (jsonapi.Collection, error)

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
