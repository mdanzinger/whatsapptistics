package main

import (
	"github.com/mdanzinger/mywhatsapp/chat"
	"github.com/mdanzinger/mywhatsapp/dynamodb"
	"github.com/mdanzinger/mywhatsapp/http"
	"github.com/mdanzinger/mywhatsapp/redis"
	"github.com/mdanzinger/mywhatsapp/report"
	"github.com/mdanzinger/mywhatsapp/s3"
	"log"
)

func main() {
	//Initialize stuff
	l := &log.Logger{}

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
