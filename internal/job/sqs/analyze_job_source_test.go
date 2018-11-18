package sqs

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/mdanzinger/whatsapptistics/internal/job"
)

func TestNewAnalyzeJobSource(t *testing.T) {
	b := NewAnalyzeJobSource(nil)
	for i := 0; i < 10; i++ {
		testchat := job.Chat{
			ChatID: []string{"1159-" + strconv.Itoa(i), "1160-" + strconv.Itoa(i)}[rand.Intn(2)],
		}
		err := b.QueueJob(&testchat)
		if err != nil {
			t.Error(err)
		}
	}

	time.Sleep(time.Second * 2)

	for i := 0; i < 10; i++ {
		j, err := b.NextJob()
		if err != nil {
			t.Error(err)
		}
		fmt.Println(j.ChatID)
	}

	//j2, err := b.NextJob()
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println(j2.ChatID)
}
