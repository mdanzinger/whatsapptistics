package sns

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/mdanzinger/mywhatsapp/internal/pkg/report"
)

type Notifier struct {
	sns *sns.SNS
}

func (n *Notifier) Notify(ctx context.Context, r *report.Report) error {
	params := &sns.PublishInput{
		Message:  aws.String(r.ReportID),
		TopicArn: aws.String("arn:aws:sns:us-east-2:582875565416:whatsapp_chats"),
	}

	if r.Email != " " {
		params.MessageAttributes = map[string]*sns.MessageAttributeValue{
			"EmailAddress": {
				DataType:    aws.String("String"), // Required
				StringValue: aws.String(r.Email),
			},
		}
	}

	fmt.Println(params)
	//
	resp, err := n.sns.Publish(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(awsutil.StringValue(resp))
	//return nil
	return nil
}

func NewReportNotifier() *Notifier {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Errorf("Error creating session -> %s", err)
	}
	sns := sns.New(sess)

	return &Notifier{
		sns: sns,
	}
}
