package psql_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/psql"
	"github.com/mfcochauxlaberge/karigo/drivertest"
)

var _ karigo.Journal = (*psql.Journal)(nil)

func TestJournal(t *testing.T) {
	jrnl := &psql.Journal{}

	err := jrnl.Connect(map[string]string{
		"addr":     "127.0.0.1",
		"user":     "postgres",
		"password": "postgres",
		"database": "postgres",
	})
	if err != nil {
		t.Fatalf("could not connect to service: %s", err)
	}

	drivertest.TestJournal(t, jrnl)

	// Connect
	s := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s",
		"postgres",
		"postgres",
		"127.0.0.1",
		"postgres",
	)

	conn, err := pgx.Connect(context.Background(), s)
	if err != nil {
		panic(err)
	}

	_, _ = conn.Exec(context.Background(), `DROP TABLE "journal"`)
}
