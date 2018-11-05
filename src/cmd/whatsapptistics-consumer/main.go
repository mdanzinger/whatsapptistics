package main

import (
	"github.com/mdanzinger/whatsapptistics/src/analyzer"
	"github.com/mdanzinger/whatsapptistics/src/store/dynamodb"
	"github.com/mdanzinger/whatsapptistics/src/store/redis"
	"log"
)

func main() {
	rc := redis.NewReportCacheRepo()
	rp := dynamodb.NewReportRepo(rc)
	a := analyzer.NewAnalyzerService(rp, &log.Logger{})


}
