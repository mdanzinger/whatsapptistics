package sqs

import (
	"log"
	"os"
	"sync"
	"testing"
)

func TestReportPoller_Poll(t *testing.T) {
	ch := make(chan []string)
	wg := &sync.WaitGroup{}
	ps := NewReportPoller(log.New(os.Stdout, "POLLER", log.Lshortfile))

	ps.Poll(ch, wg)

}
