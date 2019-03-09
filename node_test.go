package karigo_test

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/sources"
)

func TestNode(t *testing.T) {
	mem := &sources.Memory{
		ID:       "memory",
		Location: "local",
	}

	node := karigo.NewNode(mem)

	go node.Run()

	req := &karigo.Request{}
	res := node.Handle(req)
	if len(res.Errors) == 0 {
		t.Errorf("No errors occured.\n")
	}

	err := node.Close()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
