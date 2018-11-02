package main

import (
	"log"
	"os"

	"github.com/mdanzinger/whatsapptistics/src/chat"
	"github.com/mdanzinger/whatsapptistics/src/http"
	"github.com/mdanzinger/whatsapptistics/src/report"
	"github.com/mdanzinger/whatsapptistics/src/store/dynamodb"
	"github.com/mdanzinger/whatsapptistics/src/store/redis"
	"github.com/mdanzinger/whatsapptistics/src/store/s3"
)

func main() {
	//Initialize stuff
	l := log.New(os.Stdout, "whatsapptistics:", 2)

	// Chat service initialization
	cr := s3.NewChatRepo()
	cs := chat.NewChatService(cr, l)

	// Report Service initialization
	rc := redis.NewReportCacheRepo()
	rr := dynamodb.NewReportRepo(rc)
	rs := report.NewReportService(rr, l)

	// Create server, inject services
	s := http.NewServer(cs, rs, l)

	s.Start()

}
