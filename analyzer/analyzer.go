package analyzer

import (
	"github.com/mdanzinger/mywhatsapp/chat"
	"github.com/mdanzinger/mywhatsapp/report"
	"log"
)

// ChatAnalyzer represents a chat analyzer
type ChatAnalyzer interface {
	Analyze(*chat.Chat) (*report.ChatAnalytics, error)
}

// analyzerService represents a service for analyzing chats via an injected analyzer
type analyzerService struct {
	a      ChatAnalyzer
	logger log.Logger
}

func (as *analyzerService) Analyze(c *chat.Chat) (*report.ChatAnalytics, error) {
	ca, err := as.a.Analyze(c)
	if err != nil {
		as.logger.Print(err)
		return nil, err
	}
	return ca, nil
}

// NewAnalyzerService returns an instance of an analyzer service
func NewAnalyzerService(a ChatAnalyzer, logger log.Logger) *analyzerService {
	return &analyzerService{
		a:      a,
		logger: logger,
	}
}
