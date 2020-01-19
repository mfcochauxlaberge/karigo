package karigo

// Action ...
type Action func(*Checkpoint)

// ActionDefault ...
func ActionDefault(cp *Checkpoint) {}
