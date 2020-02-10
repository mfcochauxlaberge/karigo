package memory_test

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo/drivertest"
	. "github.com/mfcochauxlaberge/karigo/memory"
)

func TestMemorySource(t *testing.T) {
	src := &Source{}
	jrnl := &Journal{}

	drivertest.Test(t, src, jrnl)
}
