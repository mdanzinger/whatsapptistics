package analyzer

import (
	"log"
	"sync"

	"github.com/mdanzinger/whatsapptistics/internal/notify"

	"github.com/mdanzinger/whatsapptistics/internal/chat"
	"github.com/mdanzinger/whatsapptistics/internal/report"

	"github.com/mdanzinger/whatsapptistics/internal/job"
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
	reportService report.ReportService
	chatService   chat.ChatService
	emailNotifier *notify.EmailNotifier

	logger *log.Logger
	jobs   job.Source
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
				as.handler(j)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// handler gets chat from an injected chatService and starts the analyzer
func (as *analyzerService) handler(j *job.Chat) {
	as.logger.Printf("Handling chat %s \n", j.ChatID)

	// Download chat
	c, err := as.chatService.Retrieve(j.ChatID)
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
		ReportID:      j.ChatID,
	}

	// Create new report
	err = as.reportService.New(r)
	if err != nil {
		as.logger.Printf("Error storing report: %v \n", err)
	}

	// Notify chat owner
	if len(j.UploaderEmail) > 0 {
		if err := as.emailNotifier.Notify(j.ChatID, j.UploaderEmail); err != nil {
			as.logger.Fatalf("error sending notification to %s, got error: %v", j.UploaderEmail, err)
		}
	}

}

// analyze analyzes the supplied chat and returns a ChatAnalytics
func (as *analyzerService) analyze(c *chat.Chat) (*report.ChatAnalytics, error) {
	as.logger.Println("beginning analyze method")
	var analytics *report.ChatAnalytics
	var analyzer = analyzer{}

	if stripCtlAndExtFromUTF8(string(c.Content[:8]))[0] == '[' { // Is IOS chat
		analyzer.parser = &iosParser{}
		as.logger.Println("using ios parser")
	} else {
		analyzer.parser = &androidParser{}
		as.logger.Println("using android parser")
	}

	as.logger.Println("starting analyzer")
	analytics, err := analyzer.Analyze(c)
	if err != nil {
		return nil, err
	}
	//fmt.Println(analytics)

	return analytics, nil
}

// NewAnalyzerService returns an instance of an analyzer service
func NewAnalyzerService(rs report.ReportService, cs chat.ChatService, jobSource job.Source, notifier notify.Notifier, logger *log.Logger) *analyzerService {
	return &analyzerService{
		reportService: rs,
		chatService:   cs,
		jobs:          jobSource,
		emailNotifier: notify.NewNotifier(notifier),
		logger:        logger,
	}
}
