package handlers

import (
	"database/sql"
	"net/http"

	"github.com/mdanzinger/whatsapp/internal/app/mywhatsapp/db"

	"github.com/mdanzinger/whatsapp/internal/pkg/reports"
)

// New Report Creates a New Report
func newReport(w http.ResponseWriter, r *http.Request) {
	// Some validation..

	store := db.NewReportStore(&sql.DB{})
	report := reports.NewReport(store)

	report.Read(123)

}
