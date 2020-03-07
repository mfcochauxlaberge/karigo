package karigo_test

import (
	"testing"

	. "github.com/mfcochauxlaberge/karigo"

	"github.com/stretchr/testify/assert"
)

func TestOpEncode(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		ops  []Op
	}{
		{
			name: "nil ops",
			ops:  nil,
		}, {
			name: "no ops",
			ops:  []Op{},
		}, {
			name: "1 op",
			ops: []Op{
				{
					Key: Key{
						Set:   "set",
						ID:    "id",
						Field: "field",
					},
					Op:    OpSet,
					Value: "string value",
				},
			},
		},
	}

	for _, test := range tests {
		data, err := Encode(0, test.ops)
		assert.NoError(err)

		ops, err := Decode(0, data)
		assert.NoError(err)

		assert.Equal(test.ops, ops)
	}
}
