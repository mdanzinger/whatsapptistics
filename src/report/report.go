package report

import (
	"context"
	"time"
)

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
	Participants map[string]ParticipantAnalytics
	HourList     Hourlist

	// HourMap is used by the analyzer, it is not stored.
	HourMap map[string]int `json:"-"`
}

// ParticipantAnalytics represents the analytics for a single participant in a chat
type ParticipantAnalytics struct {
	//Name         string
	WordsSent    int
	MessagesSent int
	WordList     Wordlist
	MonthList    Monthlist

	// These fields are used by the analyzer, they are not stored.
	WordMap  map[string]int `json:"-"`
	MonthMap map[string]int `json:"-"`
}

// Hour represents a single hour
type Hour struct {
	Hour     string
	Messages int
}

// Hourlist is a slice of hours that implements the sort interface to sort by hour.
type Hourlist []Hour

func (hl Hourlist) Len() int      { return len(hl) }
func (hl Hourlist) Swap(i, j int) { hl[i], hl[j] = hl[j], hl[i] }
func (hl Hourlist) Less(i, j int) bool {
	// Convert to numerical format to have an int to sort by
	ti, _ := time.Parse("3:04 PM", hl[i].Hour)
	tj, _ := time.Parse("3:04 PM", hl[j].Hour)
	return ti.Format("15") < tj.Format("15")
}

// Word represents a unique word
type Word struct {
	Word  string
	Usage int
}

// Wordlist is a slice of words that implements sort interface to sort by count
type Wordlist []Word

func (wl Wordlist) Len() int           { return len(wl) }
func (wl Wordlist) Swap(i, j int)      { wl[i], wl[j] = wl[j], wl[i] }
func (wl Wordlist) Less(i, j int) bool { return wl[i].Usage > wl[j].Usage }

// Month represents a single month
type Month struct {
	Month    string
	Messages int
}

// Monthlist is a slice of months that implements the sort interface to sort by month
type Monthlist []Month

func (ml Monthlist) Len() int      { return len(ml) }
func (ml Monthlist) Swap(i, j int) { ml[i], ml[j] = ml[j], ml[i] }
func (ml Monthlist) Less(i, j int) bool {
	// Convert to numerical format to have an int to sort by
	ti, _ := time.Parse("Jan 2006", ml[i].Month)
	tj, _ := time.Parse("Jan 2006", ml[j].Month)
	return ti.Format("0601") < tj.Format("0601")
}
