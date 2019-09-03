package karigo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	s.logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		QuoteEmptyFields: true,
	}
	// s.logger.Formatter = &logrus.JSONFormatter{}
	s.logger.SetLevel(logrus.DebugLevel)

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

	for _, node := range s.Nodes {
		node.logger = s.logger
	}

	// Listen and serve
	err := http.ListenAndServe(":8080", s)
	if err != http.ErrServerClosed {
		panic(err)
	}
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.NewV4().String()[:8]

	// Parse domain and port
	domain, port := domainAndPort(r.Host)

	// Populate logger with rid
	logger := s.logger.WithField("rid", requestID)

	logger.WithFields(logrus.Fields{
		"event":  "read_request",
		"domain": domain,
		"port":   port,
	}).Info("New request incoming")

	// Find node from domain
	var (
		node *Node
		ok   bool
	)
	if node, ok = s.Nodes[domain]; !ok {
		logger.WithField("event", "unknown_domain").Warn("App not found from domain")
		w.WriteHeader(http.StatusNotFound)
		logger.WithField("event", "set_http_status_code").Warn("See HTTP status code")
		logger.WithField("event", "send_response").Warn("Send response")
		return
	}

	logger.WithFields(logrus.Fields{
		"app": node.Name,
	}).Debug("App found")

	// Parse URL
	url, err := jsonapi.NewURLFromRaw(node.schema, r.URL.String())
	if err != nil {
		s.logger.WithError(err).WithField("url", r.URL.String()).Warn("Invalid URL")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set default page size
	if url.Params.PageSize == 0 {
		url.Params.PageSize = 10
	}

	logger.WithFields(logrus.Fields{
		"event": "parse_url",
		"url":   url.String(),
	}).Debug("URL parsed")

	// Build request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req := &Request{
		ID:     requestID,
		Method: r.Method,
		URL:    url,
		Body:   body,
	}

	// Send request to node
	doc := node.Handle(req)

	// Marshal response
	pl, err := jsonapi.Marshal(doc, url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 Internal Server Error"))
	}
	if r.Method == DELETE && len(doc.Errors) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Indent response
	out := &bytes.Buffer{}
	_ = json.Indent(out, pl, "", "\t")

	// Send response
	_, _ = w.Write(out.Bytes())
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
