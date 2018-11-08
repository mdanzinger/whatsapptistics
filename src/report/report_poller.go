package report

import (
	"sync"
)

const (
	// MAX_CONCURRENT is the maximum amount of chats that'll be analyzed concurrently
	MAX_CONCURRENT = 10
)

// chatID represents an ID of a chat that's queued to be analyzed
//type chatID string

// Poller is an interface for polling from a queue, and sending a slice of chatIDs
// to a supplied channel.
type Poller interface {
	// Poll sends a slice of strings that represent chat IDS to a supplied channel, and waits for the
	// waitgroup before continuing to poll.
	Poll(chan []string, *sync.WaitGroup)
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
