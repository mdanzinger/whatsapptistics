package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mdanzinger/mywhatsapp/old/internalrnal/pkg/report"
	"os"
)

type reportStore struct {
	db    *dynamodb.DynamoDB
	cache report.ReportStore
}

func (s *reportStore) Get(ctx context.Context, key string) (*report.Report, error) {
	r, err := s.cache.Get(ctx, key)

	if err == nil {
		return r, err
	}

	result, err := s.db.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("mywhatsapp_reports"),
		Key: map[string]*dynamodb.AttributeValue{
			"ReportID": {
				S: aws.String(key),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	report := &report.Report{}

	err = dynamodbattribute.UnmarshalMap(result.Item, report)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if err = s.cache.Store(ctx, report); err != nil {
		return report, err
	}

	return report, nil
}

func (s *reportStore) Store(ctx context.Context, r *report.Report) error {
	av, err := dynamodbattribute.MarshalMap(r)

	if err != nil {
		fmt.Errorf("Error creating marhsal map -> %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("mywhatsapp_reports"),
	}

	_, err = s.db.PutItemWithContext(ctx, input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
		return err
	}

	if err = s.cache.Store(ctx, r); err != nil {
		fmt.Printf("Failed to add report to cache")
		return err
	}

	fmt.Printf("Successfully added Report %s to Reports table \n", r.ReportID)

	return nil
}
func NewReportStore(cache report.ReportStore) *reportStore {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Errorf("Error creating session -> %s", err)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return &reportStore{
		db:    svc,
		cache: cache,
	}
}
