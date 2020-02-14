package psql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// Journal is the PostgreSQL implementation of karigo.Journal.
type Journal struct {
	index uint
	conn  *pgx.Conn
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
		return fmt.Errorf("could not connect to postgresql: %s", err)
	}

	// Create table if it does not already exist.
	row := conn.QueryRow(
		context.Background(),
		`
		SELECT "table_name"
		FROM "information_schema"."tables"
		WHERE "table_name" = 'journal'
		`,
	)

	var name string
	err = row.Scan(&name)
	if err != nil && err != pgx.ErrNoRows {
		return fmt.Errorf("journal table does not exist")
	}

	if name != "journal" {
		_, err = conn.Exec(
			context.Background(),
			`
			CREATE TABLE "journal" (
				"index" BIGINT PRIMARY KEY,
				"entry" BYTEA
			)
			`,
		)
		if err != nil {
			return err
		}
	}

	// Get latest index.
	row = conn.QueryRow(
		context.Background(),
		`SELECT "index" FROM "journal" ORDER BY "index" DESC`,
	)

	var index uint
	err = row.Scan(&index)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	j.index = index
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

	_, err := j.conn.Exec(context.Background(), `TRUNCATE "journal"`)
	if err != nil {
		return err
	}

	j.index = 0

	return nil
}

// Append implements the corresponding method of karigo.Journal.
func (j *Journal) Append(c []byte) error {
	if j.conn == nil {
		return errors.New("journal not connected")
	}

	_, err := j.conn.Exec(
		context.Background(),
		`INSERT INTO "journal" (index, entry) VALUES ($1, $2)`,
		j.index, c,
	)
	if err != nil {
		return err
	}

	j.index++

	return nil
}

// Oldest implements the corresponding method of karigo.Journal.
func (j *Journal) Oldest() (uint, []byte, error) {
	row := j.conn.QueryRow(
		context.Background(),
		`SELECT "index", "entry" FROM "journal" ORDER BY "index" ASC`,
	)

	var (
		index uint
		entry []byte
	)
	err := row.Scan(&index, &entry)
	if err == pgx.ErrNoRows {
		return 0, nil, errors.New("karigo: journal is empty")
	} else if err != nil {
		return 0, nil, err
	}

	if entry == nil {
		entry = []byte{}
	}

	return index, entry, nil
}

// Newest implements the corresponding method of karigo.Journal.
func (j *Journal) Newest() (uint, []byte, error) {
	row := j.conn.QueryRow(
		context.Background(),
		`SELECT "index", "entry" FROM "journal" ORDER BY "index" DESC`,
	)

	var (
		index uint
		entry []byte
	)
	err := row.Scan(&index, &entry)
	if err == pgx.ErrNoRows {
		return 0, nil, errors.New("karigo: journal is empty")
	} else if err != nil {
		return 0, nil, err
	}

	if entry == nil {
		entry = []byte{}
	}

	return index, entry, nil
}

// At implements the corresponding method of karigo.Journal.
func (j *Journal) At(i uint) ([]byte, error) {
	row := j.conn.QueryRow(
		context.Background(),
		`SELECT "entry" FROM "journal" WHERE index = $1`,
		i,
	)

	var entry []byte
	err := row.Scan(&entry)
	if err == pgx.ErrNoRows {
		return nil, errors.New("karigo: index does not exist")
	} else if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	if entry == nil {
		entry = []byte{}
	}

	return entry, nil
}

// Cut implements the corresponding method of karigo.Journal.
func (j *Journal) Cut(i uint) error {
	if i > j.index {
		i = j.index
	}

	_, err := j.conn.Exec(
		context.Background(),
		`DELETE FROM "journal" WHERE "index" < $1`,
		i,
	)
	if err != nil {
		return err
	}

	return nil
}

// Range implements the corresponding method of karigo.Journal.
func (j *Journal) Range(f, t uint) ([][]byte, error) {
	return nil, nil
}
