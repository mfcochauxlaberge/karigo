package karigo

// Tx ...
type Tx func(*Checkpoint)

// TxDefault ...
func TxDefault(cp *Checkpoint) {}
