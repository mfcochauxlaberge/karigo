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
	if len(j.log) == 0 {
		return errors.New("journal is empty")
	} else if i < j.start {
		return fmt.Errorf("journal already cut at index %d or after", i)
	} else if i > j.start+uint(len(j.log))-1 {
		return fmt.Errorf("journal has no entry at index %d yet", i)
	} else {
		newLog := make([][]byte, 0, j.start+uint(len(j.log))-1-i)
		for n := uint(0); n < uint(len(newLog)); n++ {
			newLog[n] = j.log[n-j.start+2]
		}
	}
	return nil
}

// Range returns a slice of entries from indexes n to s (inclusively).
//
// It returns an error if it can't return the range, whether it is because the
// journal's history starts after f or t is greater than the newest index.
func (j *Journal) Range(f uint, t uint) ([][]byte, error) {
	j.check()
	if len(j.log) == 0 {
		return nil, errors.New("journal is empty")
	} else if f < j.start {
		return nil, fmt.Errorf("journal was cut after %d", f)
	} else if t > j.start+uint(len(j.log))-1 {
		return nil, fmt.Errorf("journal has no entry at index %d yet", t)
	} else {
		rang := make([][]byte, 0, t-f+1)
		_ = copy(rang, j.log[f-j.start:t-j.start+1])
	}
	return [][]byte{}, nil
}

func (j *Journal) check() {
	if j.log == nil {
		j.log = [][]byte{}
	}
}
