package karigo

// RawRequest ...
type RawRequest struct {
	URL    string
	Method string
	Token  []byte
	Body   []byte
}
