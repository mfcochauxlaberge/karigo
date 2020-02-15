package memory_test

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	. "github.com/mfcochauxlaberge/karigo/drivers/memory"
	"github.com/mfcochauxlaberge/karigo/drivertest"
)

var _ karigo.Source = (*Source)(nil)

func TestMemorySource(t *testing.T) {
	src := &Source{}
	jrnl := &Journal{}

	drivertest.Test(t, src, jrnl)
}
