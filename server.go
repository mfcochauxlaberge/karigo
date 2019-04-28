package karigo

import (
	"net/http"

	"github.com/mfcochauxlaberge/jsonapi"
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

	doc := s.Nodes[domain].Handle(r)

	pl, err := jsonapi.Marshal(doc, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
	}

	w.Write(pl)
}
