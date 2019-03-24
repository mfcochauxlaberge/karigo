package karigo

// Journal ...
type Journal struct{}

// Append ...
func (j *Journal) Append(c Entry) error {
	return nil
}

// Get ...
func (j *Journal) Get(id string) (Entry, error) {
	return Entry{}, nil
}

// Last ...
func (j *Journal) Last() (Entry, error) {
	return Entry{}, nil
}

// At ...
func (j *Journal) At(v uint64) (Entry, error) {
	return Entry{}, nil
}

// Tail ...
func (j *Journal) Tail(n uint) ([]Entry, error) {
	return []Entry{}, nil
}

// Range ...
func (j *Journal) Range(n uint) ([]Entry, error) {
	return []Entry{}, nil
}
