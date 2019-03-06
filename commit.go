package karigo

// Commit ...
type Commit struct {
	Version uint64 `json:"ver"`
	Unique  string `json:"uni"`
	User    string `json:"usr"`
	Ops     []Op   `json:"ops"`
}

// Log ...
type Log []Commit
