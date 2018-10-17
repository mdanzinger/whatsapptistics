package http

import "net/http"

// RegisterRoutes registers all the routes for the service
func (s *Server) RegisterRoutes() {
	// Home page
	s.r.HandleFunc("/", s.Handlers.serveIndex).Methods(http.MethodGet)

	// New Report
	s.r.HandleFunc("/report", s.Handlers.newReport).Methods(http.MethodPost)

	// Get Report
	s.r.HandleFunc("/report/:id", s.Handlers.serveReport).Methods(http.MethodGet)

}
