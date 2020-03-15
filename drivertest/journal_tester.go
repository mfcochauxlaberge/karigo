package drivertest

import (
	"strconv"
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/stretchr/testify/assert"
)

// TestJournal ...
func TestJournal(t *testing.T, jrnl karigo.Journal) {
	testBasic(t, jrnl)
	testCut(t, jrnl)
	testRange(t, jrnl)
}

func testBasic(t *testing.T, jrnl karigo.Journal) {
	assert := assert.New(t)

	err := jrnl.Reset()
	assert.NoError(err)

	v, last, err := jrnl.Newest()
	assert.Equal(0, int(v))
	assert.Equal(([]byte)(nil), last)
	assert.Error(err)

	v, last, err = jrnl.Oldest()
	assert.Equal(0, int(v))
	assert.Equal(([]byte)(nil), last)
	assert.Error(err)

	// Append an entry
	err = jrnl.Append([]byte("abc"))
	assert.NoError(err)

	v, last, err = jrnl.Newest()
	assert.Equal(0, int(v))
	assert.Equal([]byte("abc"), last)
	assert.NoError(err)

	// Append an empty entry
	err = jrnl.Append([]byte{})
	assert.NoError(err)

	v, last, err = jrnl.Newest()
	assert.Equal(1, int(v))
	assert.Equal([]byte{}, last)
	assert.NoError(err)

	// Append a nil slice (empty entry)
	err = jrnl.Append(nil)
	assert.NoError(err)

	v, last, err = jrnl.Newest()
	assert.Equal(2, int(v))
	assert.Equal([]byte{}, last)
	assert.NoError(err)

	// Get specific entries
	entry, err := jrnl.At(0)
	assert.Equal([]byte("abc"), entry)
	assert.NoError(err)

	entry, err = jrnl.At(2)
	assert.Equal([]byte{}, entry)
	assert.NoError(err)

	entry, err = jrnl.At(999)
	assert.Equal(([]byte)(nil), entry)
	assert.Error(err)
}

func testCut(t *testing.T, jrnl karigo.Journal) {
	assert := assert.New(t)

	err := jrnl.Reset()
	assert.NoError(err)

	// Empty journal
	assert.NoError(jrnl.Cut(0))
	assert.NoError(jrnl.Cut(1))

	// One element
	_ = jrnl.Append([]byte("0"))
	v, entry, _ := jrnl.Newest()
	assert.Equal(0, int(v))
	assert.Equal([]byte("0"), entry)
	assert.NoError(jrnl.Cut(0))
	assert.NoError(jrnl.Cut(0))

	for i := 1; i < 100; i++ {
		data := strconv.Itoa(i)
		_ = jrnl.Append([]byte(data))
	}

	v, entry, _ = jrnl.Newest()
	assert.Equal(99, int(v))
	assert.Equal([]byte("99"), entry)

	// Cut after newest index
	assert.NoError(jrnl.Cut(999))
	v, entry, _ = jrnl.Newest()
	assert.Equal(99, int(v))
	assert.Equal([]byte("99"), entry)

	for i := 100; i < 200; i++ {
		data := strconv.Itoa(i)
		_ = jrnl.Append([]byte(data))
	}

	// Normal cut
	assert.NoError(jrnl.Cut(150))
	v, entry, _ = jrnl.Oldest()
	assert.Equal(150, int(v))
	assert.Equal("150", string(entry))

	// Cut before the oldest index
	assert.NoError(jrnl.Cut(10))
	v, entry, _ = jrnl.Oldest()
	assert.Equal(150, int(v))
	assert.Equal("150", string(entry))
}

func testRange(t *testing.T, jrnl karigo.Journal) {
	assert := assert.New(t)

	err := jrnl.Reset()
	assert.NoError(err)

	// Empty journal
	rang, err := jrnl.Range(0, 1)
	assert.Error(err)
	assert.Len(rang, 0)

	for i := 100; i < 200; i++ {
		data := strconv.Itoa(i)
		_ = jrnl.Append([]byte(data))
	}

	_ = jrnl.Cut(50)

	// Empty range
	rang, err = jrnl.Range(0, 0)
	assert.NoError(err)
	assert.Len(rang, 0)

	// Empty range
	rang, err = jrnl.Range(60, 63)
	assert.NoError(err)
	assert.Len(rang, 4)

	// Invalid delimiters
	assert.Panics(func() {
		_, _ = jrnl.Range(1, 0)
	})

	rang, err = jrnl.Range(0, 50)
	assert.Error(err)
	assert.Len(rang, 0)

	rang, err = jrnl.Range(60, 200)
	assert.Error(err)
	assert.Len(rang, 0)
}
