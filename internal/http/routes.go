package http

import "net/http"

// RegisterRoutes registers all the routes for the service
func (s *Server) registerRoutes() {
	// Static assets
	fs := http.FileServer(http.Dir("../../internal/web/static/dist/"))
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Home page
	s.router.HandleFunc("/", s.handlers.serveIndex).Methods(http.MethodGet)

	// New Report
	s.router.HandleFunc("/report", s.handlers.newReport).Methods(http.MethodPost)

	// Get Report
	s.router.HandleFunc("/report/{id}", s.handlers.serveReport).Methods(http.MethodGet)

}
