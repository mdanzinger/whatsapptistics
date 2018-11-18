package analyzer

import (
	"bufio"
	"bytes"
	"sort"
	"strings"
	"unicode"

	"github.com/mdanzinger/whatsapptistics/internal/chat"
	"github.com/mdanzinger/whatsapptistics/internal/report"
)

type analyzer struct {
	parser parser
}

// Analyze analyzes a chat with a given parser
func (a *analyzer) Analyze(chat *chat.Chat) (*report.ChatAnalytics, error) {
	r := &report.ChatAnalytics{}
	participants := map[string]report.ParticipantAnalytics{}
	hourMap := map[string]int{}

	scanner := bufio.NewScanner(strings.NewReader(string(chat.Content)))

	// We use this to store the previous participant because some messages span multiple lines the previous participant because some messages span multiple lines the previous participant because some messages span multiple lines the previous participant because some messages span multiple lines
	var prevParticipant string

	for scanner.Scan() {
		line := stripCtlAndExtFromUTF8(scanner.Text())

		// If it's not valid, it's a continuation of the previous message
		if !a.parser.Valid(line) {
			// We know the prevParticipant exists, no need to handle errors
			p, _ := participants[prevParticipant]
			words := strings.Fields(line)
			p.WordsSent += len(words)

			r.WordsSent += len(words)

			// check and add words to wordmap
			wordMap(&p, []byte(line))

			participants[prevParticipant] = p

			continue
		}

		// Parse line
		date := a.parser.Date(line)
		hour := a.parser.Hour(line)
		message := a.parser.Message(line)
		sender := a.parser.Sender(line)

		// increase message and word count
		r.MessagesSent++

		words := strings.Fields(string(message))
		r.WordsSent += len(words)

		// Increment HourMap
		// check if we have the hour in our map
		if _, ok := hourMap[string(hour)]; !ok {
			hourMap[string(hour)] = 1
		} else {
			hourMap[string(hour)]++
		}

		// Check if the sender is in our participants map
		p, ok := participants[string(sender)]
		if !ok { // Participant not found, lets create a new participant and it to our map!
			participants[string(sender)] = report.ParticipantAnalytics{WordMap: map[string]int{}, MonthMap: map[string]int{}}
			// Hashmaps are fast... I'm just doing a lookup, but I should probably TODO: create the ParticipantsAnalytics object beforehand, and just insert it once after the conditional
			p = participants[string(sender)]
		}

		// Increment messages sent and words sent
		p.MessagesSent++
		p.WordsSent = p.WordsSent + len(words)

		// check and add words to wordmap
		wordMap(&p, message)

		// Increment MonthMap counter for participant
		if _, ok := p.MonthMap[string(date)]; !ok {
			p.MonthMap[string(date)] = 1
		} else {
			p.MonthMap[string(date)]++
		}

		// Set prevParticipant in case of multi line messages
		prevParticipant = string(sender)

		// Add participant data to map
		participants[string(sender)] = p

	}

	r.Participants = participants
	r.HourMap = hourMap

	// Create Wordlist and Monthlist from maps
	a.generateLists(r)

	return r, nil
}

// wordMap iterates over all words in a message, removes any stopwords, and populates the participants wordmap accordingly.
func wordMap(p *report.ParticipantAnalytics, message []byte) {
	// Remove punctuation marks from our message!
	m := bytes.Map(cleanMessage, message)
	words := strings.Fields(string(m))

	if p.MonthMap == nil {
		p.WordMap = make(map[string]int)
	}

	for _, word := range words {
		word := strings.ToLower(word)
		if _, ok := stopwords[word]; !ok {
			_, ok := p.WordMap[word]
			if !ok {
				p.WordMap[word] = 1
			} else {
				p.WordMap[word]++
			}
		}
	}
}

// Generate lists converts the maps into sorted lists
func (a *analyzer) generateLists(r *report.ChatAnalytics) {
	hl := make(report.Hourlist, len(r.HourMap))
	h := 0

	for k, v := range r.HourMap {
		hl[h] = report.Hour{Hour: k, Messages: v}
		h++
	}
	sort.Sort(hl)

	r.HourList = hl

	for k, p := range r.Participants {
		wl := make(report.Wordlist, len(p.WordMap))
		ml := make(report.Monthlist, len(p.MonthMap))

		w := 0
		for k, v := range p.WordMap {
			wl[w] = report.Word{Word: k, Usage: v}
			w++
		}

		m := 0
		for k, v := range p.MonthMap {
			ml[m] = report.Month{Month: k, Messages: v}
			m++
		}

		// Sort the lists
		sort.Sort(wl)
		sort.Sort(ml)

		// Assign lists to the participant
		if len(wl) > 100 {
			p.WordList = wl[:100] // Get 100 most used words!
		} else {
			p.WordList = wl
		}

		p.MonthList = ml

		// Create temp variable to assign to participant map
		temp := r.Participants[k]
		temp.WordList = p.WordList
		temp.MonthList = p.MonthList

		r.Participants[k] = temp
	}
}

// clean message is our mapping func to remove punctuation marks
func cleanMessage(r rune) rune {
	switch {
	case r == '\'' || r == ',' || r == '.' || r == '!' || r == '"' || r == '>' || r == '<' || r == '?' || r == '(' || r == ')' || r == ':' || unicode.IsDigit(r):
		return -1
	}
	return r
}

// This strips any non printable runes. IOS exports were causing issues with LRM marks
func stripCtlAndExtFromUTF8(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, str)
}
