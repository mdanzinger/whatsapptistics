package main

import (
	"github.com/mdanzinger/whatsapptistics/analyzer"
	"log"
)

func main() {
	a := analyzer.NewAnalyzerService(nil, log.Logger{})
}
