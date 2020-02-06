package memory

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo/drivertest"
)

func TestMemorySource(t *testing.T) {
	src := &Source{}
	jrnl := &Journal{}

	err := drivertest.Test(t, src, jrnl)
	if err != nil {
		t.Errorf("Source %T failed: %s", src, err)
	}
}
