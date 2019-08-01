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
	jurl, err := jsonapi.NewURLFromRaw(&jsonapi.Schema{}, r.URL.RawPath)
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
	ID     string
	Method string
	URL    *jsonapi.URL
	Body   *jsonapi.Document
}

func (r *Request) isSchemaChange() bool {
	if r.Method != GET {
		switch r.URL.ResType {
		case "0_sets", "0_attrs", "0_rels":
			return true
		}
	}
	return false
}
