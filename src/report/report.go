package report

import "context"

// Report represents an individual report
type Report struct {
	ReportID      string        `json:"ReportID"`
	Email         string        `json:"email"`
	ChatAnalytics ChatAnalytics `json:"report_analytics"`
}

// ReportRepository represents a repository of reports
type ReportRepository interface {
	Get(ctx context.Context, id string) (*Report, error)
	Store(*Report) error
}

// ChatAnalytics represents the analytics of the chat
type ChatAnalytics struct {
	WordsSent    int
	MessagesSent int
	//MessagesByHour map[int]int // Hour -> Messages sent
	MessagesByHour map[string]int
	Participants   map[string]ParticipantAnalytics
}

// ParticipantAnalytics represents the analytics for a single participant in a chat
type ParticipantAnalytics struct {
	//Name         string
	WordsSent       int
	MessagesSent    int
	MessagesByMonth map[string]int
	WordList        Wordlist

	// These fields are used by the analyzer, they are not marshaled
	WordMap map[string]int `json:"-"`
}

// Word represents a unique word
type Word struct {
	Word  string
	Usage int
}

// Wordlist is a slice of words that implements sort.Interface to sort by count
type Wordlist []Word

func (p Wordlist) Len() int           { return len(p) }
func (p Wordlist) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Wordlist) Less(i, j int) bool { return p[i].Usage < p[j].Usage }

// MessagesByHour represents the amount of messages sent in an hour
type MessagesByHours struct {
	Hour         int // 00 = 12:00 am
	MessagesSent int
}
