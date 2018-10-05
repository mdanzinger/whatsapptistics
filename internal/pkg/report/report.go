package report

import "context"

type Report struct {
	ReportID string `json:"ReportID"`
	Content  []byte `json:"content"`
}

type ReportAnalytics struct {
	MostFrequentWords types.Words
}

type ReportServer interface {
	Upload(context.Context, *Report) error
	Notify(*Report) error
	Get(string) error
}

type ReportAnalyzer interface {
	Start() error
	Analyze(*Report) *ReportAnalytics
	Store(*ReportAnalytics) error
}
