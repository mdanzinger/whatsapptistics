package mock

import "sync"

// Poller represents a mock implementation of report.Poller
type Poller struct {
	PollFn func(ch chan []string, wg *sync.WaitGroup)
	PollFnInvoked bool
}

// Poll implements the Poll method of our mock Poller
func (p *Poller) Poll(ch chan []string, wg *sync.WaitGroup ) {
	 p.PollFn(ch, wg)
}