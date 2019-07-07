package karigo

// Tx ...
type Tx func(*Checkpoint)

// TxGet ...
func TxGet(cp *Checkpoint) {}

// TxCreate ...
func TxCreate(cp *Checkpoint) {
	cp.Fail(ErrNotImplemented)
}

// TxUpdate ...
func TxUpdate(cp *Checkpoint) {
	cp.Fail(ErrNotImplemented)
}

// TxDelete ...
func TxDelete(cp *Checkpoint) {
	cp.Fail(ErrNotImplemented)
}

// TxNotImplemented ...
func TxNotImplemented(cp *Checkpoint) {
	cp.Fail(ErrNotImplemented)
}

// TxNotFound ...
func TxNotFound(cp *Checkpoint) {
	cp.Fail(ErrNotFound)
}

// tx ...
type tx struct{}
