package analyzer

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type ChatAnalyzer struct {
	ChatID string `json:"Message"`
	Email  string `json:"email_address"`
	Body   io.Reader
}

// msmbody stores the values we want from the sqs message
type msgbody struct {
	MessageId         string
	Message           string
	MessageAttributes struct {
		EmailAddress struct {
			Value string
		}
	}
}

func NewChatAnalyzer(msg *sqs.Message) error {
	m := msgbody{}
	json.Unmarshal([]byte(*msg.Body), &m)

	fmt.Println(msg)
	//json.Unmarshal([]byte(*msg.Body), &a)

	a := &ChatAnalyzer{
		ChatID: m.Message,
		Email:  m.MessageAttributes.EmailAddress.Value,
	}

	log.Printf("Recieved SQS Message: %s", &m)
	//log.Printf("Recieved SQS Message: %s", a.Email)
	go a.Start()
	return nil
}

func (c *ChatAnalyzer) Start() {

	c.getChat()
	c.getWords()

	fmt.Println("Started Chat Analysis for email: ", c.Email)
}

// func getChat receives the chat from s3
func (c *ChatAnalyzer) getChat() {
	downloadclient := s3manager.NewDownloader(Sess)
	buff := &aws.WriteAtBuffer{}

	numBytes, err := downloadclient.Download(buff,
		&s3.GetObjectInput{
			Bucket: aws.String("whatsappchats"), // Hard coding for now.. TODO: add some sort of abstraction?
			Key:    aws.String(c.ChatID),
		})

	if err != nil {
		log.Fatalf("Unable to download item %q, %v", c.ChatID, err)
	}
	fmt.Println("Downloaded", numBytes, "bytes")

	c.Body = bytes.NewReader(buff.Bytes())
}

func (c *ChatAnalyzer) getWords() {
	wordMap := map[string]int{}

	scanner := bufio.NewScanner(c.Body)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		if len(word) <= 4 {
			continue
		}
		word = strings.Trim(word, `.,"'?!`)

		curCount := 0
		if v, ok := wordMap[word]; ok {
			curCount = v
		}

		wordMap[word] = 1 + curCount
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("failed to count words, %v", err)
	}

	words := Words{}
	for word, count := range wordMap {
		words = append(words, Word{Word: word, Count: count})
	}
	sort.Sort(words)

	fmt.Println(words[:100])

	//return wordMap, nil
}

type Word struct {
	Word  string
	Count int
}

type Words []Word

func (w Words) Len() int {
	return len(w)
}
func (w Words) Less(i, j int) bool {
	return w[i].Count > w[j].Count
}
func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
