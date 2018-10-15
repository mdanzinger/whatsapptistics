package report

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type reportAnalyzerService struct {
	store     ReportStore    // Db transitions
	filestore ReportStorage  // S3, etc
	poller    ReportPoller   // sqs
	analyzer  ReportAnalyzer // custom analyzer... this can be swapped out in the future for new and improved analyzers
}

// Start starts the analyzer service
func (ra *reportAnalyzerService) Start() {
	log.Println("Starting Report Analyzer")
	for {
		// Poll for messages using the injected poller (sqs most likely).
		reports, err := ra.poller.Poll()
		if err != nil {
			log.Println(err)
			break
		}

		// We have some reports to process!
		if len(reports) > 0 {
			ra.process(reports) //
		}
	}
}

// processReports spawns goroutines to generate the report.
func (ra *reportAnalyzerService) process(reports []Report) {
	log.Println("Received %d reports", len(reports))
	// Create waitgroup. We don't want processReports to continue receiving reports until we've processed all reports given.
	var wg sync.WaitGroup
	wg.Add(len(reports))

	// Loop over each report and spawn a goroutine to handle the report
	for i := range reports {
		go func(r Report) {
			defer wg.Done()
			log.Println("worker spawned, processing report")

			rContent, err := ra.filestore.Download(r.ReportID)
			if err != nil {
				fmt.Println(err)
			}
			r.Content = rContent

			ra.analyze(&r)
			time.Sleep(time.Second * 8)
		}(reports[i])
	}
	wg.Wait()
	fmt.Println("Completed analysis, Looking for more!")
}

// AnalyzeChat analyzes the chat using the injected analyzer
func (ra *reportAnalyzerService) analyze(r *Report) error {
	analytics, err := ra.analyzer.Analyze(r)
	if err != nil {
		fmt.Println(err)
		return err
	}
	r.ReportAnalytics = *analytics
	r.Content = nil

	err = ra.store.Store(context.Background(), r)
	if err != nil {
		return err
	}
	return nil
}

// NewReportAnalyzer returns a ReportAnalyzerService
func NewReportAnalyzer(store ReportStore, filestore ReportStorage, poller ReportPoller, analyzer ReportAnalyzer) *reportAnalyzerService {
	return &reportAnalyzerService{
		store:     store,
		filestore: filestore,
		poller:    poller,
		analyzer:  analyzer,
	}
}
