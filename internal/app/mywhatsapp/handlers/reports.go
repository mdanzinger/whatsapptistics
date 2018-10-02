package handlers

import (
	"github.com/mdanzinger/whatsapp/internal/pkg/store"
	"net/http"

	"github.com/mdanzinger/whatsapp/internal/pkg/reports"
)

// New Report Creates a New Report
func newReport(w http.ResponseWriter, r *http.Request) {
	// Some validation..

	store := store.NewReportStore()
	report := reports.NewReport(store)

	report.Create(&reports.Report{
		ReportID: "123",
		Content: reports.ReportData{
			Name: "Some Reporttt",
		},
	})

}
