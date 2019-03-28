package karigo

// Journal ...
type Journal struct{}

// Append ...
func (j *Journal) Append(c []byte) error {
	return nil
}

// Last ...
func (j *Journal) Last() (uint, []byte, error) {
	return 0, []byte{}, nil
}

// At ...
func (j *Journal) At(i uint64) ([]byte, error) {
	return []byte{}, nil
}

// Cut ...
func (j *Journal) Cut(i uint64) error {
	return nil
}

// Range ...
func (j *Journal) Range(n uint, s uint) ([][]byte, error) {
	return [][]byte{}, nil
}
