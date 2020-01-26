package memory

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo/sourcetest"
)

func TestMemorySource(t *testing.T) {
	src := &Source{}
	jrnl := &Journal{}

	err := sourcetest.Test(t, src, jrnl)
	if err != nil {
		t.Errorf("Source %T failed: %s", src, err)
	}
}
