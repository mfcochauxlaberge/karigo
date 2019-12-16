package karigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/mfcochauxlaberge/jsonapi"
	"github.com/mfcochauxlaberge/karigo/memory"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

func NewServer() *Server {
	s := &Server{
		Nodes: map[string]*Node{},
	}

	s.logger = s.logger.
		Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().Timestamp().Logger()

	return s
}

// Server ...
type Server struct {
	Nodes map[string]*Node

	logger zerolog.Logger
}

// Run ...
func (s *Server) Run(port uint) {
	s.logger.Info().Str("event", "server_start")

	if s.Nodes == nil {
		s.Nodes = map[string]*Node{}
	}

	// Add cluster control node
	ctlNode := &Node{
		Name: "0_ctl_node_1",
		main: source{src: &memory.Source{}},
		log:  &memory.Journal{},
	}
	s.Nodes[ctlNode.Name] = ctlNode

	// Add empty cluster schema to node
	sc := ClusterSchema()
	ops := schemaToOps(sc)
	err := ctlNode.apply(ops)
	if err != nil {
		panic(err)
	}

	for _, node := range s.Nodes {
		node.logger = s.logger
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	})

	handler := c.Handler(s)

	// Listen and serve
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
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
		Str("endpoint", r.URL.String()).
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
	status := http.StatusOK

	if len(doc.Errors) > 0 {
		fromStr, _ := strconv.ParseInt(doc.Errors[0].Status, 10, 64)
		status = int(fromStr)
	}

	_ = sendResponse(w, status, out.Bytes(), logger)
}

func (s *Server) DisableLogger() {
	s.logger = zerolog.Nop()
}

func (s *Server) SetOutput(w io.Writer) {
	s.logger = s.logger.Output(w)
}

func sendResponse(w http.ResponseWriter, code int, body []byte, logger zerolog.Logger) error {
	var err error

	w.Header().Add("Content-Type", "application/vnd.api+json")

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
