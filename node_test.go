package karigo_test

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	. "github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/memory"

	"github.com/mfcochauxlaberge/jsonapi"
)

func TestNode(t *testing.T) {
	// Fake type
	typ := jsonapi.Type{}
	typ.Name = "things"
	typ.AddAttr(jsonapi.Attr{
		Name:     "name",
		Type:     jsonapi.AttrTypeString,
		Nullable: false,
	})

	// Schema
	schema := &jsonapi.Schema{}
	schema.AddType(typ)

	// Source
	src := &memory.Source{}
	src.Reset() // TODO Necessary?

	src.Apply(karigo.NewOpAddSet("things"))

	// Journal
	journal := &memory.Journal{}

	// Node
	node := NewNode(journal, src)
	go node.Run()

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
		t.Errorf("At least one error occured: %v\n", err)
	}
}
