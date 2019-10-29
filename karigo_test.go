package karigo_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/memory"

	"github.com/stretchr/testify/assert"
)

func TestKarigo(t *testing.T) {
	assert := assert.New(t)

	// Server
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

	go func() {
		server.Run()
	}()
	time.Sleep(100 * time.Millisecond)

	// Test
	res, err := post("/0_meta", []byte(`
		{
			"data": {
				"attributes": {
					"value": "some value"
				},
				"id": "some-key",
				"type": "0_meta"
			}
		}
		`))
	expect := `
		{
			"data": {
				"attributes": {
				"value": "some value"
				},
				"id": "some-key",
				"links": {
				"self": "/0_meta/some-key"
				},
				"relationships": {},
				"type": "0_meta"
			},
			"jsonapi": {
				"version": "1.0"
			},
			"links": {
				"self": "/0_meta?fields%5B0_meta%5D=value&page%5Bsize%5D=10&sort=value%2Cid"
			}
		}`
	assert.NoError(err)
	assert.Equal(expect, res)
}

func post(url string, body []byte) ([]byte, error) {
	client := &http.Client{}
	res, err := client.Post(
		"http://localhost:8080"+url,
		"application/vnd.api+json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, err
}
