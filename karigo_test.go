package karigo_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/cmd/karigo/util"
	"github.com/mfcochauxlaberge/karigo/internal/freeport"
	"github.com/mfcochauxlaberge/karigo/internal/gold"

	"github.com/stretchr/testify/assert"
)

var runner *gold.Runner

func TestMain(m *testing.M) {
	// Runner
	runner = gold.NewRunner("testdata")

	runner.Filters = append(runner.Filters,
		gold.FilterBcryptHashes,
		gold.FilterUUIDs,
	)

	err := runner.Prepare()
	if err != nil {
		panic(err)
	}

	// Execution
	os.Exit(m.Run())
}

func TestKarigo(t *testing.T) {
	assert := assert.New(t)

	host := startServer()

	tests := []struct {
		name     string
		requests []request
	}{
		{
			name: "meta",
			requests: []request{
				{
					method: "POST",
					path:   "/0_meta",
					payload: `
						{
							"data": {
								"attributes": {
									"value": "some value"
								},
								"id": "some-key",
								"type": "0_meta"
							}
						}
					`,
				}, {
					method: "GET",
					path:   "/0_meta/some-key",
				}, {
					method: "PATCH",
					path:   "/0_meta/some-key",
					payload: `
						{
							"data": {
								"attributes": {
									"value": "value changed"
								},
								"id": "some-key",
								"type": "0_meta"
							}
						}
					`,
				}, {
					method: "GET",
					path:   "/0_meta/some-key",
				}, {
					method: "DELETE",
					path:   "/0_meta/some-key",
				}, {
					method: "GET",
					path:   "/0_meta/some-key",
				},
			},
		}, {
			name: "basic security with password in meta",
			requests: []request{
				{
					method: "POST",
					path:   "/0_meta",
					payload: `
						{
							"data": {
								"attributes": {
									"value": "some value"
								},
								"id": "some-key",
								"type": "0_meta"
							}
						}
					`,
				}, {
					method: "POST",
					path:   "/0_meta",
					payload: `
						{
							"data": {
								"attributes": {
									"value": "p@ssw0rd"
								},
								"id": "password",
								"type": "0_meta"
							}
						}
					`,
				}, {
					method: "POST",
					path:   "/0_meta",
					payload: `
						{
							"data": {
								"attributes": {
									"value": "no password, rejected"
								},
								"id": "another-key",
								"type": "0_meta"
							}
						}
					`,
				}, {
					method: "POST",
					path:   "/0_meta",
					payload: `
						{
							"data": {
								"attributes": {
									"value": "another value"
								},
								"id": "another-key",
								"type": "0_meta"
							},
							"meta": {
								"password": "p@ssw0rd"
							}
						}
					`,
				}, {
					method: "PATCH",
					path:   "/0_meta/another-key",
					payload: `
						{
							"data": {
								"attributes": {
									"value": "no password, rejected"
								},
								"id": "another-key",
								"type": "0_meta"
							}
						}
					`,
				}, {
					method: "PATCH",
					path:   "/0_meta/another-key",
					payload: `
						{
							"data": {
								"attributes": {
									"value": "new value"
								},
								"id": "another-key",
								"type": "0_meta"
							},
							"meta": {
								"password": "p@ssw0rd"
							}
						}
					`,
				}, {
					method: "GET",
					path:   "/0_meta/another-key",
				}, {
					method: "DELETE",
					path:   "/0_meta/password",
					payload: `
						{
							"meta": {
								"password": "wrongpassword"
							}
						}
					`,
				}, {
					method: "DELETE",
					path:   "/0_meta/password",
					payload: `
						{
							"meta": {
								"password": "p@ssw0rd"
							}
						}
					`,
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

func startServer() string {
	server := util.CreateServer(karigo.Config{
		Hosts: []string{"127.0.0.1"},
		Port:  freeport.Find(),
	})

	server.DisableLogger()

	// Run server
	go func() {
		server.Run()
	}()
	time.Sleep(100 * time.Millisecond)

	return "127.0.0.1:" + strconv.Itoa(int(server.Port))
}

func do(method, host, path string, body []byte) (int, []byte, []byte, error) {
	// Build request
	req, err := http.NewRequest(
		method,
		"http://"+host+path,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return 0, nil, nil, err
	}

	// Send request
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer res.Body.Close()

	// Read response
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, nil, err
	}

	// Format response body
	if len(resBody) > 0 {
		resBody = gold.FilterFormatJSON(resBody)
	}

	return res.StatusCode, buildSortedHeader(res.Header), resBody, err
}

func buildSortedHeader(h http.Header) []byte {
	header := []string{}

	for key, vals := range h {
		header = append(header, "")
		j := len(header) - 1

		header[j] = key + ": "

		if key == "Date" {
			header[j] += "Abc, 01 Def 2345 67:89:01 GMT"
		} else {
			for i := 0; i < len(vals); i++ {
				header[j] += vals[i]
			}
		}

		header[j] += "\n"
	}

	sort.Strings(header)

	headerBytes := []byte{}
	for _, v := range header {
		headerBytes = append(headerBytes, []byte(v)...)
	}

	return headerBytes
}

type request struct {
	method  string
	path    string
	payload string
}
