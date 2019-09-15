package karigo

// Journal ...
type Journal interface {
	Append([]byte) error
	First() (uint, []byte, error)
	Last() (uint, []byte, error)
	At(uint) ([]byte, error)
	Cut(uint) error
	Range(uint, uint) ([][]byte, error)
}
