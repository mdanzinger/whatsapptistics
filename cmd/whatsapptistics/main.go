package main

import (
	"log"
	"os"

	"github.com/mdanzinger/whatsapptistics/internal/job/sqs"

	"github.com/mdanzinger/whatsapptistics/internal/chat"
	"github.com/mdanzinger/whatsapptistics/internal/http"
	"github.com/mdanzinger/whatsapptistics/internal/report"
	"github.com/mdanzinger/whatsapptistics/internal/store/dynamodb"
	"github.com/mdanzinger/whatsapptistics/internal/store/gocache"
	"github.com/mdanzinger/whatsapptistics/internal/store/s3"
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
