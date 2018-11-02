package analyzer

import (
	"fmt"
	"log"

	"github.com/mdanzinger/whatsapptistics/src/chat"
	"github.com/mdanzinger/whatsapptistics/src/report"
)

// ChatAnalyzer represents a chat analyzer
type ChatAnalyzer interface {
	Analyze(*chat.Chat) (*report.ChatAnalytics, error)
}

// analyzerService represents a service for analyzing chats via an injected analyzer
type analyzerService struct {
	repo   report.ReportRepository
	logger *log.Logger
}

// AnalyzeAndStore analyzes the chat and stores the report in an injected repo
func (as *analyzerService) AnalyzeAndStore(c *chat.Chat) error {
	var analytics *report.ChatAnalytics
	var analyzer = analyzer{}

	if c.Content[0] == '[' { // Is IOS chat
		analyzer.parser = &iosParser{}
	} else {
		analyzer.parser = &androidParser{}
	}

	analytics, err := analyzer.Analyze(c)
	if err != nil {
		as.logger.Printf("Error analyzing chat: %v", err)
	}
	fmt.Println(analytics)

	var r = &report.Report{
		ChatAnalytics: *analytics,
	}

	err = as.repo.Store(r)
	if err != nil {
		as.logger.Printf("Error storing report: %v", err)
	}

	return nil
}

// NewAnalyzerService returns an instance of an analyzer service
func NewAnalyzerService(r report.ReportRepository, logger *log.Logger) *analyzerService {
	return &analyzerService{
		repo:   r,
		logger: logger,
	}
}
