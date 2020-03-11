package karigo_test

import (
	"testing"

	. "github.com/mfcochauxlaberge/karigo"

	"github.com/stretchr/testify/assert"
)

func TestOpEncodeV0(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		ops  []Op
		err  string
	}{
		{
			name: "nil ops",
			ops:  nil,
			err:  "no ops",
		}, {
			name: "no ops",
			ops:  []Op{},
			err:  "no ops",
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
		}, {
			name: "2 ops",
			ops: []Op{
				{
					Key: Key{
						Set:   "set",
						ID:    "id",
						Field: "field",
					},
					Op:    OpSet,
					Value: "string value",
				}, {
					Key: Key{
						Set:   "set2",
						ID:    "id2",
						Field: "field2",
					},
					Op:    OpSet,
					Value: 42,
				},
			},
		},
	}

	for _, test := range tests {
		data, err := Encode(0, test.ops)

		if test.err == "" {
			assert.NoError(err, test.name)

			ops, err := Decode(0, data)
			assert.NoError(err, test.name)
			assert.Equal(test.ops, ops, test.name)
		} else {
			assert.EqualError(err, test.err, test.name)
		}
	}
}
