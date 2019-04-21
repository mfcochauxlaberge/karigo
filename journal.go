package karigo

// Journal ...
type Journal interface {
	Append([]byte) error
	Last() (uint64, []byte, error)
	At(uint64) ([]byte, error)
	Cut(uint64) error
	Range(uint64, uint) ([][]byte, error)
}
