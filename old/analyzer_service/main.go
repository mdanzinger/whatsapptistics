package main

import (
	"log"
	"whatsapp/analyzer_service/analyzer"
	"whatsapp/analyzer_service/worker"
)

func main() {
	w, err := worker.NewChatReceiver("whatsapp-chat")
	if err != nil {
		log.Printf("Error creating new Worker Service: %s\n", err)
	}
	log.Printf("Starting SQS Consumer")

	err = analyzer.Init()
	if err != nil {
		log.Printf("Error creating analyzer session: %s\n", err)
	}

	w.Start(worker.HandlerFunc(analyzer.NewChatAnalyzer))
}
