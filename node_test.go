package karigo_test

import (
	"fmt"
	"math/rand"
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

	_ = src.Apply(NewOpAddSet("things"))
	_ = src.Apply(NewOpActivateSet("things"))

	// Journal
	journal := &memory.Journal{}

	// Node
	node := NewNode(journal, src)

	go func() { _ = node.Run() }()

	url, err := jsonapi.NewURLFromRaw(schema, "/things")
	if err != nil {
		panic(err)
	}

	req := &Request{
		Method: "GET",
		URL:    url,
	}
	rnd := rand.Uint32()
	fmt.Printf("start %d\n", rnd)
	res := node.Handle(req)
	fmt.Printf("done %d\n", rnd)

	fmt.Printf("res: %v\n", res)
	if len(res.Errors) > 0 {
		t.Errorf("At least one error occurred: %v\n", err)
	}
}
