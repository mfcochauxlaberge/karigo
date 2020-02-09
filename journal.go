package karigo

// Journal ...
type Journal interface {
	Service

	// Append appends an entry to the journal.
	Append([]byte) error

	// Oldest returns the oldest known entry.
	Oldest() (uint, []byte, error)

	// Newest returns the newest entry.
	Newest() (uint, []byte, error)

	// At returns the entry indexed at i, or none if it
	// does not exist.
	At(i uint) ([]byte, error)

	// Cut removes all entries from the oldest one to
	// the one at i minus one.
	//
	// If i is lower than the oldest known index,
	// nothing gets cut. If i is greater than the newest
	// index, it will be interpreted as the newest index,
	// and therefore everything will be cut except i,
	// leaving a journal of length one.
	Cut(i uint) error

	// Range returns a slice of entries from indexes f
	// to t (inclusively).
	//
	// It returns an error if it can't return the range,
	// whether it is because the journal's history starts
	// after f or t is greater than the newest index.
	Range(f uint, t uint) ([][]byte, error)
}

// source is a thin convenient wrapper for a Journal.
type journal struct {
	jrnl Journal
}
