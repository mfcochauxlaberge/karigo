package memory

import "fmt"

// Journal ...
type Journal struct {
	log   [][]byte
	start uint
}

// Append ...
func (j *Journal) Append(c []byte) error {
	j.check()
	j.log = append(j.log, c)
	return nil
}

// Oldest ...
func (j *Journal) Oldest() (uint, []byte, error) {
	return 0, nil, nil
}

// Newest ...
func (j *Journal) Newest() (uint, []byte, error) {
	j.check()
	if len(j.log) > 0 {
		last := j.start + uint(len(j.log)) - 1
		return last, j.log[len(j.log)-1], nil
	}
	return 0, nil, nil
}

// At ...
func (j *Journal) At(i uint) ([]byte, error) {
	j.check()
	if i < j.start || i > j.start+uint(len(j.log))-1 {
		return nil, fmt.Errorf("karigo: index %d does not exist", i)
	}
	return j.log[i-j.start], nil
}

// Cut ...
func (j *Journal) Cut(i uint) error {
	j.check()
	// TODO
	return nil
}

// Range ...
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
