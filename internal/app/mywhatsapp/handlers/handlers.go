package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/report", newReport).Methods(http.MethodPost)
}
