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

// // Append ...
// func (l *Log) Append(c Commit) error {
// 	return nil
// }

// // Get ...
// func (l *Log) Get(id string) (Commit, error) {
// 	return Commit{}, nil
// }

// // Last ...
// func (l *Log) Last() (Commit, error) {
// 	return Commit{}, nil
// }

// // At ...
// func (l *Log) At(v uint64) (Commit, error) {
// 	return Commit{}, nil
// }

// // Tail ...
// func (l *Log) Tail(n uint) ([]Commit, error) {
// 	return []Commit{}, nil
// }

// // Range ...
// func (l *Log) Range(n uint) ([]Commit, error) {
// 	return []Commit{}, nil
// }
