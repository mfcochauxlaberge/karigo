package psql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// NewJournal ...
func NewJournal(c string) (*Journal, error) {
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
//
// Params:
//  - user
//  - password
//  - addr (hostname and port)
func (j *Journal) Connect(params map[string]string) error {
	s := fmt.Sprintf(
		"postgresql://%s:%s@%s",
		params["user"],
		params["password"],
		params["addr"],
	)

	conn, err := pgx.Connect(context.Background(), s)
	if err != nil {
		return nil, err
	}

	return nil
}

// Ping implements the corresponding method of karigo.Journal.
func (j *Journal) Ping() bool {
	if j.conn != nil {
		j.conn.Ping(context.Background())
		return err == nil
	}

	return false
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
