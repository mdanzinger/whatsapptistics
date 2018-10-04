package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	//"github.com/mdanzinger/whatsapp/mywhatsapp/common"
	//"github.com/mdanzinger/whatsapp/mywhatsapp/report"
	"whatsapp/app/common"
	"whatsapp/app/reports"
)

func main() {
	//// Initializes AWS Session
	common.Init()

	r := mux.NewRouter()
	reports.RegisterRoutes(r)

	//r.

	log.Fatal(http.ListenAndServe(":8000", r))

}
