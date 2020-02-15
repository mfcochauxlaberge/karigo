package karigo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompile(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name          string
		code          []byte
		expectedError error
	}{
		{
			name:          "nil code",
			code:          nil,
			expectedError: errors.New("karigo: can't compile empty code"),
		},
		{
			name:          "empty slice code",
			code:          []byte{},
			expectedError: errors.New("karigo: can't compile empty code"),
		},
	}

	for _, test := range tests {
		_, err := compile(test.code)
		assert.Equal(test.expectedError, err, test.name)
	}
}
