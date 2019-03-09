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
func NewRequest(method, url string, body *jsonapi.Document) (*Request, error) {
	jurl, err := jsonapi.ParseRawURL(nil, url)
	if err != nil {
		return nil, err
	}

	req := &Request{
		Method: method,
		URL:    jurl,
		Body:   body,
	}

	return req, nil
}

// Request ...
type Request struct {
	Method string
	URL    *jsonapi.URL
	Body   *jsonapi.Document
}

// Response ...
type Response jsonapi.Document
