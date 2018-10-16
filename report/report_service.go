package report

import "log"

// ReportService represents a service for interacting with the Report Repository
type ReportService interface {
	Get(id string) (*Report, error)
	Store(*Report) error
}

type reportService struct {
	rr     ReportRepository
	logger log.Logger
}

// Get retrieves a report from an injected db
func (rs *reportService) Get(id string) (*Report, error) {
	r, err := rs.rr.Get(id)
	if err != nil {
		rs.logger.Print(err)
		return nil, err
	}

	return r, nil
}

// Store stores the report in an injected db
func (rs *reportService) Store(r *Report) error {
	if err := rs.rr.Store(r); err != nil {
		rs.logger.Print(err)
		return err
	}
	return nil
}

// NewReportService returns a ReportService instance with the dependencies injected
func NewReportService(rr ReportRepository, logger log.Logger) *reportService {
	return &reportService{
		rr:     rr,
		logger: logger,
	}
}
