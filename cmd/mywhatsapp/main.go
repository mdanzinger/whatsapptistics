package main

import (
	"net/http"

	"github.com/mdanzinger/whatsapp/internal/app/mywhatsapp/handlers"
	"github.com/mdanzinger/whatsapp/internal/app/mywhatsapp/server"
)

func main() {
	r := server.NewServer()
	handlers.RegisterRoutes(r.Router)

	http.ListenAndServe(":8081", r)
}
