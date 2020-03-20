package freeport_test

import (
	"testing"

	. "github.com/mfcochauxlaberge/karigo/internal/freeport"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	port := Find()
	assert.NotEqual(t, 0, port)
}
