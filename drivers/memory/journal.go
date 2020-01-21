package memory

import (
	"errors"
	"fmt"
	"sync"
)

// Journal is a simple memory-based implementation of karigo.Journal.
//
// It is meant to be used for testing and examples. It is not meant to be used
// in production.
type Journal struct {
	log   [][]byte
	start uint

	m sync.Mutex
}

// Append appends c to the log.
func (j *Journal) Append(c []byte) error {
	j.m.Lock()
	defer j.m.Unlock()
	j.check()

	if c == nil {
		c = []byte{}
	}

	j.log = append(j.log, c)

	return nil
}

// Oldest returns the oldest known entry.
func (j *Journal) Oldest() (uint, []byte, error) {
	j.m.Lock()
	defer j.m.Unlock()
	j.check()

	if len(j.log) > 0 {
		return j.start, j.log[0], nil
	}

	return 0, nil, errors.New("karigo: journal is empty")
}

// Newest returns the newest entry.
func (j *Journal) Newest() (uint, []byte, error) {
	j.m.Lock()
	defer j.m.Unlock()
	j.check()

	if len(j.log) > 0 {
		newest := j.start + uint(len(j.log)) - 1
		return newest, j.log[len(j.log)-1], nil
	}

	return 0, nil, errors.New("karigo: journal is empty")
}

// At returns the entry indexed at i.
func (j *Journal) At(i uint) ([]byte, error) {
	j.m.Lock()
	defer j.m.Unlock()
	j.check()

	if i < j.start || i > j.start+uint(len(j.log))-1 {
		return nil, fmt.Errorf("karigo: index %d does not exist", i)
	}

	return j.log[i-j.start], nil
}

// Cut removes all entries from the oldest one to the one at i minus one.
//
// If i is lower than the oldest known index, nothing gets cut. If i is greater
// than the newest index, i will be interpreted as the newest index, and
// therefore everything will be cut except the latest index, leaving a journal
// of length one.
func (j *Journal) Cut(i uint) error {
	j.m.Lock()
	defer j.m.Unlock()
	j.check()

	// If i is 0, there is nothing to cut.
	if i == 0 {
		return nil
	}

	// If the log is empty, there is nothing to cut.
	if len(j.log) == 0 {
		return nil
	}

	// If i is before or equal to the oldest known
	// entry, there is nothing to cut.
	if i <= j.start {
		return nil
	}

	// If i is after or equal to the newest entry,
	// cut everything except the newest entry.
	if i > j.start+uint(len(j.log)) {
		j.start = j.start + uint(len(j.log)) - 1
		j.log = [][]byte{j.log[len(j.log)-1]}

		return nil
	}

	// Here, i must be somewhere between the oldest
	// entry and the newest one.
	l := j.start + uint(len(j.log)) - 1 - i
	newLog := make([][]byte, l)

	for n := uint(0); n < uint(len(newLog)-1); n++ {
		newLog[n] = j.log[i-j.start]
	}

	j.start = i
	j.log = newLog

	return nil
}

// Range returns a slice of entries from indexes f to t (inclusively).
//
// It returns an error if it can't return the range, whether it is because the
// journal's history starts after f or t is greater than the newest index.
func (j *Journal) Range(f, t uint) ([][]byte, error) {
	j.m.Lock()
	defer j.m.Unlock()
	j.check()

	if f > t {
		panic("f > t")
	}

	switch {
	case len(j.log) == 0:
		return nil, errors.New("journal is empty")
	case f == t:
		return [][]byte{}, nil
	case f < j.start:
		return nil, fmt.Errorf("journal was cut after %d", f)
	case t > j.start+uint(len(j.log))-1:
		return nil, fmt.Errorf("journal has no entry at index %d yet", t)
	default:
		rang := make([][]byte, t-f+1)
		_ = copy(rang, j.log[f-j.start:t-j.start+1])

		return rang, nil
	}
}

func (j *Journal) check() {
	if j.log == nil {
		j.log = [][]byte{}
	}
}
