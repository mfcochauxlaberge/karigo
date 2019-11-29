package karigo_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/memory"

	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update-golden-files", false, "update the golden files")

func TestKarigo(t *testing.T) {
	assert := assert.New(t)

	host := startServer()

	tests := []struct {
		name    string
		method  string
		path    string
		payload string
	}{
		{
			name:   "example",
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
		},
	}

	for _, test := range tests {
		payload, headers, err := do(test.method, host, test.path, []byte(test.payload))
		fmt.Printf("err: %v\n", err)
		fmt.Printf("payload: %v\n", payload)

		// Golden file
		filename := strings.Replace(test.name, " ", "_", -1)
		path := filepath.Join("testdata", "goldenfiles", "replays", filename)

		if !*update {
			// Retrieve the expected result from a file
			expected, _ := ioutil.ReadFile(path)

			assert.NoError(err, test.name)
			assert.JSONEq(string(expected), string(payload))
		} else {
			// Write the result to a file
			dst := &bytes.Buffer{}
			err = json.Indent(dst, payload, "", "\t")
			assert.NoError(err)
			out := headers
			out = append(out, '\n')
			out = append(out, dst.Bytes()...)
			// TODO Figure out whether 0644 is okay or not.
			err = ioutil.WriteFile(path, out, 0644)
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

func do(method, host, path string, body []byte) ([]byte, []byte, error) {
	fmt.Printf("url: %s\n", host+path)

	req, err := http.NewRequest(
		method,
		"http://"+host+path,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("errrrrr: %v\n", err)
		return nil, nil, err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	header := []byte{}

	for key, vals := range res.Header {
		if len(header) > 0 {
			header = append(header, '\n')
		}

		header = append(header, []byte(key)...)
		header = append(header, []byte(": ")...)

		if key == "Date" {
			header = append(header, []byte("Abc, 01 Def 2345 67:89:01 GMT")...)
		} else {
			for i := 0; i < len(vals); i++ {
				header = append(header, []byte(vals[i])...)
			}
		}
	}

	header = append(header, header...)

	return resBody, header, err
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
