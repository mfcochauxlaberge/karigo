package karigo

import (
	"net/http"
)

// Server ...
type Server struct {
	Nodes map[string]*Node
}

// Run ...
func (s *Server) Run() {
	err := http.ListenAndServe(":8080", s)
	if err != http.ErrServerClosed {
		panic(err)
	}
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Host

	if _, ok := s.Nodes[domain]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// node := s.Nodes[domain]

	// rawreq := &RawRequest{}
	// rawreq.URL = r.URL.RawPath

	_ = s.Nodes[domain].Handle(r)
}
