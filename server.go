package karigo

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
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
	// s.logger.Formatter = &logrus.TextFormatter{}
	s.logger.Formatter = &logrus.JSONFormatter{}

	s.logger.WithField("event", "server_started").Info("Server started")

	// Nodes
	s.Nodes = map[string]*Node{}

	local := &Node{}
	s.Nodes["localhost"] = local

	// Listen and serve
	err := http.ListenAndServe(":8080", s)
	if err != http.ErrServerClosed {
		panic(err)
	}
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.NewV4()
	shortID := makeShortID(requestID)

	// URL
	domain, _ := domainAndPort(r.Host)
	if domain == "" || domain == "127.0.0.1" {
		domain = "localhost"
	}

	logger := s.logger.WithFields(logrus.Fields{
		"id":     shortID,
		"domain": domain,
	})

	logger.WithField("event", "incoming_request").Info("New request incoming")

	if _, ok := s.Nodes[domain]; !ok {
		logger.WithField("event", "unknown_domain").Warn("Domain not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	url, err := jsonapi.ParseRawURL(&jsonapi.Schema{}, r.URL.String())
	if err != nil {
		s.logger.WithField("url", r.URL.String()).Warn("Invalid URL")
	}

	logger = s.logger.WithFields(logrus.Fields{
		"url": url,
	})

	req := &Request{
		Method: r.Method,
		URL:    url,
	}

	doc := s.Nodes[domain].Handle(req)

	pl, err := jsonapi.Marshal(doc, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
	}

	w.Write(pl)
}

func domainAndPort(host string) (string, int) {
	fragments := strings.Split(host, ":")

	domain := fragments[0]

	var (
		port int
		err  error
	)
	if len(fragments) > 1 {
		port, err = strconv.Atoi(fragments[1])
		if err != nil {
			port = 0
		}
	}

	return domain, port
}

func makeShortID(uuid uuid.UUID) string {
	return uuid.String()[0:8]
}
