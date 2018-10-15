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
	ReportID        string `json:"ReportID"`
	Email           string
	Content         []byte          `json:"content"`
	ReportAnalytics ReportAnalytics `json:"report_analytics"`
}

type ReportAnalytics struct {
	SomeCoolStats string `json:"somethingCool"`
	//MostFrequentWords types.Words
}

type ReportServer interface {
	Upload(context.Context, *Report) error
	Get(context.Context, string) (*Report, error)
	ReportNotifier
}

type ReportAnalyzerService interface {
	Start()
}

type ReportStore interface {
	Get(context.Context, string) (*Report, error)
	Store(context.Context, *Report) error
}

type ReportStorage interface {
	Download(string) ([]byte, error)
	Upload(context.Context, *Report) error
}

//ReportPoller polls a service and returns a slice of report ids.
type ReportPoller interface {
	Poll() ([]Report, error)
}

type ReportAnalyzer interface {
	Analyze(*Report) (*ReportAnalytics, error)
}

type ReportNotifier interface {
	Notify(context.Context, *Report) error
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
