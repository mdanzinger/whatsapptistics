package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// DefaultAddr is the default port to listen to
const DefaultPort = ":8000"

// Server represents an HTTP server
type Server struct {
	r *mux.Router

	// Handler to use
	Handlers *Handler

	// port to listen to
	Port string
}

func NewServer() *Server {
	return &Server{
		Port: DefaultPort,
	}
}

// Start starts the server
func (s *Server) Start() {
	s.
		log.Fatal(http.ListenAndServe(s.Port, s.r))
}
