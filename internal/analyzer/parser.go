package analyzer

import (
	"strings"
	"time"
	"unicode"
)

// parse parses a whatsapp line
type parser interface {
	// Name returns the name of the sender
	Sender(line string) []byte
	// Date returns the Month and year the message was sent
	Date(line string) []byte // January 2018
	// Hour returns the hour the Message was sent
	Hour(line string) []byte // 2:00 AM represents all messages sent between 2:00 AM and 3:00AM
	// Message returns the actual Message
	Message(line string) []byte

	// Valid returns false if the message if the line is not a valid message.
	Valid(line string) bool
}

// iosParser is a parser implementation for IOS
type iosParser struct{}

func (p *iosParser) Sender(line string) []byte {
	return []byte(line[strings.Index(line, "]")+2 : strings.Index(line, ": ")])
}

func (p *iosParser) Date(line string) []byte {
	d := line[strings.Index(line, "[")+1 : strings.Index(line, ",")]
	tp, _ := time.Parse("2006-01-02", d)
	return []byte(tp.Format("Jan 2006"))
}

func (p *iosParser) Hour(line string) []byte {
	t := line[strings.Index(line, ",")+2 : strings.Index(line, "]")]
	tp, _ := time.Parse("3:04:05 PM", t)
	i := []byte(tp.Format("3:00 PM"))
	return i

}
func (p *iosParser) Message(line string) []byte {
	//return []byte(strings.SplitAfterN(line, ": ", 2)[1]) // We need to split to ensure we correctly parse the whole iosMessage, and not break for messages that have colons
	return []byte(line[strings.Index(line, ": ")+2:])
}

func (p *iosParser) Valid(line string) bool {
	return len(line) > 0 && (strings.Index(line, "[")) != -1
}

// androidParser is a parser implementation for android
type androidParser struct{}

func (p *androidParser) Sender(line string) []byte {
	return []byte(line[strings.Index(line, "-")+2 : strings.Index(line, ": ")])
}
func (p *androidParser) Date(line string) []byte {
	d := line[0:strings.Index(line, ",")]
	tp, _ := time.Parse("1/2/06", d)
	return []byte(tp.Format("Jan 2006"))
}
func (p *androidParser) Hour(line string) []byte {
	t := line[strings.Index(line, ",")+2 : strings.Index(line, "-")-1]
	tp, _ := time.Parse("3:04 PM", t)
	i := []byte(tp.Format("3:00 PM"))
	return i
}
func (p *androidParser) Message(line string) []byte {
	return []byte(line[strings.Index(line, ": ")+2:])
}

func (p *androidParser) Valid(line string) bool {
	if len(line) == 0 {
		return false
	}
	char1 := rune(line[0])
	return unicode.IsDigit(char1) && strings.Index(line, " - ") != -1
	//char1 := rune(line[0])
	//char2 := rune(line[1])
	//char3 := rune(line[2])
	//return unicode.IsDigit(char1) && (unicode.IsDigit(char2) || char2 == '/' && (unicode.IsDigit(char3) || char3 == '/'))
}
