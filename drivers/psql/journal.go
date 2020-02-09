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

// Connect implements the corresponding method of karigo.Journal.
func (j *Journal) Connect(_ map[string]string) error {
	return nil
}

// Ping implements the corresponding method of karigo.Journal.
func (j *Journal) Ping() bool {
	return true
}

// Append implements the corresponding method of karigo.Journal.
func (j *Journal) Append(c []byte) error {
	return nil
}

// Oldest implements the corresponding method of karigo.Journal.
func (j *Journal) Oldest() (uint, []byte, error) {
	return 0, nil, errors.New("karigo: journal is empty")
}

// Newest implements the corresponding method of karigo.Journal.
func (j *Journal) Newest() (uint, []byte, error) {
	return 0, nil, errors.New("karigo: journal is empty")
}

// At implements the corresponding method of karigo.Journal.
func (j *Journal) At(i uint) ([]byte, error) {
	return nil, nil
}

// Cut implements the corresponding method of karigo.Journal.
func (j *Journal) Cut(i uint) error {
	return nil
}

// Range implements the corresponding method of karigo.Journal.
func (j *Journal) Range(f, t uint) ([][]byte, error) {
	return nil, nil
}
