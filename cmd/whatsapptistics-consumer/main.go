package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/mdanzinger/whatsapptistics/notify/ses"

	"github.com/mdanzinger/whatsapptistics/chat"
	"github.com/mdanzinger/whatsapptistics/job/sqs"
	"github.com/mdanzinger/whatsapptistics/report"
	"github.com/mdanzinger/whatsapptistics/store/gocache"
	"github.com/mdanzinger/whatsapptistics/store/s3"

	"github.com/mdanzinger/whatsapptistics/store/dynamodb"

	"github.com/mdanzinger/whatsapptistics/analyzer"
)

func main() {
	// Create dependencies
	logger := log.New(os.Stdout, "whatsapptistics: ", 0)

	cr := s3.NewChatRepo()
	jobSource := sqs.NewJobSource(logger)
	cs := chat.NewChatService(cr, jobSource, logger)
	rc := gocache.NewReportCache()
	rp := dynamodb.NewReportRepo(rc)
	rs := report.NewReportService(rp, logger)

	mailer := ses.NewSESNotifier()

	a := analyzer.NewAnalyzerService(rs, cs, jobSource, mailer, logger)

	a.Start()

}
