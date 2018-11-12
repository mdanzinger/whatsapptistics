package analyzer

import (
	"log"
	"sync"

	"github.com/mdanzinger/whatsapptistics/src/chat"
	"github.com/mdanzinger/whatsapptistics/src/report"
)

const (
	MAX_CONCURRENT = 10
)

// ChatAnalyzer represents a chat analyzer
type ChatAnalyzer interface {
	Analyze(*chat.Chat) (*report.ChatAnalytics, error)
}

// analyzerService represents a service for analyzing chats via an injected analyzer
type analyzerService struct {
	rs report.ReportService
	cs chat.ChatService

	logger *log.Logger
	poller report.Poller
}

// Start starts and initializes the analyzer service. It will use the injected poller
// to poll and get chats from an message queue
func (as *analyzerService) Start() {
	chatIDs := make(chan []string)
	semaphore := make(chan int, MAX_CONCURRENT) // we use this to limit the amount of analyzers running
	wg := &sync.WaitGroup{}                     // we use this to prevent polling while chats jobs are still running

	// Begin polling!
	go as.poller.Poll(chatIDs, wg)

	//Listen on supplied channels for a slice of ids to come in
	for sliceIDs := range chatIDs {
		for _, id := range sliceIDs {
			go as.handler(id, semaphore, wg)
			//go func(id string, sem chan int, wg *sync.WaitGroup) {
			//	sem <- 1
			//	fmt.Println("Handled Chat")
			//	wg.Done()
			//	<-sem
			//}(id, semaphore, wg)
		}
	}

}

// handler gets chat from an injected chatService and starts the analyzer
func (as *analyzerService) handler(id string, sem chan int, wg *sync.WaitGroup) {
	sem <- 1 // Pass value to semaphore
	as.logger.Printf("Handling chat %s \n", id)

	// Download chat
	c, err := as.cs.Retrieve(id)
	if err != nil {
		as.logger.Printf("Error downloading chat: %v", err)
	}

	// Analyze and store result
	analytics, err := as.analyze(c)
	if err != nil {
		as.logger.Printf("Error analzying: %v \n", err)
	}

	// Create report obj
	r := &report.Report{
		ChatAnalytics: *analytics,
		ReportID:      id,
	}

	// Create new report
	err = as.rs.New(r)
	if err != nil {
		as.logger.Printf("Error storing report: %v \n", err)
	}

	// release waitgroup and semaphore
	wg.Done()
	<-sem
}

// analyze analyzes the supplied chat and returns a ChatAnalytics
func (as *analyzerService) analyze(c *chat.Chat) (*report.ChatAnalytics, error) {
	var analytics *report.ChatAnalytics
	var analyzer = analyzer{}

	if c.Content[0] == '[' { // Is IOS chat
		analyzer.parser = &iosParser{}
	} else {
		analyzer.parser = &androidParser{}
	}

	analytics, err := analyzer.Analyze(c)
	if err != nil {
		return nil, err
	}
	//fmt.Println(analytics)

	return analytics, nil
}

// NewAnalyzerService returns an instance of an analyzer service
func NewAnalyzerService(rs report.ReportService, cs chat.ChatService, logger *log.Logger, poller report.Poller) *analyzerService {
	return &analyzerService{
		rs:     rs,
		cs:     cs,
		logger: logger,
		poller: poller,
	}
}
