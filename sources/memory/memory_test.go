package memory

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo/sourcetest"
)

func TestMemory(t *testing.T) {
	src := &Memory{}
	err := sourcetest.Test(src)
	if err != nil {
		t.Errorf("Source %T failed: %s", src, err)
	}
}
