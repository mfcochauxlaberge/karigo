package psql_test

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/psql"
	"github.com/mfcochauxlaberge/karigo/drivertest"
	"github.com/stretchr/testify/assert"
)

var _ karigo.Journal = (*psql.Journal)(nil)

// func Main(m *testing.M) {
// 	conn, err := pgx.Connect(context.Background(), c)
// 	if err != nil {
// 		panic("cannot connect to PostgreSQL instance")
// 	}

// 	// conn.Exec()

// 	m.Run()
// }

func TestJournal(t *testing.T) {
	assert := assert.New(t)

	jrnl := &psql.Journal{}

	err := jrnl.Connect(map[string]string{
		"addr":     "127.0.0.1",
		"user":     "test",
		"password": "test",
		"database": "test_journal",
	})
	if err != nil {
		t.Skipf("psql journal test skipped: %s", err)
	}

	err = jrnl.Reset()
	assert.NoError(err)

	assert.Equal(1, 2)

	drivertest.TestJournal(t, jrnl)
}
