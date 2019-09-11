package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJournalSource(t *testing.T) {
	assert := assert.New(t)

	// Empty journal
	journal := &Journal{}

	v, last, err := journal.Last()
	assert.Equal(uint(0), v)
	assert.Equal(([]byte)(nil), last)
	assert.Error(err)

	// Append an entry
	err = journal.Append([]byte("abc"))
	assert.NoError(err)

	v, last, err = journal.Last()
	assert.Equal(uint(0), v)
	assert.Equal([]byte("abc"), last)
	assert.NoError(err)

	// Append an empty entry
	err = journal.Append([]byte{})
	assert.NoError(err)

	v, last, err = journal.Last()
	assert.Equal(uint(1), v)
	assert.Equal([]byte{}, last)
	assert.NoError(err)

	// Append a nil slice (empty entry)
	err = journal.Append(nil)
	assert.NoError(err)

	v, last, err = journal.Last()
	assert.Equal(uint(2), v)
	assert.Equal([]byte{}, last)
	assert.NoError(err)

	// Get specific entries
	entry, err := journal.At(0)
	assert.Equal([]byte("abc"), entry)
	assert.NoError(err)

	entry, err = journal.At(2)
	assert.Equal([]byte{}, entry)
	assert.NoError(err)

	entry, err = journal.At(999)
	assert.Equal(([]byte)(nil), entry)
	assert.Error(err)
}
