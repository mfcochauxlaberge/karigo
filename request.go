package karigo

import (
	"net/http"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Methods
const (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

// NewRequest ...
func NewRequest(r *http.Request) (*Request, error) {
	jurl, err := jsonapi.ParseRawURL(&jsonapi.Schema{}, r.URL.RawPath)
	// if err != nil {
	// 	return nil, err
	// }

	req := &Request{
		Method: r.Method,
		URL:    jurl,
		Body:   &jsonapi.Document{},
	}

	return req, err
}

// Request ...
type Request struct {
	Method string
	URL    *jsonapi.URL
	Body   *jsonapi.Document
}

// Response ...
type Response jsonapi.Document
