package main

import (
	"github.com/mdanzinger/mywhatsapp/old/internalrnal/app/mywhatsapp/http"
	"github.com/mdanzinger/mywhatsapp/old/internalrnal/pkg/notify/sns"
	"github.com/mdanzinger/mywhatsapp/old/internalrnal/pkg/report"
	"github.com/mdanzinger/mywhatsapp/old/internalrnal/pkg/storage/s3"
	"github.com/mdanzinger/mywhatsapp/old/internalrnal/pkg/store/dynamodb"
	"github.com/mdanzinger/mywhatsapp/old/internalrnal/pkg/store/redis"
)

func main() {
	// Create server dependencies
	cacher := redis.NewReportCache()
	db := dynamodb.NewReportStore(cacher)
	notifier := sns.NewReportNotifier()
	storage := s3.NewReportStorage()

	reportServer := report.NewReportServer(db, notifier, storage)

	// Create Server
	s := http.NewServer(reportServer)

	s.ListenAndServe(":8080")

}
