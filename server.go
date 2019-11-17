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

	defer func() {
		// This is for separating the requests in
		// the logger's output. It is useful for
		// debugging in the early stages of this
		// project, but it should be removed
		// eventually.
		fmt.Println()
	}()

	logger.Info().
		Str("event", "read_request").
		Str("domain", domain).
		Int("port", port).
		Str("method", r.Method).
		Str("url", r.URL.String()).
		Msg("Incoming request")

	// Find node from domain
	var (
		node *Node
		ok   bool
	)

	if node, ok = s.Nodes[domain]; !ok {
		logger.Info().
			Str("event", "unknown_domain").
			Msg("App not found from domain")

		_ = sendResponse(w, http.StatusNotFound, nil, logger)

		return
	}

	// Parse URL
	url, err := jsonapi.NewURLFromRaw(node.schema, r.URL.String())
	if err != nil {
		logger.
			Err(err).
			Str("url", r.URL.String()).
			Msg("Invalid URL")

		_ = sendResponse(w, http.StatusInternalServerError, nil, logger)

		return
	}

	// Set default page size
	if url.Params.PageSize == 0 {
		url.Params.PageSize = 10
	}

	logger.Debug().
		Str("event", "parse_url").
		Str("url", url.UnescapedString()).
		Msg("URL parsed")

	// Build request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Err(err).Send()

		_ = sendResponse(w, http.StatusInternalServerError, nil, logger)

		return
	}

	req := &Request{
		ID:     requestID,
		Method: r.Method,
		URL:    url,
		Body:   body,
		Logger: logger,
	}

	// Send request to node
	doc := node.Handle(req)

	if doc == nil {
		_ = sendResponse(w, http.StatusInternalServerError, nil, logger)

		return
	}

	// Marshal response
	pl, err := jsonapi.MarshalDocument(doc, url)
	if err != nil {
		logger.Err(err).Send()

		_ = sendResponse(
			w,
			http.StatusInternalServerError,
			[]byte("500 Internal Server Error"),
			logger,
		)

		return
	}

	if r.Method == DELETE && len(doc.Errors) == 0 {
		_ = sendResponse(w, http.StatusNoContent, nil, logger)

		return
	}

	// Indent response
	out := &bytes.Buffer{}
	_ = json.Indent(out, pl, "", "\t")

	// Send response
	_ = sendResponse(w, http.StatusOK, out.Bytes(), logger)
}

func sendResponse(w http.ResponseWriter, code int, body []byte, logger zerolog.Logger) error {
	var err error

	w.WriteHeader(code)

	if len(body) != 0 {
		_, err = w.Write(body)
	}

	logger.Info().
		Str("event", "send_response").
		Int("status_code", code).
		Str("status_code_text", http.StatusText(code)).
		Msg("Send response")

	return err
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
