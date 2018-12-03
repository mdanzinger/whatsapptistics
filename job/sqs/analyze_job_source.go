package sqs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mdanzinger/whatsapptistics/job"
	"github.com/mdanzinger/whatsapptistics/job/sns"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// TODO: Make these env variables
const (
	QUEUE_URL    = "https://sqs.us-east-2.amazonaws.com/582875565416/whatsapp-chats"
	MAX_MESSAGES = 10 // sqs has a 10 message limit.
)

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

type sqsSource struct {
	sqs       *sqs.SQS
	snsQueuer *sns.SnsQueuer // we used SNS to indirectly add items to queue.
	logger    *log.Logger

	currentBatch []job.Chat
}

func (s *sqsSource) NextJob() (j *job.Chat, err error) {
	if len(s.currentBatch) == 0 {
		// No messages in current batch, need to poll for more.
		// Once we have a set of messages, we store them in
		// s.currentBatch.
		for {
			fmt.Println("Polling SQS for messages...")
			resp, err := s.sqs.ReceiveMessage(params)
			if err != nil {
				log.Println(err)
			}

			// Loop through messages returned, add to batch
			if len(resp.Messages) > 0 {
				fmt.Printf("Got %v messages! \n", len(resp.Messages))
				for _, m := range resp.Messages {
					// Extract body from message
					mjson := message{}
					// Unmarshal body to get message
					err := json.Unmarshal([]byte(*m.Body), &mjson)
					if err != nil {
						return nil, err
					}
					// unmarshal message to get job
					cj := job.Chat{}
					err = json.Unmarshal([]byte(mjson.Message), &cj)
					if err != nil {
						return nil, err
					}
					s.currentBatch = append(s.currentBatch, cj)

					// Delete message from sqs
					s.deleteMessage(m.ReceiptHandle)
				}

				// Break out of our loop!
				break
			}
		}
	} else {
		fmt.Println("Found in local cache")
	}
	job := s.currentBatch[0]
	s.currentBatch = s.currentBatch[1:]
	return &job, nil
}

func (s *sqsSource) QueueJob(j *job.Chat) error {
	return s.snsQueuer.QueueJob(j)
}

// deleteMessage removes the message from the queue
func (s *sqsSource) deleteMessage(handle *string) error {
	_, err := s.sqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(QUEUE_URL),
		ReceiptHandle: handle,
	})
	if err != nil {
		return err
	}

	return nil
}

// NewJobSource returns an sqs implementation of job.Source
func NewJobSource(l *log.Logger) *sqsSource {
	sess, err := session.NewSession()
	if err != nil {
		log.Printf("Error creating session -> %s \n", err)
	}

	q := sqs.New(sess)

	return &sqsSource{
		sqs:       q,
		snsQueuer: sns.NewSnsQueuer(),
		logger:    l,
	}
}
