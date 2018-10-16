package worker

import (
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// ChatHandler interface
type ChatHandler interface {
	HandleChat(msg *sqs.Message) error
}

type ChatReceiver struct {
	AWSSession *session.Session
	JobSQS     *sqs.SQS
	JobSQSURL  string
}

// InvalidMessageError for message that can't be processed and should be deleted
type InvalidMessageError struct {
	SQSMessage string
	LogMessage string
}

type HandlerFunc func(msg *sqs.Message) error

func (e InvalidMessageError) Error() string {
	return fmt.Sprintf("[Invalid Message: %s] %s", e.SQSMessage, e.LogMessage)
}

// NewInvalidMessageError to create new error for messages that should be deleted
func NewInvalidMessageError(SQSMessage, logMessage string) InvalidMessageError {
	return InvalidMessageError{SQSMessage: SQSMessage, LogMessage: logMessage}
}

func (f HandlerFunc) HandleChat(msg *sqs.Message) error {
	return f(msg)
}

//
var (
	// MaxNumberOfMessage at one poll
	MaxNumberOfMessage int64 = 10
	// WaitTimeSecond for each poll
	WaitTimeSecond int64 = 20
)

// NewChatReceiver creates a worker service that retrieves chat from a queue
func NewChatReceiver(n string) (*ChatReceiver, error) {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	s := sqs.New(sess)

	resultURL, err := s.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(n),
	})

	if err != nil {
		log.Println("Can't get the SQS queue")
		return nil, err
	}

	builder := &ChatReceiver{
		AWSSession: sess,
		JobSQS:     s,
		JobSQSURL:  aws.StringValue(resultURL.QueueUrl),
	}

	return builder, nil
}

func (c *ChatReceiver) Start(h ChatHandler) {
	for {
		params := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(c.JobSQSURL), // Required
			MaxNumberOfMessages: aws.Int64(MaxNumberOfMessage),
			MessageAttributeNames: []*string{
				aws.String("All"), // Required
			},
			WaitTimeSeconds: aws.Int64(WaitTimeSecond),
		}

		resp, err := c.JobSQS.ReceiveMessage(params)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(resp.Messages) > 0 {
			run(c, h, resp.Messages)
		}
	}
}

func run(c *ChatReceiver, h ChatHandler, messages []*sqs.Message) {
	numMessages := len(messages)
	var wg sync.WaitGroup
	wg.Add(numMessages)
	for i := range messages {
		go func(m *sqs.Message) {
			// launch goroutine
			defer wg.Done()
			if err := handleMessage(c, m, h); err != nil {
				log.Println(err.Error())
			}
		}(messages[i])
	}

	wg.Wait()
}

func handleMessage(c *ChatReceiver, m *sqs.Message, h ChatHandler) error {
	err := h.HandleChat(m)
	if _, ok := err.(InvalidMessageError); ok {
		// Invalid message encountered. Swallow the error and delete the message
		log.Println(err.Error())
	} else if err != nil {
		// Message is valid but there is an error proccesing it. Keeping it in the
		// queue or send to DLQ to try again
		return err
	}

	// Delete the processed (or invalid) message
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(c.JobSQSURL), // Required
		ReceiptHandle: m.ReceiptHandle,         // Required
	}
	_, err = c.JobSQS.DeleteMessage(params)
	if err != nil {
		return err
	}

	return nil
}
