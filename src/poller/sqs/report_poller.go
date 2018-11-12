package sqs

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// TODO: Make these env variables
const (
	QUEUE_URL    = "https://sqs.us-east-2.amazonaws.com/582875565416/whatsapp-chats"
	MAX_MESSAGES = 10 // sqs has a 10 message limit.
)

type Message struct {
	Message string `json:"Message"`
}

type Tag struct {
	Records []struct {
		S3 struct {
			Object struct {
				Etag string `json:"etag"`
			}
		}
	}
}

var (
	params = &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(QUEUE_URL), // Required
		MaxNumberOfMessages: aws.Int64(MAX_MESSAGES),
		MessageAttributeNames: []*string{
			aws.String("All"), // Required
		},
		WaitTimeSeconds: aws.Int64(20), // 20 = long polling
	}
)

type reportPoller struct {
	sqs    *sqs.SQS
	logger *log.Logger
}

func (rp *reportPoller) Poll(c chan []string, wg *sync.WaitGroup) {
	mb := Message{}
	e := Tag{}
	for {
		resp, err := rp.sqs.ReceiveMessage(params)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(resp.Messages) > 0 {
			wg.Add(len(resp.Messages))
			//c <- []string{"chat 1", "chat 2", "chat 3", "chat 4", "chat 5", "chat 6", "chat 7", "chat 8", "chat 9", "chat 10"}
			for _, m := range resp.Messages {
				fmt.Print("RECEIVED MESSAGE:")
				fmt.Println(*m.Body)
				fmt.Println(m.Attributes)

				if err := json.Unmarshal([]byte(*m.Body), &mb); err != nil {
					fmt.Println(err)
				}
				if err := json.Unmarshal([]byte(mb.Message), &e); err != nil {
					fmt.Println(err)
				}
				fmt.Println(e)
			}
			wg.Wait()
		}
	}
}

// NewReportPoller returns an SQS implementation of a report poller
func NewReportPoller(l *log.Logger) *reportPoller {
	sess, err := session.NewSession()
	if err != nil {
		log.Printf("Error creating session -> %s \n", err)
	}

	q := sqs.New(sess)

	return &reportPoller{
		sqs:    q,
		logger: l,
	}
}
