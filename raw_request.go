package karigo

// RawRequest ...
type RawRequest struct {
	Path   string
	Method string
	Token  []byte
	Body   []byte
}
