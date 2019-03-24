package karigo

// Journal ...
type Journal struct{}

// Append ...
func (j *Journal) Append(c []byte) error {
	return nil
}

// Get ...
func (j *Journal) Get(id string) ([]byte, error) {
	return []byte{}, nil
}

// Last ...
func (j *Journal) Last() ([]byte, error) {
	return []byte{}, nil
}

// At ...
func (j *Journal) At(v uint64) ([]byte, error) {
	return []byte{}, nil
}

// Tail ...
func (j *Journal) Tail(n uint) ([][]byte, error) {
	return [][]byte{}, nil
}

// Range ...
func (j *Journal) Range(n uint) ([][]byte, error) {
	return [][]byte{}, nil
}
