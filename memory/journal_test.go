package memory_test

import (
	"strconv"
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/memory"

	"github.com/stretchr/testify/assert"
)

var _ = (karigo.Journal)(nil)

func TestJournalSource(t *testing.T) {
	assert := assert.New(t)

	// Empty journal
	var journal karigo.Journal = &memory.Journal{}

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

	var journal karigo.Journal = &memory.Journal{}
	assert.NoError(journal.Cut(0))

	_ = journal.Append([]byte("0"))
	v, entry, _ := journal.Last()
	assert.Equal(uint(0), v)
	assert.Equal([]byte("0"), entry)
	assert.NoError(journal.Cut(0))
	assert.NoError(journal.Cut(0))

	for i := 0; i < 100; i++ {
		data := strconv.Itoa(i)
		_ = journal.Append([]byte(data))
	}

	assert.NoError(journal.Cut(999))
	v, entry, _ = journal.First()
	assert.Equal(uint(0), v)
	assert.Equal([]byte("0"), entry)

	assert.NoError(journal.Cut(10))
	v, entry, _ = journal.First()
	assert.Equal(uint(11), v)
	assert.Equal([]byte("11"), entry)

	assert.NoError(journal.Cut(10))
	v, entry, _ = journal.First()
	assert.Equal(uint(11), v)
	assert.Equal([]byte("11"), entry)
}
