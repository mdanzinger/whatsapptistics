package report

import "context"

type Report struct {
	ReportID      string `json:"ReportID"`
	Email         string `json:"email"`
	ChatAnalytics `json:"report_analytics"`
}

type ChatAnalytics map[string]interface{}

type ReportRepository interface {
	Get(ctx context.Context, id string) (*Report, error)
	Store(*Report) error
}
