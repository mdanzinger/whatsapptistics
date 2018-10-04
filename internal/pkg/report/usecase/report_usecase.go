package usecase

import "github.com/mdanzinger/whatsapp/internal/pkg/report"

type reportUsecase struct {
	reportRepo  report.ReportRepository
	reportCache report.ReportCache
}

func NewReportUsecase(r report.ReportRepository, c report.ReportCache) report.ReportUsecase {
	return &reportUsecase{
		reportRepo:  r,
		reportCache: c,
	}
}

func (r *reportUsecase) Create(report2 *report.Report) error {

}
