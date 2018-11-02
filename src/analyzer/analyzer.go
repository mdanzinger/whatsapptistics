package analyzer

import (
	"bufio"
	"bytes"
	"sort"
	"strings"

	"github.com/mdanzinger/whatsapptistics/src/chat"
	"github.com/mdanzinger/whatsapptistics/src/report"
)

type analyzer struct {
	parser parser
}

// Analyze analyzes a chat with a given parser
func (a *analyzer) Analyze(chat *chat.Chat) (*report.ChatAnalytics, error) {
	r := &report.ChatAnalytics{}
	participants := map[string]report.ParticipantAnalytics{}
	messagesByHour := map[string]int{}

	scanner := bufio.NewScanner(strings.NewReader(string(chat.Content)))
	for scanner.Scan() {
		line := scanner.Text()

		// Parse line
		date := a.parser.Date(line)
		hour := a.parser.Hour(line)
		message := a.parser.Message(line)
		sender := a.parser.Sender(line)

		// increase message and word count
		r.MessagesSent++
		r.WordsSent = r.WordsSent + len(message)

		// Increment MessagesByHour
		// check if we have the hour in our map
		if _, ok := messagesByHour[string(hour)]; !ok {
			messagesByHour[string(hour)] = 1
		} else {
			messagesByHour[string(hour)]++
		}

		// Check if the sender is in our participants map
		p, ok := participants[string(sender)]
		if !ok { // Participant not found, lets create a new participant and it to our map!
			participants[string(sender)] = report.ParticipantAnalytics{WordMap: map[string]int{}, MessagesByMonth: map[string]int{}}
			// Hashmaps are fast... I'm just doing a lookup, but I should probably TODO: create the ParticipantsAnalytics object beforehand, and just insert it once after the conditional
			p = participants[string(sender)]
		}

		// Increment messages sent and words sent
		p.MessagesSent++
		p.WordsSent = p.WordsSent + len(message)

		// check and add words to wordmap
		wordMap(&p, message)

		// Increment MessagesByMonth counter for participant
		if _, ok := p.MessagesByMonth[string(date)]; !ok {
			p.MessagesByMonth[string(date)] = 1
		} else {
			p.MessagesByMonth[string(date)]++
		}

		// Add participant data to map
		participants[string(sender)] = p

	}

	r.Participants = participants
	r.MessagesByHour = messagesByHour

	// Create Wordlist from WordMap
	a.generateWordlist(r)

	return r, nil
}

// wordMap iterates over all words in a message, removes any stopwords, and populates the participants wordmap accordingly.
func wordMap(p *report.ParticipantAnalytics, message []byte) {
	// Remove punctuation marks from our message!
	m := bytes.Map(cleanMessage, message)
	words := strings.Fields(string(m))
	for _, word := range words {
		word := strings.ToLower(word)
		if _, ok := stopwords[word]; !ok {
			_, ok := p.WordMap[word]
			if !ok {
				//p.WordMap[word] = report.Word{Word: word, Usage: 1}
				p.WordMap[word] = 1
			} else {
				p.WordMap[word]++
			}
		}
	}
}

// generateWordList sorts the participant Wordmap and generates a WordList from the top 100 most used words
func (a *analyzer) generateWordlist(r *report.ChatAnalytics) {
	for k, p := range r.Participants {
		wl := make(report.Wordlist, len(p.WordMap))

		i := 0
		for k, v := range p.WordMap {
			wl[i] = report.Word{Word: k, Usage: v}
			i++
		}

		sort.Sort(wl)
		if len(wl) > 100 {
			p.WordList = wl[:100] // Get 100 most used words!
			break
		}
		p.WordList = wl[:len(wl)-1]


		// Create temp variable to assign to participant map
		temp := r.Participants[k]
		temp.WordList = p.WordList

		r.Participants[k] = temp
	}

}

// clean message is our mapping func to remove punctuation marks
func cleanMessage(r rune) rune {
	switch {
	case r == '\'' || r == ',' || r == '.' || r == '!' || r == '"':
		return -1
	}
	return r
}
