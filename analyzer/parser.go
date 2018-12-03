package analyzer

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

// parse parses a whatsapp line
type parser interface {
	// Name returns the name of the sender
	Sender(line string) ([]byte, error)
	// Date returns the Month and year the message was sent
	Date(line string) ([]byte, error) // January 2018
	// Hour returns the hour the Message was sent
	Hour(line string) ([]byte, error) // 2:00 AM represents all messages sent between 2:00 AM and 3:00AM
	// Message returns the actual Message
	Message(line string) ([]byte, error)

	// Valid returns false if the message is a continuation of the previous message
	Valid(line string) bool

	//Announcement returns true if message was a group chat announcement (ex: X has added Y to the chat)
	Announcement(line string) bool
}

var (
	ErrNoSender  = fmt.Errorf("error finding sender")
	ErrNoDate    = fmt.Errorf("error finding date")
	ErrNoHour    = fmt.Errorf("error finding hour")
	ErrNoMessage = fmt.Errorf("error finding message")
)

// iosParser is a parser implementation for IOS
type iosParser struct{}

func (p *iosParser) Sender(line string) ([]byte, error) {
	if strings.Index(line, "[") == -1 || strings.Index(line, "]") == -1 {
		return nil, ErrNoSender
	}
	return []byte(line[strings.Index(line, "]")+2 : strings.Index(line, ": ")]), nil
}

func (p *iosParser) Date(line string) ([]byte, error) {
	if strings.Index(line, "[") == -1 || strings.Index(line, "]") == -1 {
		return nil, ErrNoDate
	}
	d := line[strings.Index(line, "[")+1 : strings.Index(line, ",")]
	tp, _ := time.Parse("2006-01-02", d)
	return []byte(tp.Format("Jan 2006")), nil
}

func (p *iosParser) Hour(line string) ([]byte, error) {
	if strings.Index(line, "[") == -1 || strings.Index(line, "]") == -1 || !strings.Contains(line, ",") || strings.Index(line, "]") <= 20 {
		return nil, ErrNoHour
	}
	t := line[strings.Index(line, ",")+2 : strings.Index(line, "]")]
	tp, _ := time.Parse("3:04:05 PM", t)
	i := []byte(tp.Format("3:00 PM"))
	return i, nil

}
func (p *iosParser) Message(line string) ([]byte, error) {
	if strings.Index(line, ":") == -1 {
		return nil, ErrNoMessage
	}
	//return []byte(strings.SplitAfterN(line, ": ", 2)[1]) // We need to split to ensure we correctly parse the whole iosMessage, and not break for messages that have colons
	return []byte(line[strings.Index(line, ": ")+2:]), nil
}

func (p *iosParser) Valid(line string) bool {
	return len(line) > 0 && strings.Contains(line, "[")
}
func (p *iosParser) Announcement(line string) bool {
	return strings.Contains(line, "]") && !strings.Contains(line, ": ")
}

// androidParser is a parser implementation for android
type androidParser struct{}

func (p *androidParser) Sender(line string) ([]byte, error) {
	if !strings.Contains(line, " - ") || strings.Count(line, ":") < 2 {
		return nil, ErrNoSender
	}
	return []byte(line[strings.Index(line, "-")+2 : strings.Index(line, ": ")]), nil
}
func (p *androidParser) Date(line string) ([]byte, error) {
	d := line[0:strings.Index(line, ",")]
	tp, err := time.Parse("1/2/06", d)
	if err != nil {
		return nil, ErrNoDate
	}
	return []byte(tp.Format("Jan 2006")), nil
}
func (p *androidParser) Hour(line string) ([]byte, error) {
	t := line[strings.Index(line, ",")+2 : strings.Index(line, "-")-1]
	tp, err := time.Parse("3:04 PM", t)
	if err != nil {
		return nil, ErrNoHour
	}
	i := []byte(tp.Format("3:00 PM"))
	return i, nil
}
func (p *androidParser) Message(line string) ([]byte, error) {
	if strings.Index(line, ": ")+2 > len(line) {
		return nil, ErrNoMessage
	}
	return []byte(line[strings.Index(line, ": ")+2:]), nil
}

func (p *androidParser) Valid(line string) bool {
	if len(line) == 0 {
		return false
	}
	char1 := rune(line[0])
	return unicode.IsDigit(char1) && strings.Index(line, "M - ") != -1 && strings.Index(line, "M -") <= 20
}

func (p *androidParser) Announcement(line string) bool {
	return strings.Contains(line, "M -") && !strings.Contains(line, ": ")
}
