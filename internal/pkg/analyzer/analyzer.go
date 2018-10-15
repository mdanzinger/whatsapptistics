package analyzer

import (
	"fmt"
	"github.com/mdanzinger/mywhatsapp/internal/pkg/report"
)

type analyzer struct {
}

func (a *analyzer) Analyze(r *report.Report) (*report.ReportAnalytics, error) {
	fmt.Println("--------------------- analyzing ------------------")
	return &report.ReportAnalytics{}, nil
}

func NewAnalyzer() *analyzer {
	return &analyzer{}
}
