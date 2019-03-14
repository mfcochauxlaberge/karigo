package karigo

import (
	"errors"
	"sync"

	"github.com/mfcochauxlaberge/jsonapi"
)

// NewNode ...
func NewNode(src Source) *Node {
	node := &Node{
		log: []Commit{},
		main: source{
			src: src,
		},

		currSchema: NewSchema(),
		nextSchema: NewSchema(),

		snapLock: &sync.Mutex{},
		locks:    map[string]*sync.RWMutex{},
	}

	return node
}

// Node ...
type Node struct {
	// Run
	log  Log
	main source

	// Schema
	currSchema *Schema
	nextSchema *Schema

	// Requests
	snapLock *sync.Mutex
	locks    map[string]*sync.RWMutex

	// Error
	err error

	// Internal
	sync.Mutex
}

// Run ...
func (n *Node) Run() error {
	n.Lock()
	n.Unlock()

	for n.err == nil {
		select {}
	}

	return n.err
}

// // AddSource ...
// func (n *Node) AddSource(name string, src Source) error {
// 	n.Lock()
// 	defer n.Unlock()

// 	if n.sources == nil {
// 		n.sources = []source{}
// 	}

// 	if len(name) == 0 {
// 		return errors.New("karigo: source name cannot be empty")
// 	}

// 	for i := range n.sources {
// 		if n.sources[i].name == name {
// 			return errors.New("karigo: source name already used")
// 		}
// 	}

// 	n.sources = append(n.sources, source{
// 		name: name,
// 		src:  src,
// 	})

// 	// go n.sources[len(n.sources)-1].run()

// 	return nil
// }

// Handle ...
func (n *Node) Handle(req *Request) *Response {
	n.Lock()
	defer n.Unlock()

	var tx Tx
	switch req.Method {
	case GET:
		tx = n.currSchema.getFuncs[req.URL.ResType]
	case POST:
		tx = n.currSchema.createFuncs[req.URL.ResType]
	case PATCH:
		tx = n.currSchema.updateFuncs[req.URL.ResType]
	case DELETE:
		tx = n.currSchema.deleteFuncs[req.URL.ResType]
	}
	if tx == nil {
		tx = TxNotImplemented
	}

	snap := &Snapshot{
		node:  n,
		locks: map[string]bool{},
		ops:   []Op{},
	}

	n.snapLock.Lock()
	tx(snap)
	snap.Commit()

	res := &Response{}

	if snap.err != nil {
		var jaErr jsonapi.Error
		switch snap.err {
		case ErrNotImplemented:
			jaErr = jsonapi.NewErrNotImplemented()
		default:
			jaErr = jsonapi.NewErrInternalServerError()
		}
		res.Errors = []jsonapi.Error{jaErr}
	} else {
		res.Data = nil

		// Aggregate operations
		// Commit to log
		// If success:
		//   * Commit sources
		//   * Return success
		// If failure:
		//   * Rollback sources
		//   * Try to transfer request to master node
		//   * Return response
	}

	return res
}

// Close ...
func (n *Node) Close() error {
	n.Lock()
	defer n.Unlock()
	if n.err == nil {
		n.err = errors.New("karigo: node execution has been closed")
		return n.err
	}
	return n.err
}

// Shutdown ...
func (n *Node) Shutdown() error {
	n.Lock()
	defer n.Unlock()
	if n.err == nil {
		n.err = errors.New("karigo: node execution has been shut down")
		return n.err
	}
	return n.err
}

// Resource ...
func (n *Node) Resource(qry QueryRes) (jsonapi.Resource, error) {
	// for i := range n.sources {
	// 	if n.sources[i].versions[qry.Set] == version {
	// 		_, err := n.sources[i].src.Resource(qry)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 	}
	// }

	return nil, errors.New("karigo: no source could handle the query")
}

// Collection ...
func (n *Node) Collection(qry QueryCol) ([]jsonapi.Resource, error) {
	// TODO Validate the query?
	// TODO Complete the sorting rule

	// for i := range n.sources {
	// 	if n.sources[i].versions[qry.Set] == version {
	// 		_, err := n.sources[i].src.Collection(qry)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 	}
	// }

	return n.main.src.Collection(qry)
}

