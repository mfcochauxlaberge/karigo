package karigo

import (
	"errors"
)

var (
	// ErrCantDo should be returned by a source that is given a valid query
	// that it is unable to do.
	ErrCantDo = errors.New("can't do the given query")

	// ErrNotFound is returned when an endpoint does not exist.
	ErrNotFound = errors.New("resource not found")

	// ErrUnexpected is returned when an unexpected error occurs.
	ErrUnexpected = errors.New("unexpected error")

	// ErrNotImplemented is returned when an endpoint exists but is not
	// implemented.
	ErrNotImplemented = errors.New("not implemented")
)
