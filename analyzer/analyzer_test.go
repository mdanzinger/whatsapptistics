package analyzer

import (
	"io/ioutil"
	"testing"

	"github.com/mdanzinger/whatsapptistics/chat"
)

func TestAnalyzer_Analyze(t *testing.T) {
	a := analyzer{
		parser: &androidParser{},
		//parser: &iosParser{},
	}
	androidChat, err := ioutil.ReadFile("../resource/android_testchat.txt")
	//androidChat, err := ioutil.ReadFile("../resource/_test.txt")

	if err != nil {
		t.Fatal("Error opening up ios test chat")
	}

	c := &chat.Chat{
		Content: androidChat,
	}

	r, err := a.Analyze(c)

	if err != nil {
		t.Fatalf("Got error analyzing chat")
	}

	if r.WordsSent != 127 {
		t.Errorf("WordsSent for test chat should be '127', is %v", r.WordsSent)
	}

	// TODO implement full coverage of the analyzer
	//if r.WordsSent != 117 {
	//	t.Fatalf("Report has counted %v words, should be %v", r.WordsSent, 117)
	//}

	// print chat for debugging
	//j, err := json.Marshal(r)
	//if err != nil {
	//	fmt.Println("Could not marshal json")
	//}

	//fmt.Println(string(j))
}
