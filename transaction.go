package karigo

// Tx ...
type Tx func(*Checkpoint)

// TxDefault ...
func TxDefault(cp *Checkpoint, ops []Op) {
	cp.Apply(ops)
}

// // TxNotImplemented ...
// func TxNotImplemented(cp *Checkpoint) {
// 	cp.Fail(ErrNotImplemented)
// }

// // TxNotFound ...
// func TxNotFound(cp *Checkpoint) {
// 	cp.Fail(ErrNotFound)
// }
