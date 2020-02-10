package memory

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo/drivertest"
)

func TestMemorySource(t *testing.T) {
	src := &Source{}
	jrnl := &Journal{}

	drivertest.Test(t, src, jrnl)
}
