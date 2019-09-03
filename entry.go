package karigo

// Entry ...
type Entry struct {
	Version uint64 `json:"ver"`
	ID      string `json:"id"`
	User    string `json:"user"`
	Method  string `json:"method"`
	URL     string `json:"url"`
	Body    []byte `json:"body"`
	Ops     []Op   `json:"ops"`
}
