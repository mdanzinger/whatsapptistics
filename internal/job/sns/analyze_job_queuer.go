package sns

import (
	"encoding/json"
	"fmt"

	"github.com/mdanzinger/whatsapptistics/internal/job"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	TopicARN = "arn:aws:sns:us-east-2:582875565416:whatsapp_chats"
)

type SnsQueuer struct {
	sns *sns.SNS
}

func (s *SnsQueuer) QueueJob(job *job.Chat) error {
	j, err := json.Marshal(job)
	if err != nil {
		return fmt.Errorf("error marshalling job: %v", err)
	}

	params := &sns.PublishInput{
		Message:  aws.String(string(j)),
		TopicArn: aws.String(TopicARN),
	}

	fmt.Println("Params: ", params)
	//
	resp, err := s.sns.Publish(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(resp)
	return nil
}

func NewSnsQueuer() *SnsQueuer {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Errorf("Error creating session -> %s", err)
	}
	sns := sns.New(sess)

	return &SnsQueuer{
		sns: sns,
	}
}
