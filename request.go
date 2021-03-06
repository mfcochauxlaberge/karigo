package karigo

import (
	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/rs/zerolog"
)

// Methods
const (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

// Request ...
type Request struct {
	ID string

	Method string
	URL    *jsonapi.URL

	Body []byte
	Doc  *jsonapi.Document

	Logger zerolog.Logger
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
