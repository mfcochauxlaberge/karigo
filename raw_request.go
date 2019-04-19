package karigo

import (
	"io/ioutil"
	"net/http"
)

// NewRawRequest ...
func NewRawRequest(r *http.Request) *RawRequest {
	rr := &RawRequest{}
	return rr
}

// RawRequest ...
type RawRequest struct {
	URL    string
	Method string
	Token  []byte
	Body   []byte
}

func encodeRawRequest(r *http.Request) []byte {
	rr := []byte{}
	body, _ := ioutil.ReadAll(r.Body)
	rr = body
	return rr
}
