package handlers

import (
	"github.com/mdanzinger/whatsapp/internal/pkg/store"
	"net/http"

	"github.com/mdanzinger/whatsapp/internal/pkg/report"
)

// New Report Creates a New Report
func newReport(w http.ResponseWriter, r *http.Request) {
	// Some validation..

	store := store.NewReportStore()
	report := report.NewReport(store)

	report.Create(&report.Report{
		ReportID: "123",
		Content: report.ReportData{
			Name: "Some Reporttt",
		},
	})

}
