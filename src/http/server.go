package http

import (
	"github.com/gorilla/mux"
	"github.com/mdanzinger/whatsapptistics/src/chat"
	"github.com/mdanzinger/whatsapptistics/src/report"
	"log"
	"net/http"
)

// DefaultPort is the default port to listen to
const DefaultPort = ":8000"

// Server represents an HTTP server
type Server struct {
	r *mux.Router

	// handler to use
	handlers *handler

	// port to listen to
	Port string

	// logger
	logger *log.Logger
}

// NewServer creates a new server with the services passed in
func NewServer(cs chat.ChatService, rs report.ReportService, l *log.Logger) *Server {
	router := mux.NewRouter()

	return &Server{
		r:    router,
		Port: DefaultPort,
		handlers: &handler{
			ChatService:   cs,
			ReportService: rs,
		},
		logger: l,
	}
}

// Start starts the server
func (s *Server) Start() {
	// Parse Templates
	s.handlers.CompileTemplates()

	// Register Routes
	s.registerRoutes()

	//Start Server
	s.logger.Fatal(http.ListenAndServe(s.Port, s.r))
}
