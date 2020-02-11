package psql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// Journal is the PostgreSQL implementation of karigo.Journal.
type Journal struct {
	conn *pgx.Conn
}

// Connect implements the corresponding method of karigo.Journal.
//
// Params:
//  - addr (hostname and port)
//  - user
//  - password
//  - database
func (j *Journal) Connect(params map[string]string) error {
	s := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s",
		params["user"],
		params["password"],
		params["addr"],
		params["database"],
	)

	conn, err := pgx.Connect(context.Background(), s)
	if err != nil {
		return err
	}

	j.conn = conn

	return nil
}

// Ping implements the corresponding method of karigo.Journal.
func (j *Journal) Ping() bool {
	if j.conn != nil {
		err := j.conn.Ping(context.Background())
		return err == nil
	}

	return false
}

// Reset implements the corresponding method of karigo.Journal.
func (j *Journal) Reset() error {
	if j.conn == nil {
		return errors.New("journal not connected")
	}

	_, err := j.conn.Exec(context.Background(), `TRUNCATE "karigo_journal"`)
	if err != nil {
		return err
	}

	return nil
}

// Append implements the corresponding method of karigo.Journal.
func (j *Journal) Append(c []byte) error {
	if j.conn == nil {
		return errors.New("journal not connected")
	}

	// tag, err := j.conn.Exec(
	// 	context.Background(),
	// 	`INSERT INTO $1 (index, entry) VALUES ($2, $3)`,
	// )

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
