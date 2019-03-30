package karigo

import (
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
func NewRequest(rr *RawRequest) (*Request, error) {
	jurl, err := jsonapi.ParseRawURL(&jsonapi.Registry{}, rr.URL)
	// if err != nil {
	// 	return nil, err
	// }

	req := &Request{
		Method: rr.Method,
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
