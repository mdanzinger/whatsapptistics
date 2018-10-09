package main

import (
	"github.com/mdanzinger/mywhatsapp/internal/app/mywhatsapp/handlers"
	"github.com/mdanzinger/mywhatsapp/internal/app/mywhatsapp/http"
)

func main() {
	// Create http and setup routes
	s := http.NewServer()
	handlers.Setup(s)

	// Start http
	s.Start(":8080")
}
