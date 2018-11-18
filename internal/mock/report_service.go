package mock

import (
	"context"

	"github.com/mdanzinger/whatsapptistics/internal/report"
)

//ReportService represents a mock of report.ReportService
type ReportService struct {
	GetFunc        func(ctx context.Context, id string) (*report.Report, error)
	GetFuncInvoked bool

	NewFunc        func(*report.Report) error
	NewFuncInvoked bool
}

// Get invokes the mock implementation and marks the function as invoked
func (rs *ReportService) Get(ctx context.Context, id string) (*report.Report, error) {
	rs.GetFuncInvoked = true
	return rs.GetFunc(ctx, id)
}

// New invokes the mock implementation and marks the function as invoked
func (rs *ReportService) New(r *report.Report) error {
	rs.NewFuncInvoked = true
	return rs.NewFunc(r)
}
