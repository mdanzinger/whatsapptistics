package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	*mux.Router
}

// NewServer creates a server instance and starts it
func NewServer() *Server {
	// Create router
	r := mux.NewRouter()

	//// Add version prefix
	//vr := r.PathPrefix(version).Subrouter()

	//http.ListenAndServe(":8080", r)

	return &Server{
		r,
	}

}

// Start starts the server with the supplied port
func (s *Server) Start(port string) {
	http.ListenAndServe(port, s.Router)
}
