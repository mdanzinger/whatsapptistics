package analyzer

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/mdanzinger/whatsapptistics/src/job"

	"github.com/mdanzinger/whatsapptistics/src/chat"

	"github.com/mdanzinger/whatsapptistics/src/report"

	"github.com/mdanzinger/whatsapptistics/src/mock"
)

var ReportStore = make(map[string]*report.Report)
var ChatStore = make(map[string]*chat.Chat)

func TestAnalyzerService_Start(t *testing.T) {
	// Inject mock services into analyzer service
	var rs mock.ReportService
	var cs mock.ChatService
	var js mock.AnalyzeJobSource
	var as analyzerService

	as.rs = &rs
	as.cs = &cs
	as.jobs = &js
	as.logger = log.New(os.Stdout, "whatsapptistics: ", log.Lshortfile)

	// Create mock funcs
	rs.NewFunc = func(r *report.Report) error {
		ReportStore[r.ReportID] = r
		rs.NewFuncInvoked = true
		return nil
	}

	jobIndex := 0
	js.NextJobFn = func() (*job.AnalyzeJob, error) {
		jobs := []job.AnalyzeJob{
			{ChatID: "1159"},
			{ChatID: "1160"},
		}
		js.NextJobFnInvoked = true
		ji := jobIndex
		jobIndex++
		return &jobs[ji], nil
	}

	cs.RetrieveFn = func(id string) (*chat.Chat, error) {
		if id == "1159" {
			cc, err := ioutil.ReadFile("../../resource/android_testchat.txt")
			if err != nil {
				t.Errorf("Error opening up android test chat")
			}
			c := &chat.Chat{
				Content: cc,
			}
			cs.RetrieveInvoked = true
			return c, nil
		}
		if id == "1160" {
			cc, err := ioutil.ReadFile("../../resource/ios_testchat.txt")
			if err != nil {
				t.Errorf("Error opening up ios test chat")
			}
			c := &chat.Chat{
				Content: cc,
			}
			cs.RetrieveInvoked = true
			return c, nil
		}
		t.Fatalf("Exprect chat id to be either 1159 or 1160, got %s", id)
		return nil, nil
	}

	// Mock Start Method
	for i := 0; i < 2; i++ {
		j, err := as.jobs.NextJob()

		if err != nil {
			as.logger.Printf("Error fetching next job: %v \n", err)
		}
		as.handler(j.ChatID)
	}

	if !rs.NewFuncInvoked {
		t.Errorf("Expected New() to be Invoked")
	}

	if !js.NextJobFnInvoked {
		t.Errorf("Expected NextAnalyzeJobFn() to be Invoked")
	}

	if !cs.RetrieveInvoked {
		t.Errorf("Expected retrive() to be Invoked")
	}

	if ReportStore["1159"] == nil {
		t.Errorf("ReportStore should have report 1159")
	}

	if ReportStore["1160"] == nil {
		t.Errorf("ReportStore should have report 1160")
	}

}
