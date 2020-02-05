package psql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

// NewJournal ...
func NewJournal(c string) (*Journal, error) {
	conn, err := pgx.Connect(context.Background(), c)
	if err != nil {
		return nil, err
	}

	jrnl := &Journal{
		conn: conn,
	}

	return jrnl, nil
}

// Journal is the PostgreSQL implementation of karigo.Journal.
type Journal struct {
	conn *pgx.Conn
}

// Append appends c to the log.
func (j *Journal) Append(c []byte) error {
	return nil
}

// Oldest returns the oldest known entry.
func (j *Journal) Oldest() (uint, []byte, error) {
	return 0, nil, errors.New("karigo: journal is empty")
}

// Newest returns the newest entry.
func (j *Journal) Newest() (uint, []byte, error) {
	return 0, nil, errors.New("karigo: journal is empty")
}

// At returns the entry indexed at i.
func (j *Journal) At(i uint) ([]byte, error) {
	return nil, nil
}

// Cut removes all entries from the oldest one to the one at i minus one.
//
// If i is lower than the oldest known index, nothing gets cut. If i is greater
// than the newest index, i will be interpreted as the newest index, and
// therefore everything will be cut except the latest index, leaving a journal
// of length one.
func (j *Journal) Cut(i uint) error {
	return nil
}

// Range returns a slice of entries from indexes f to t (inclusively).
//
// It returns an error if it can't return the range, whether it is because the
// journal's history starts after f or t is greater than the newest index.
func (j *Journal) Range(f, t uint) ([][]byte, error) {
	return nil, nil
}
