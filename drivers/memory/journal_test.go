package memory_test

import (
	"testing"

	"github.com/mfcochauxlaberge/karigo"
	. "github.com/mfcochauxlaberge/karigo/drivers/memory"
	"github.com/mfcochauxlaberge/karigo/drivertest"
)

var _ karigo.Journal = (*Journal)(nil)

func TestJournal(t *testing.T) {
	jrnl := &Journal{}

	drivertest.TestJournal(t, jrnl)
}
