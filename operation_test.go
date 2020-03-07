package karigo_test

import (
	"testing"

	. "github.com/mfcochauxlaberge/karigo"

	"github.com/stretchr/testify/assert"
)

func TestOpBinary(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		op          Op
		expectedBin []byte
		expectedErr error
	}{
		{
			op: Op{
				Key: Key{
					Set:   "set",
					ID:    "id",
					Field: "field",
				},
				Op:    OpSet,
				Value: "string",
			},
			expectedBin: []byte{
				0x73, 0x65, 0x74, 0x2e, 0x69, // set.
				0x64, 0x2e, 0x66, 0x69, // id.
				0x65, 0x6c, 0x64, // field
				0x20, 0x3d, 0x20, // =
				0x1, 0x0, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, // string
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		// Marshal into binary
		data, err := test.op.MarshalBinary()
		assert.Equal(test.expectedErr, err)
		assert.Equal(test.expectedBin, data)

		// Unmarshal from binary
		op := &Op{}
		err = op.UnmarshalBinary(data)
		assert.Equal(test.expectedErr, err)
		assert.Equal(test.op, *op)
	}
}
