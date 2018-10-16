package report

import (
	"context"
	"github.com/nu7hatch/gouuid"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Report struct {
	ReportID      string                 `json:"ReportID"`
	Email         string                 `json:"email"`
	ChatAnalytics map[string]interface{} `json:"report_analytics"`
}

type Chat struct {
}

func NewReport(r io.Reader) *Report {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	// Process ID
	idString := strings.Replace(id.String(), "-", "", -1)

	reportContent, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return &Report{
		ReportID: idString,
		Content:  reportContent,
	}
}
