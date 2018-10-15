package sqs

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/mdanzinger/mywhatsapp/internal/pkg/report"
	"log"
)

type reportPoller struct {
	SQS      *sqs.SQS
	QueueUrl string
}

// Poll satisifies the ReportPoller interface by polling sqs for messages and returning a slice of reports from their ids
func (rp *reportPoller) Poll() ([]report.Report, error) {
	fmt.Println("Polling for messages...")
	var reports []report.Report
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(rp.QueueUrl), // Required
		MaxNumberOfMessages: aws.Int64(10),
		MessageAttributeNames: []*string{
			aws.String("All"), // Required
		},
		WaitTimeSeconds: aws.Int64(20),
	}

	resp, err := rp.SQS.ReceiveMessage(params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(resp.Messages) > 0 {
		reports = parse(resp.Messages)
	}
	return reports, nil
}

type msgbody struct {
	MessageId         string
	Message           string
	MessageAttributes struct {
		EmailAddress struct {
			Value string
		}
	}
}

func parse(msgs []*sqs.Message) []report.Report {
	var reports []report.Report

	for _, m := range msgs {
		mb := msgbody{}
		if err := json.Unmarshal([]byte(*m.Body), &mb); err != nil {
			fmt.Errorf("parse Amazon S3 Event message %v", err)
		}
		reports = append(reports, report.Report{
			ReportID: mb.Message,
			Email:    mb.MessageAttributes.EmailAddress.Value,
		})
	}
	return reports
}

// NewReportPoller returns a new sqs reportPoller that can be polled to receive new reports.
func NewReportPoller() *reportPoller {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Errorf("Error creating session -> %s", err)
	}
	s := sqs.New(sess)

	resultURL, err := s.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("whatsapp-chats"),
	})

	return &reportPoller{
		SQS:      s,
		QueueUrl: aws.StringValue(resultURL.QueueUrl),
	}
}
