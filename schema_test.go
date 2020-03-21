package karigo_test

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/mfcochauxlaberge/karigo/internal/gold"
	"github.com/stretchr/testify/assert"
)

func TestSchemaChanges(t *testing.T) {
	assert := assert.New(t)

	host := startServer()

	tests := []struct {
		name     string
		requests []request
	}{
		{
			name: "add set",
			requests: []request{
				{
					method: "POST",
					path:   "/0_sets",
					payload: `
						{
							"data": {
								"attributes": {
									"name": "new_set"
								},
								"type": "0_sets"
							}
						}
					`,
				}, {
					method: "GET",
					path:   "/0_sets/new_set",
				},
			},
		},
	}

	for _, test := range tests {
		out := []byte{}

		for i, req := range test.requests {
			status, headers, body, err := do(req.method, host, req.path, []byte(req.payload))
			assert.NoError(err)

			out = append(out, req.method...)
			out = append(out, ' ')
			out = append(out, req.path...)
			out = append(out, '\n')
			out = append(out, '\n')

			out = append(out, []byte(strconv.Itoa(status))...)
			out = append(out, ' ')
			out = append(out, []byte(http.StatusText(status))...)
			out = append(out, '\n')

			out = append(out, headers...)
			out = append(out, '\n')

			if len(body) > 0 {
				out = append(out, body...)
				out = append(out, '\n')
				out = append(out, '\n')
			}

			if i < len(test.requests)-1 {
				out = append(out, []byte("##################################################")...)
				out = append(out, '\n')
				out = append(out, '\n')
			}
		}

		// Golden file
		filename := strings.Replace(test.name, " ", "_", -1)
		path := filepath.Join("goldenfiles", "replays", filename)

		err := runner.Test(path, out)
		if _, ok := err.(gold.ComparisonError); ok {
			assert.Fail("file is different", test.name)
		} else if err != nil {
			panic(err)
		}
	}
}
