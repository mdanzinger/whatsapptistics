package http

import "net/http"

// RegisterRoutes registers all the routes for the service
func (s *Server) registerRoutes() {
	// Home page
	s.router.HandleFunc("/", s.handlers.serveIndex).Methods(http.MethodGet)

	// New Report
	s.router.HandleFunc("/report", s.handlers.newReport).Methods(http.MethodPost)

	// Get Report
	s.router.HandleFunc("/report/{id}", s.handlers.serveReport).Methods(http.MethodGet)

}
