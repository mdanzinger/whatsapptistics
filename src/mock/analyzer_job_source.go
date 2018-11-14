package mock

import (
	"github.com/mdanzinger/whatsapptistics/src/job"
)

// Poller represents a mock implementation of report.Poller
type AnalyzeJobSource struct {
	NextJobFn        func() (*job.AnalyzeJob, error)
	NextJobFnInvoked bool

	QueueJobFn        func(*job.AnalyzeJob) error
	QueueJobFnInvoked bool

	// some queue services return a batch of messages, so currentBatch
	// represents a batch of messages (jobs) that may have been returned
	currentBatch []job.AnalyzeJob
}

// Poll implements the Poll method of our mock Poller
func (js *AnalyzeJobSource) NextJob() (*job.AnalyzeJob, error) {
	js.NextJobFnInvoked = true
	return js.NextJobFn()
}

func (js *AnalyzeJobSource) QueueJob(analyzeJob *job.AnalyzeJob) error {
	js.QueueJobFnInvoked = true
	return js.QueueJobFn(analyzeJob)
}
