package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdanzinger/whatsapptistics/chat"
	"github.com/mdanzinger/whatsapptistics/report"
)

// DefaultPort is the default port to listen to
const DefaultPort = ":8000"

// Server represents an HTTP server
type Server struct {
	router *mux.Router

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
		router: router,
		Port:   DefaultPort,
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
	s.logger.Fatal(http.ListenAndServe(s.Port, s.router))
}
