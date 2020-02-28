package karigo_test

import (
	"testing"

	. "github.com/mfcochauxlaberge/karigo"

	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/stretchr/testify/assert"
)

func TestParseOp(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		str string
		typ int
		op  Op
		err error
	}{
		{
			str: `set.id.field = "value"`,
			typ: jsonapi.AttrTypeString,
			op:  Op{},
			err: nil,
		},
	}

	for _, test := range tests {
		op, err := ParseOp(test.str, test.typ)
		assert.Equal(test.err, err)
		assert.Equal(test.op, op)
	}
}
