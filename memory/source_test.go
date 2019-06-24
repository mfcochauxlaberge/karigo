package memory

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo/sourcetest"
)

func TestMemorySource(t *testing.T) {
	src := &Source{}
	err := sourcetest.Test(src)
	if err != nil {
		t.Errorf("Source %T failed: %s", src, err)
	}
	// t.Errorf("FAIL!")
}