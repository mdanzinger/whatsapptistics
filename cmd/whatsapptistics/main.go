package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/mdanzinger/whatsapptistics/job/sqs"

	"github.com/mdanzinger/whatsapptistics/chat"
	"github.com/mdanzinger/whatsapptistics/http"
	"github.com/mdanzinger/whatsapptistics/report"
	"github.com/mdanzinger/whatsapptistics/store/dynamodb"
	"github.com/mdanzinger/whatsapptistics/store/gocache"
	"github.com/mdanzinger/whatsapptistics/store/s3"
)

func main() {
	//Initialize stuff
	l := log.New(os.Stdout, "whatsapptistics: ", 2)

	// Chat service initialization
	cr := s3.NewChatRepo()
	jobSource := sqs.NewJobSource(l)
	cs := chat.NewChatService(cr, jobSource, l)

	// Report Service initialization
	rc := gocache.NewReportCache()
	rr := dynamodb.NewReportRepo(rc)
	rs := report.NewReportService(rr, l)

	// Create server, inject services
	s := http.NewServer(cs, rs, l)

	s.Start()

}
