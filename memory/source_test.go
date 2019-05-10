package memory

import (
	"fmt"
	"testing"

	"github.com/mfcochauxlaberge/karigo/sourcetest"
)

func TestMemorySource(t *testing.T) {
	src := &Source{}
	err := sourcetest.Test(src)
	if err != nil {
		t.Errorf("Source %T failed: %s", src, err)
	} else {
		fmt.Printf("It didn't fail!\n")
	}
	t.Errorf("FAIL!")
}
