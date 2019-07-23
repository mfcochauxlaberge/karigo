package karigo

import (
	"bytes"
	"encoding/json"
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
	s.logger.Formatter = &logrus.TextFormatter{}
	// s.logger.Formatter = &logrus.JSONFormatter{}

	s.logger.WithField("event", "server_start").Info("Server has started")

	if node, ok := s.Nodes["localhost"]; ok {
		ops := []Op{}
		ops = append(ops, NewOpSet("0_meta", "", "id", "password"))
		ops = append(ops, NewOpSet("0_meta", "password", "value", "123456seven"))
		ops = append(ops, NewOpSet("0_meta", "", "id", "name"))
		ops = append(ops, NewOpSet("0_meta", "name", "value", "test"))
		ops = append(ops, NewOpSet("0_meta", "", "id", "name-again"))
		ops = append(ops, NewOpSet("0_meta", "name-again", "value", "test"))

		ops = append(ops, NewOpAddSet("users")...)
		ops = append(ops, NewOpAddAttr("users", "username", "string", false)...)
		ops = append(ops, NewOpAddAttr("users", "name", "string", false)...)
		ops = append(ops, NewOpAddAttr("users", "password", "string", false)...)
		// ops = append(ops, NewOpAddAttr("users", "created-at", "string", time.Now())...)
		ops = append(ops, NewOpSet("users", "", "id", "abc123"))
		ops = append(ops, NewOpSet("users", "abc123", "username", "user1"))
		ops = append(ops, NewOpSet("users", "abc123", "name", "Bob"))
		ops = append(ops, NewOpSet("users", "abc123", "password", "j2K2sN1s7"))
		ops = append(ops, NewOpSet("users", "", "id", "def456"))
		ops = append(ops, NewOpSet("users", "def456", "username", "user2"))
		ops = append(ops, NewOpSet("users", "def456", "name", "John"))
		ops = append(ops, NewOpSet("users", "def456", "password", "K1nas82J2"))
		ops = append(ops, NewOpSet("users", "", "id", "ghi789"))
		ops = append(ops, NewOpSet("users", "ghi789", "username", "user3"))
		ops = append(ops, NewOpSet("users", "ghi789", "name", "Ted"))
		ops = append(ops, NewOpSet("users", "ghi789", "password", "aJ2n2s8sa"))

		err := node.apply(ops)
		if err != nil {
			panic(err)
		}
	}

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

	entry := s.logger.WithFields(logrus.Fields{
		"id":     shortID,
		"domain": domain,
	})

	entry.WithField("event", "incoming_request").Info("New request incoming")

	var (
		node *Node
		ok   bool
	)
	if node, ok = s.Nodes[domain]; !ok {
		entry.WithField("event", "unknown_domain").Warn("Domain not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	node.logger = s.logger

	url, err := jsonapi.NewURLFromRaw(node.schema, r.URL.String())
	if err != nil {
		s.logger.WithError(err).WithField("url", r.URL.String()).Warn("Invalid URL")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if url.Params.PageSize == 0 {
		url.Params.PageSize = 10
	}

	entry = s.logger.WithFields(logrus.Fields{
		"url": url,
	})

	req := &Request{
		Method: r.Method,
		URL:    url,
	}

	doc := node.Handle(req)

	pl, err := jsonapi.Marshal(doc, url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
	}

	out := &bytes.Buffer{}
	err = json.Indent(out, pl, "", "\t")

	w.Write(out.Bytes())
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
