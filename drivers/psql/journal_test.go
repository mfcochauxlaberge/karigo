package psql_test

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/psql"
	"github.com/mfcochauxlaberge/karigo/drivertest"

	"github.com/jackc/pgx/v4"
)

var _ karigo.Journal = (*psql.Journal)(nil)

func TestJournal(t *testing.T) {
	// Connect
	s := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s",
		"test",
		"test",
		"127.0.0.1",
		"test",
	)

	conn, err := pgx.Connect(context.Background(), s)
	if err != nil {
		panic(err)
	}

	// Add a schema with a random name.
	// This is to isolate tests so that they
	// can run concurrently.
	rand.Seed(time.Now().UnixNano())
	schema := strconv.Itoa(rand.Intn(99999))

	_, err = conn.Exec(context.Background(), "CREATE SCHEMA test_"+schema)
	if err != nil {
		panic(err)
	}

	jrnl := &psql.Journal{}

	err = jrnl.Connect(map[string]string{
		"addr":     "127.0.0.1",
		"user":     "test",
		"password": "test",
		"database": "test",
		"schema":   schema,
	})
	if err != nil {
		t.Skipf("psql journal test skipped: %s", err)
	}

	drivertest.TestJournal(t, jrnl)

	// Drop schema
	_, err = conn.Exec(context.Background(), "DROP SCHEMA test_"+schema)
	if err != nil {
		panic(err)
	}
}
