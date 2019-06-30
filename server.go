package karigo

import (
	"net/http"

	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	Nodes map[string]*Node

	logger *logrus.Logger
}

// Run ...
func (s *Server) Run() {
	// Logger
	s.logger = logrus.New()
	s.logger.Formatter = &logrus.TextFormatter{}

	// Nodes
	s.Nodes = map[string]*Node{}

	local := &Node{}
	s.Nodes["127.0.0.1"] = local

	// Listen and serve
	err := http.ListenAndServe(":8080", s)
	if err != http.ErrServerClosed {
		panic(err)
	}
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Host

	s.logger.Logf(logrus.InfoLevel, "Received request from %s", domain)

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
