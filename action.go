package karigo

import "github.com/mfcochauxlaberge/karigo/query"

// Action ...
type Action func(*Checkpoint)

// ActionDefault ...
func ActionDefault(cp *Checkpoint) {}

// ActionPostSet ...
func ActionPostSet(cp *Checkpoint) {
	id := cp.Res.Get("name").(string)

	// Overwrite ID before the operation gets applied.
	cp.Apply([]query.Op{query.NewOpSet("0_sets", "_", "id", id)})
}

// ActionPostAttr ...
func ActionPostAttr(cp *Checkpoint) {
	set := cp.Res.GetToOne("set")
	name := cp.Res.Get("name").(string)
	id := set + "_" + name

	// Overwrite ID before the operation gets applied.
	cp.Apply([]query.Op{query.NewOpSet("0_attrs", "_", "id", id)})
}

// ActionPostRel ...
func ActionPostRel(cp *Checkpoint) {
	set := cp.Res.GetToOne("set")
	name := cp.Res.Get("name").(string)
	id := set + "_" + name

	// Overwrite ID before the operation gets applied.
	cp.Apply([]query.Op{query.NewOpSet("0_attrs", "_", "id", id)})
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
