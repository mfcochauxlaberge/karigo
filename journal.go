package karigo

// Journal ...
type Journal interface {
	Append([]byte) error
	Oldest() (uint, []byte, error)
	Newest() (uint, []byte, error)
	At(uint) ([]byte, error)
	Cut(uint) error
	Range(uint, uint) ([][]byte, error)
}
