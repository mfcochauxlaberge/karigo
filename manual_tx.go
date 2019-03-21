package karigo

// // NewManualTx ...
// func NewManualTx(src DirectSource) *ManualTx {
// 	return &ManualTx{
// 		src: src,
// 	}
// }

// // ManualTx ...
// type ManualTx struct {
// 	src  Source
// 	undo []Op

// 	sync.Mutex
// }

// // Apply ...
// func (m *ManualTx) Apply(ops []Op) error {
// 	m.Lock()
// 	defer m.Unlock()

// 	m.Apply(ops)

// 	return nil
// }

// // Rollback ...
// func (m *ManualTx) Rollback() error {
// 	m.Lock()
// 	defer m.Unlock()

// 	m.Apply(m.undo)

// 	return nil
// }

// // Commit ...
// func (m *ManualTx) Commit() error {
// 	m.Lock()
// 	defer m.Unlock()

// 	m.undo = nil

// 	return nil
// }
