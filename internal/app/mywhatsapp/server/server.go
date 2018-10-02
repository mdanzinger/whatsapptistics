package server

import "github.com/gorilla/mux"

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	// Create router
	r := mux.NewRouter()

	//// Add version prefix
	//vr := r.PathPrefix(version).Subrouter()

	return &Server{
		r,
	}

}
