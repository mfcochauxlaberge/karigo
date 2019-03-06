package karigo

// Source ...
type Source interface {
	Reset() error

	Resource(QueryRes) (map[string]interface{}, error)
	Collection(QueryCol) ([]map[string]interface{}, error)

	Begin() (SourceTx, error)
}

// source ...
type source struct {
	name     string
	src      Source
	version  uint64
	versions map[string]uint64
	// availability [2]map[string]bool
	node *Node
}

// func (s *source) run() {
// 	for {
// 		time.Sleep(2 * time.Second)
// 	}
// }

// SourceTx ...
type SourceTx interface {
	Apply([]Op) error
	Rollback() error
	Commit() error
}
