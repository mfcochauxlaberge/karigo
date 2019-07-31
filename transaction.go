package karigo

// Tx ...
type Tx func(*Checkpoint)

// TxNothing ...
func TxNothing(*Checkpoint, []Op) {}

// TxGet ...
func TxGet(cp *Checkpoint, ops []Op) {}

// TxCreate ...
func TxCreate(cp *Checkpoint, ops []Op) {}

// TxUpdate ...
func TxUpdate(cp *Checkpoint, ops []Op) {}

// TxDelete ...
func TxDelete(cp *Checkpoint, ops []Op) {}

// TxNotImplemented ...
func TxNotImplemented(cp *Checkpoint) {
	cp.Fail(ErrNotImplemented)
}

// TxNotFound ...
func TxNotFound(cp *Checkpoint) {
	cp.Fail(ErrNotFound)
}

// tx ...
// type tx struct{}
