package analyzer

import (
	"log"
	"sync"

	"github.com/mdanzinger/whatsapptistics/src/job"

	"github.com/mdanzinger/whatsapptistics/src/chat"
	"github.com/mdanzinger/whatsapptistics/src/report"
)

const (
	MAX_CONCURRENT = 2
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
	jobs   job.AnalyzeJobSource
}

// Start starts and initializes the analyzer service. It will use the injected poller
// to poll and get chats from an message queue
func (as *analyzerService) Start() {
	wg := &sync.WaitGroup{}
	for {
		for i := 0; i < MAX_CONCURRENT; i++ {
			j, err := as.jobs.NextJob()

			if err != nil {
				as.logger.Printf("Error fetching next job: %v \n", err)
			}

			wg.Add(1)
			go func() {
				as.handler(j.ChatID)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// handler gets chat from an injected chatService and starts the analyzer
func (as *analyzerService) handler(id string) {
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
func NewAnalyzerService(rs report.ReportService, cs chat.ChatService, logger *log.Logger, jobSource job.AnalyzeJobSource) *analyzerService {
	return &analyzerService{
		rs:     rs,
		cs:     cs,
		logger: logger,
		jobs:   jobSource,
	}
}
