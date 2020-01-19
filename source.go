package karigo

// A Source is the interface used by karigo to query data from a database and
// apply operations to mutate the data.
//
// Each of the interface's methods have comments to explain how to implement a
// Source. See the current implementations for more details.
type Source interface {
	// Reset wipes all data and brings the underlying database
	// to a clean state.
	//
	// A clean state represents the schema returned by FirstSchema.
	Reset() error

	// NewTx returns a new Tx object.
	NewTx() (Tx, error)
}

// source is a thin convenient wrapper for a Source.
type source struct {
	src Source
	// versions map[string]uint64
	// node     *Node
}
