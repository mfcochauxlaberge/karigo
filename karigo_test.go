package karigo_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/memory"

	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update-golden-files", false, "update the golden files")

func TestMain(m *testing.M) {
	if *update {
		err := os.RemoveAll("testdata")
		if err != nil {
			panic(err)
		}

		err = os.Mkdir(filepath.Join("testdata"), os.ModePerm)
		if err != nil {
			panic(err)
		}

		err = os.Mkdir(filepath.Join("testdata", "goldenfiles"), os.ModePerm)
		if err != nil {
			panic(err)
		}

		err = os.Mkdir(filepath.Join("testdata", "goldenfiles", "replays"), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

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
		path := filepath.Join("testdata", "goldenfiles", "replays", filename)

		if !*update {
			// Retrieve the expected result from a file
			expected, err := ioutil.ReadFile(path)
			assert.NoError(err, test.name)
			assert.Equal(string(expected), string(out), test.name)
		} else {
			// Write the result to a file
			// TODO Figure out whether 0644 is okay or not.
			err := ioutil.WriteFile(path, out, 0644)
			assert.NoError(err)
		}
	}
}

func startServer() string {
	server := &karigo.Server{
		Nodes: map[string]*karigo.Node{},
	}

	src := &memory.Source{}
	_ = src.Reset()
	node := karigo.NewNode(&memory.Journal{}, src)
	node.Name = "test"
	node.Domains = []string{"localhost", "127.0.0.1"}

	for _, domain := range node.Domains {
		server.Nodes[domain] = node
	}

	port := findFreePort()

	// Run server
	go func() {
		server.Run(port)
	}()
	time.Sleep(100 * time.Millisecond)

	return "127.0.0.1:" + strconv.Itoa(int(port))
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
		dst := &bytes.Buffer{}
		err = json.Indent(dst, resBody, "", "\t")
		resBody = dst.Bytes()
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

func findFreePort() uint {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0
	}
	defer l.Close()

	return uint(l.Addr().(*net.TCPAddr).Port)
}

type request struct {
	method  string
	path    string
	payload string
}
