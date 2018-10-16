package main

import (
	"github.com/mdanzinger/mywhatsapp/analyzer"
	"log"
)

func main() {
	a := analyzer.NewAnalyzerService(nil, log.Logger{})
}
