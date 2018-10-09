package main

import (
	"github.com/mdanzinger/mywhatsapp/internal/app/mywhatsapp/http"
	"github.com/mdanzinger/mywhatsapp/internal/pkg/notify/sns"
	"github.com/mdanzinger/mywhatsapp/internal/pkg/report"
	"github.com/mdanzinger/mywhatsapp/internal/pkg/storage/s3"
	"github.com/mdanzinger/mywhatsapp/internal/pkg/store/dynamodb"
	"github.com/mdanzinger/mywhatsapp/internal/pkg/store/redis"
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
