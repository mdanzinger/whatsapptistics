package analyzer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/mdanzinger/whatsapptistics/src/chat"
)

func TestAnalyzer_Analyze(t *testing.T) {
	a := analyzer{
		//parser: &androidParser{},
		parser: &iosParser{},
	}
	//android_chat, err := ioutil.ReadFile("../../resource/android_testchat.txt")
	android_chat, err := ioutil.ReadFile("../../resource/_chat-ios.txt")

	if err != nil {
		t.Errorf("Error opening up android test chat")
	}

	c := &chat.Chat{
		Content: android_chat,
	}

	r, err := a.Analyze(c)

	// TODO implement full coverage of the analyzer
	//if r.WordsSent != 117 {
	//	t.Fatalf("Report has counted %v words, should be %v", r.WordsSent, 117)
	//}



	// print chat for debugging
	j, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Could not marshal json")
	}

	fmt.Println(string(j))
}
