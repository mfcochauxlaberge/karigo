package karigo

// Tx ...
type Tx func(*Checkpoint)

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
