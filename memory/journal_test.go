package memory

import (
	"strconv"
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

func TestJournalCut(t *testing.T) {
	assert := assert.New(t)

	journal := &Journal{}
	assert.EqualError(journal.Cut(0), "journal is empty")

	_ = journal.Append([]byte("0"))
	v, entry, _ := journal.Last()
	assert.Equal(uint(0), v)
	assert.Equal([]byte("0"), entry)
	assert.NoError(journal.Cut(0))
	assert.EqualError(journal.Cut(0), "journal is empty")

	for i := 0; i < 100; i++ {
		data := strconv.Itoa(i)
		_ = journal.Append([]byte(data))
	}
	assert.EqualError(
		journal.Cut(999),
		"journal has no entry at index 999 yet",
	)
	_ = journal.Cut(10)
	assert.EqualError(
		journal.Cut(10),
		"journal already cut at index 10 or after",
	)
}
