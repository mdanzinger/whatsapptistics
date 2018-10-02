package db

import (
	"database/sql"
	"fmt"

	"github.com/mdanzinger/whatsapp/internal/pkg/reports"
)

type ReportStore struct {
	db *sql.DB
}

func (s *ReportStore) Create(r *reports.Report) (bool, error) {
	fmt.Printf("Creating report ID: %v ", r.ReportID)
	return true, nil
}

func (s *ReportStore) Read(r int) (*reports.Report, error) {
	fmt.Printf("Getting report ID: %v ", r)
	return &reports.Report{}, nil
}

func (s *ReportStore) Update(r *reports.Report) (bool, error) {
	fmt.Printf("Updating report ID: %v ", r.ReportID)
	return true, nil
}

func (s *ReportStore) Delete(r *reports.Report) (bool, error) {
	fmt.Printf("Deleting report ID: %v ", r.ReportID)
	return true, nil
}

func NewReportStore(conn *sql.DB) *ReportStore {
	return &ReportStore{
		db: conn,
	}
}
