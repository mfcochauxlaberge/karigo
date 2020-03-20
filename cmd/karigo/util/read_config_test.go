package util_test

import (
	"testing"

	. "github.com/mfcochauxlaberge/karigo"
	. "github.com/mfcochauxlaberge/karigo/cmd/karigo/util"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig("testdata/configs/example.yml")
	assert.NoError(t, err)
	assert.Equal(
		t,
		Config{
			Port: 6820,
			Hosts: []string{
				"127.0.0.1", "localhost", "example.com",
			},
			Journal: map[string]string{
				"driver":   "psql",
				"addr":     "127.0.0.1",
				"database": "exampledb",
				"user":     "username",
				"password": "secret",
			},
			Sources: map[string]map[string]string{
				"main": {
					"driver": "memory",
				},
			},
		},
		config,
	)
}
