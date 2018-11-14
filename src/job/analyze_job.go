package job

// AnalyzeJob represents a chat that needs to be analyzed and created into a report
type AnalyzeJob struct {
	ChatID        string
	UploaderEmail string
}

//AnalyzeJobSource represents a queue of jobs
type AnalyzeJobSource interface {
	QueueJob(job *AnalyzeJob) error
	NextJob() (*AnalyzeJob, error)
}