// RLock ...
func (n *Node) RLock(set string) error {
	if _, ok := n.locks[set]; !ok {
		return errors.New("karigo: cannot read-lock inexistent set")
	}
	n.locks[set].RLock()
	return nil
}

// WLock ...
func (n *Node) WLock(set string) error {
	if _, ok := n.locks[set]; !ok {
		return errors.New("karigo: cannot write-lock inexistent set")
	}
	n.locks[set].Lock()
	return nil
}

// RUnlock ...
func (n *Node) RUnlock(set string) error {
	if _, ok := n.locks[set]; !ok {
		return errors.New("karigo: cannot read-unlock inexistent set")
	}
	n.locks[set].Unlock()
	return nil
}

// WUnlock ...
func (n *Node) WUnlock(set string) error {
	if _, ok := n.locks[set]; !ok {
		return errors.New("karigo: cannot write-unlock inexistent set")
	}
	n.locks[set].Unlock()
	return nil
}

// // Commit ...
// func (n *Node) Commit(version uint64) error {
// 	for i := range n.sources {
// 		if n.sources[i].versions["abc"] == version-1 {
// 			n.sources[i].versions["abc"] = version
// 		}
// 	}

// 	return nil
// }

// Apply ...
// func (n *Node) Apply(ops []Op) error {
// 	for i := range n.sources {
// 		err := n.sources[i].src.Apply(ops)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return errors.New("karigo: an operation could not be executed")
// }

// func (n *Node) handleRequest(in chan *Request) error {
// 	for {
// 		select {
// 		case req := <-in:
// 			var tx Tx
// 			if req.Method == GET {
// 				tx = n.currSchema.getFuncs[req.URL.ResType]
// 			} else if req.Method == POST {
// 				tx = n.currSchema.createFuncs[req.URL.ResType]
// 			} else if req.Method == PATCH {
// 				tx = n.currSchema.updateFuncs[req.URL.ResType]
// 			} else if req.Method == DELETE {
// 				tx = n.currSchema.deleteFuncs[req.URL.ResType]
// 			}
// 			if tx == nil {
// 				tx = TxNotImplemented
// 			}

// 			snap := &Snapshot{
// 				node:  n,
// 				locks: map[string]bool{},
// 				ops:   []Op{},
// 			}

// 			n.snapLock.Lock()
// 			tx(snap) // TODO Provide fields
// 			snap.Commit()

// 			if snap.err != nil {
// 				jaErr := jsonapi.NewErrNotImplemented()
// 				req.out.Errors = []jsonapi.Error{jaErr}
// 			} else {
// 				req.out.Data = nil

// 				// Aggregate operations
// 				// Commit to log
// 				// If success:
// 				//   * Commit sources
// 				//   * Return success
// 				// If failure:
// 				//   * Rollback sources
// 				//   * Try to transfer request to master node
// 				//   * Return response
// 			}
// 		}
// 	}
// }

// // versions ...
// type versions struct {
// 	versions map[string]uint64

// 	first   uint64
// 	commits []bool

// 	sync.Mutex
// }

// func (v *versions) reportCompletedTx(ver uint64) {
// 	v.Lock()
// 	defer v.Unlock()

// 	i := ver - v.first
// 	if uint64(len(v.commits)) <= i {
// 		v.commits = append(
// 			v.commits,
// 			make([]bool, i+1-uint64(len(v.commits)))...,
// 		)
// 	}
// 	v.commits[i] = true

// 	// Cleanup
// 	shift := 0
// 	for shift < len(v.commits) {
// 		if v.commits[shift] {
// 			shift++
// 		} else {
// 			break
// 		}
// 	}
// 	for j := range v.commits {
// 		if j+shift < len(v.commits) {
// 			v.commits[j] = v.commits[j+shift]
// 		} else {
// 			v.commits[j] = false
// 		}
// 	}

// }
