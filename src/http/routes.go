package http

import "net/http"

// RegisterRoutes registers all the routes for the service
func (s *Server) registerRoutes() {
	// Home page
	s.r.HandleFunc("/", s.handlers.serveIndex).Methods(http.MethodGet)

	// New Report
	s.r.HandleFunc("/report", s.handlers.newReport).Methods(http.MethodPost)

	// Get Report
	s.r.HandleFunc("/report/{id}", s.handlers.serveReport).Methods(http.MethodGet)

}
