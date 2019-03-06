package karigo_test

import (
	"fmt"
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/sources"
)

func TestNode(t *testing.T) {
	node := karigo.NewNode()

	mem := &sources.Memory{
		ID:       "memory",
		Location: "local",
	}
	err := node.AddSource("memory", mem)
	if err != nil {
		fmt.Printf("Run error: %s\n", err)
	}

	go node.Run()

	req := &karigo.Request{}
	res := node.Handle(req)
	if len(res.Errors) == 0 {
		t.Errorf("No errors occured.\n")
	}

	err = node.Close()
	// err = node.Close()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
