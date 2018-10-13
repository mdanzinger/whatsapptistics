package report

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)

type reportAnalyzerService struct {
	store    ReportStore    // Db transitions
	storage  ReportStore    // S3, etc
	poller   ReportPoller   // sqs
	analyzer ReportAnalyzer // custom analyzer... this can be swapped out in the future for new and improved analyzers
}

// Start starts the analyzer service
func (ra *reportAnalyzerService) Start() error {

	emptyError := errors.New("")
	for {
		log.Println("Starting Report Analyzer")
		// Poll for messages using the injected poller (sqs most likely).
		reportsToAnalze, err := ra.poller.Poll()
		if err != nil {
			log.Println(err)
			emptyError = err
			break

		}

		// We have some reports to process!
		if len(reportsToAnalze) > 0 {
			ra.processReports(reportsToAnalze) //
		}
	}
	return emptyError
}

// processReports spawns goroutines to generate the report.
func (ra *reportAnalyzerService) processReports(reports []Report) {
	log.Println("Received %d reports", len(reports))
	// Create waitgroup. We don't want processReports to continue receiving reports until we've processed all reports given.
	var wg sync.WaitGroup
	wg.Add(len(reports))

	// Loop over each report and spawn a goroutine to handle the report
	for i := range reports {
		go func(r Report) {
			log.Println("worker spawned, processing report")
			defer wg.Done()
			ra.AnalyzeAndStore(&r)
		}(reports[i])
	}
}

// AnalyzeChat analyzes the chat using the injected analyzer
func (ra *reportAnalyzerService) AnalyzeAndStore(r *Report) error {
	analytics, err := ra.analyzer.Analyze(r)
	if err != nil {
		fmt.Println(err)
		return err
	}
	r.ReportAnalytics = *analytics

	err = ra.store.Store(context.Background(), r)
	if err != nil {
		return err
	}
	return nil
}

// GetChat retrieves a chat from an injected storage source
func (ra *reportAnalyzerService) GetChat(id string) (*Report, error) {
	ctx := context.Background()
	report, err := ra.storage.Get(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return report, nil
}

func (ra *reportAnalyzerService) StoreReport(ctx context.Context, r *Report) error {
	err := ra.store.Store(ctx, r)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// NewReportAnalyzer returns a ReportAnalyzerService
func NewReportAnalyzer(s ReportStore, st ReportStore, p ReportPoller, a ReportAnalyzer) *reportAnalyzerService {
	return &reportAnalyzerService{
		store:    s,
		storage:  st,
		poller:   p,
		analyzer: a,
	}
}
