package main

import (
	"github.com/mdanzinger/whatsapp/internal/app/mywhatsapp/handlers"
	"github.com/mdanzinger/whatsapp/internal/app/mywhatsapp/server"
)

func main() {
	// Create server and setup routes
	s := server.NewServer()
	handlers.Setup(s)

	// Start server
	s.Start(":8080")
}
