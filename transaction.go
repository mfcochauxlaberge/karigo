package karigo

// Tx ...
type Tx func(*Checkpoint)

// TxGet ...
func TxGet(cp *Checkpoint) {}

// TxCreate ...
func TxCreate(cp *Checkpoint) {}

// TxUpdate ...
func TxUpdate(cp *Checkpoint) {}

// TxDelete ...
func TxDelete(cp *Checkpoint) {}

// TxNotImplemented ...
// func TxNotImplemented(cp *Checkpoint) {
// 	cp.Fail(ErrNotImplemented)
// }

// TxNotFound ...
// func TxNotFound(cp *Checkpoint) {
// 	cp.Fail(ErrNotFound)
// }

// tx ...
type tx struct{}
