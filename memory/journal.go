package memory

import (
	"errors"
	"fmt"
)

// Journal is a simple memory-based implementation of karigo.Journal.
//
// It is meant to be used for testing and examples. It is not meant to be used
// in production.
type Journal struct {
	log   [][]byte
	start uint
}

// Append appends c to the log.
func (j *Journal) Append(c []byte) error {
	j.check()
	if c == nil {
		c = []byte{}
	}
	j.log = append(j.log, c)
	return nil
}

// Last returns the newest entry.
func (j *Journal) Last() (uint, []byte, error) {
	j.check()
	if len(j.log) > 0 {
		last := j.start + uint(len(j.log)) - 1
		return last, j.log[len(j.log)-1], nil
	}
	return 0, nil, errors.New("karigo: journal is empty")
}

// At returns the entry indexed at i.
func (j *Journal) At(i uint) ([]byte, error) {
	j.check()
	if i < j.start || i > j.start+uint(len(j.log))-1 {
		return nil, fmt.Errorf("karigo: index %d does not exist", i)
	}
	return j.log[i-j.start], nil
}

// Cut removes all entries from the oldest one to the one at index i.
//
// It returns an error if i is greater than the newest index.
func (j *Journal) Cut(i uint) error {
	j.check()
	if i > j.start+uint(len(j.log)) {

	} else if i < j.start {
		// Already cut
		return nil
	} else {

	}
	return nil
}

// Range returns a slice of entries from indexes n to s (inclusively).
//
// It returns an error if it can't return the ranger, whether it is because the
// journal's history starts after n or s is greater than the newest index.
func (j *Journal) Range(n uint, s uint) ([][]byte, error) {
	j.check()
	// TODO
	return [][]byte{}, nil
}

func (j *Journal) check() {
	if j.log == nil {
		j.log = [][]byte{}
	}
}
