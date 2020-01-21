package karigo_test

import (
	"testing"

	"github.com/mfcochauxlaberge/jsonapi"

	. "github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/memory"
)

func TestNode(t *testing.T) {
	// Fake type
	typ := jsonapi.Type{}
	typ.Name = "things"
	_ = typ.AddAttr(jsonapi.Attr{
		Name:     "name",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})

	// Schema
	schema := &jsonapi.Schema{}
	_ = schema.AddType(typ)

	// Source
	src := &memory.Source{}
	_ = src.Reset() // TODO Necessary?

	tx, _ := src.NewTx()

	_ = tx.Apply(NewOpAddSet("things"))
	_ = tx.Apply(NewOpActivateSet("things"))

	// Journal
	journal := &memory.Journal{}

	// Node
	node := NewNode(journal, src)

	url, err := jsonapi.NewURLFromRaw(schema, "/things")
	if err != nil {
		panic(err)
	}

	req := &Request{
		Method: "GET",
		URL:    url,
	}
	res := node.Handle(req)

	if len(res.Errors) > 0 {
		t.Errorf("At least one error occurred: %v\n", err)
	}
}
