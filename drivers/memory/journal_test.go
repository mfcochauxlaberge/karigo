package memory_test

import (
	"strconv"
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/memory"

	"github.com/stretchr/testify/assert"
)

var _ = (karigo.Journal)(nil)

func TestJournalSource(t *testing.T) {
	assert := assert.New(t)

	// Empty journal
	var journal karigo.Journal = &memory.Journal{}

	v, last, err := journal.Newest()
	assert.Equal(0, int(v))
	assert.Equal(([]byte)(nil), last)
	assert.Error(err)

	v, last, err = journal.Oldest()
	assert.Equal(0, int(v))
	assert.Equal(([]byte)(nil), last)
	assert.Error(err)

	// Append an entry
	err = journal.Append([]byte("abc"))
	assert.NoError(err)

	v, last, err = journal.Newest()
	assert.Equal(0, int(v))
	assert.Equal([]byte("abc"), last)
	assert.NoError(err)

	// Append an empty entry
	err = journal.Append([]byte{})
	assert.NoError(err)

	v, last, err = journal.Newest()
	assert.Equal(1, int(v))
	assert.Equal([]byte{}, last)
	assert.NoError(err)

	// Append a nil slice (empty entry)
	err = journal.Append(nil)
	assert.NoError(err)

	v, last, err = journal.Newest()
	assert.Equal(2, int(v))
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

	journal := &memory.Journal{}

	// Empty journal
	assert.NoError(journal.Cut(0))
	assert.NoError(journal.Cut(1))

	// One element
	_ = journal.Append([]byte("0"))
	v, entry, _ := journal.Newest()
	assert.Equal(0, int(v))
	assert.Equal([]byte("0"), entry)
	assert.NoError(journal.Cut(0))
	assert.NoError(journal.Cut(0))

	for i := 1; i < 100; i++ {
		data := strconv.Itoa(i)
		_ = journal.Append([]byte(data))
	}

	// Cut after newest index
	assert.NoError(journal.Cut(999))
	v, entry, _ = journal.Oldest()
	assert.Equal(99, int(v))
	assert.Equal([]byte("99"), entry)

	for i := 100; i < 200; i++ {
		data := strconv.Itoa(i)
		_ = journal.Append([]byte(data))
	}

	// Normal cut
	assert.NoError(journal.Cut(150))
	v, entry, _ = journal.Oldest()
	assert.Equal(150, int(v))
	assert.Equal("150", string(entry))

	// Cut before the oldest index
	assert.NoError(journal.Cut(10))
	v, entry, _ = journal.Oldest()
	assert.Equal(150, int(v))
	assert.Equal("150", string(entry))
}

func TestJournalRange(t *testing.T) {
	assert := assert.New(t)

	journal := &memory.Journal{}

	// Empty journal
	rang, err := journal.Range(0, 1)
	assert.Error(err)
	assert.Len(rang, 0)

	journal = &memory.Journal{}

	for i := 100; i < 200; i++ {
		data := strconv.Itoa(i)
		_ = journal.Append([]byte(data))
	}

	_ = journal.Cut(50)

	// Empty range
	rang, err = journal.Range(0, 0)
	assert.NoError(err)
	assert.Len(rang, 0)

	// Empty range
	rang, err = journal.Range(60, 63)
	assert.NoError(err)
	assert.Len(rang, 4)

	// Invalid delimiters
	assert.Panics(func() {
		_, _ = journal.Range(1, 0)
	})

	rang, err = journal.Range(0, 50)
	assert.Error(err)
	assert.Len(rang, 0)

	rang, err = journal.Range(60, 200)
	assert.Error(err)
	assert.Len(rang, 0)
}
