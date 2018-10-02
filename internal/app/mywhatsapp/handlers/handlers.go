package handlers

import (
	"github.com/mdanzinger/whatsapp/internal/app/mywhatsapp/server"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(server *server.Server) {
	registerReportRoutes(server.Router)
}

func registerReportRoutes(r *mux.Router) {
	r.HandleFunc("/report", newReport).Methods(http.MethodPost)
}
