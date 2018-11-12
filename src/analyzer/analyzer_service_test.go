package analyzer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"testing"
	"time"

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
	var poller mock.Poller
	var as analyzerService

	as.rs = &rs
	as.cs = &cs
	as.poller = &poller
	as.logger = log.New(os.Stdout, "whatsapptistics: ", log.Lshortfile)

	// Create mock funcs
	rs.NewFunc = func(r *report.Report) error {
		ReportStore[r.ReportID] = r
		rs.NewFuncInvoked = true
		return nil
	}

	poller.PollFn = func(ch chan []string, wg *sync.WaitGroup) {
		fmt.Println("Polling for messages from Mock Implementation")
		time.Sleep(time.Millisecond * 500)
		wg.Add(2)
		ch <- []string{
			"1159",
			"1160",
		}
		wg.Wait()
		close(ch)
		poller.PollFnInvoked = true
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
		t.Errorf("Exprect chat id to be either 1159 or 1170, got %s", id)
		return nil, nil
	}
	as.Start()

	if !rs.NewFuncInvoked {
		t.Errorf("Expected New() to be Invoked")
	}

	if !poller.PollFnInvoked {
		t.Errorf("Expected Poll() to be Invoked")
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
