package karigo

// Action ...
type Action func(*Checkpoint)

// ActionDefault ...
func ActionDefault(cp *Checkpoint) {}

// ActionPostSet ...
func ActionPostSet(cp *Checkpoint) {
	cp.SetID("$0", cp.Get("0_sets", "$0", "name").(string))
}

// ActionPostAttr ...
func ActionPostAttr(cp *Checkpoint) {
	set := cp.Get("0_sets", "$0", "set").(string)
	name := cp.Get("0_sets", "$0", "name").(string)

	cp.SetID("$0", set+"_"+name)
}

// ActionPostRel ...
func ActionPostRel(cp *Checkpoint) {
	set := cp.Get("0_sets", "$0", "set").(string)
	name := cp.Get("0_sets", "$0", "from-name").(string)

	cp.SetID("$0", set+"_"+name)
}

// ActionPatchSet ...
func ActionPatchSet(cp *Checkpoint) {}

// ActionPatchAttr ...
func ActionPatchAttr(cp *Checkpoint) {}

// ActionPatchRel ...
func ActionPatchRel(cp *Checkpoint) {}

// ActionDeleteSet ...
func ActionDeleteSet(cp *Checkpoint) {}

// ActionDeleteAttr ...
func ActionDeleteAttr(cp *Checkpoint) {}

// ActionDeleteRel ...
func ActionDeleteRel(cp *Checkpoint) {}
