package main

import (
	"log"
	"os"

	"github.com/mdanzinger/whatsapptistics/internal/notify/ses"

	"github.com/mdanzinger/whatsapptistics/internal/chat"
	"github.com/mdanzinger/whatsapptistics/internal/job/sqs"
	"github.com/mdanzinger/whatsapptistics/internal/report"
	"github.com/mdanzinger/whatsapptistics/internal/store/gocache"
	"github.com/mdanzinger/whatsapptistics/internal/store/s3"

	"github.com/mdanzinger/whatsapptistics/internal/store/dynamodb"

	"github.com/mdanzinger/whatsapptistics/internal/analyzer"
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
