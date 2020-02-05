package psql_test

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/psql"
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

func TestJournalSource(t *testing.T) {
	assert := assert.New(t)

	jrnl, err := psql.NewJournal("postgresql://test:test@localhost")
	assert.NoError(err)

	err = jrnl.Append([]byte("abc"))
	assert.NoError(err)
}
