package report

import (
	"github.com/mdanzinger/whatsapptistics/src/chat"
	"log"
	"sync"
)

const (
	// MAX_CONCURRENT is the maximum amount of chats that'll be analyzed concurrently
	MAX_CONCURRENT = 10
)


// chatID represents an ID of a chat that's queued to be analyzed
type chatID string


// Poller is an interface for polling against some service and returning a slice of chat id's
// to be analyzed
type Poller interface {
	Poll(chan []chatID) error
}


type pollService struct {
	poller Poller
	reportService ReportService
	chatService chat.ChatService
	log *log.Logger
}

func (ps *pollService) Start() {
	// Make channel to retrieve chat ids
	chatIDs := make(chan []chatID)
	semaphore := make(chan int, MAX_CONCURRENT)

	if err := ps.poller.Poll(chatIDs); err != nil {
		ps.log.Println("Error polling: ", err)
	}

	for id := range chatIDs {

	}


}

//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//type chatID string
//
//
//func main() {
//	wg := &sync.WaitGroup{}
//	chatIDs := make(chan []chatID)
//	semaphore := make(chan int, 2)
//
//	go Poll(chatIDs, wg)
//
//	for ids := range chatIDs {
//		for _, id := range ids {
//			go handler(id, semaphore, wg)
//		}
//	}
//
//}
//
//func Poll(c chan []chatID, wg *sync.WaitGroup) {
//	for {
//		fmt.Println("Waiting for all chats to be processed before 'fetching' more!")
//		time.Sleep(time.Millisecond * 500)
//		wg.Add(10)
//		c <- []chatID{"chat 1", "chat 2", "chat 3", "chat 4", "chat 5", "chat 6", "chat 7", "chat 8", "chat 9", "chat 10"}
//		wg.Wait()
//	}
//}
//
//func handler(id chatID, sem chan int, wg *sync.WaitGroup) {
//	sem <- 1
//	fmt.Println("Handling", id)
//	time.Sleep(time.Second)
//	wg.Done()
//	<-sem
//}