package karigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/rs/zerolog"
)

// Server ...
type Server struct {
	Nodes map[string]*Node

	logger zerolog.Logger
}

// Run ...
func (s *Server) Run(port uint) {
	// Logger
	s.logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	s.logger = s.logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	s.logger.Info().Str("event", "server_start")

	for _, node := range s.Nodes {
		node.logger = s.logger
	}

	// Listen and serve
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), s)
	if err != http.ErrServerClosed {
		panic(err)
	}
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()[:8]

	// Parse domain and port
	domain, port := domainAndPort(r.Host)

	// Populate logger with rid
	logger := s.logger.With().Str("rid", requestID).Logger()

	logger.Info().
		Str("event", "read_request").
		Str("domain", domain).
		Int("port", port).
		Str("method", r.Method).
		Str("url", r.URL.String()).
		Msg("New request incoming")

	// Find node from domain
	var (
		node *Node
		ok   bool
	)

	if node, ok = s.Nodes[domain]; !ok {
		logger.Warn().
			Str("event", "unknown_domain").
			Msg("App not found from domain")

		w.WriteHeader(http.StatusNotFound)
		logger.Warn().
			Str("event", "send_response").
			Int("status_code", http.StatusNotFound).
			Msg("Send response")

		return
	}

	logger.Debug().
		Str("app", node.Name).
		Msg("App found")

	// Parse URL
	url, err := jsonapi.NewURLFromRaw(node.schema, r.URL.String())
	if err != nil {
		errlogger := logger.With().Err(err).Logger()
		errlogger.Warn().
			Str("url", r.URL.String()).
			Msg("Invalid URL")

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	// Set default page size
	if url.Params.PageSize == 0 {
		url.Params.PageSize = 10
	}

	logger.Debug().
		Str("event", "parse_url").
		Str("url", url.String()).
		Msg("URL parsed")

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
	pl, err := jsonapi.MarshalDocument(doc, url)
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
