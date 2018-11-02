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
		parser: &androidParser{},
	}
	android_chat, err := ioutil.ReadFile("../../resource/_chat.txt")
	if err != nil {
		t.Errorf("Error opening up android test chat")
	}

	c := &chat.Chat{
		Content: android_chat,
	}

	r, err := a.Analyze(c)

	j, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Could not marshal json")
	}

	fmt.Println(string(j))
}
