package karigo_test

import (
	"net/http/httptest"
	"testing"

	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/memory"
)

func TestNode(t *testing.T) {
	// Fake type
	typ := jsonapi.Type{}
	typ.Name = "things"
	typ.AddAttr(jsonapi.Attr{
		Name: "name",
		Type: jsonapi.AttrTypeString,
		Null: false,
	})

	// Schema
	schema := jsonapi.Schema{}
	schema.AddType(typ)

	// Source
	src := &memory.Source{
		ID:       "memory",
		Location: "local",
	}

	// Journal
	journal := &memory.Journal{}

	// Node
	node := karigo.NewNode(journal, src)
	go node.Run()

	//

	req := httptest.NewRequest("GET", "/things", nil)
	res := node.Handle(req)
	if len(res.Errors) == 0 {
		t.Errorf("No errors occured.\n")
	}

	err := node.Close()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
