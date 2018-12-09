package dynamodb

import (
	"context"
	"fmt"

	_ "github.com/joho/godotenv/autoload"

	"github.com/mdanzinger/whatsapptistics/report"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// reportRepo is an implementation of a ReportRepository
type reportRepo struct {
	db    *dynamodb.DynamoDB
	cache report.ReportRepository
}

func (s *reportRepo) Get(ctx context.Context, key string) (*report.Report, error) {
	// Get report from cache
	r, err := s.cache.Get(ctx, key)

	// Cache hit
	if err == nil {
		fmt.Printf("Serving report %s from cache \n", key)
		return r, err
	}
	fmt.Println(err)

	// Cache misses.. get report from db
	result, err := s.db.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("AWS_DYNAMODB_REPORTS")),
		Key: map[string]*dynamodb.AttributeValue{
			"ReportID": {
				S: aws.String(key),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	// No errors, create report and populate
	report := &report.Report{}

	err = dynamodbattribute.UnmarshalMap(result.Item, report)

	if err != nil {
		return nil, err
	}

	// Store it in cache for future use..
	if err = s.cache.Store(report); err != nil {
		return report, err
	}

	fmt.Printf("Serving report %s from dynamodb \n", key)
	return report, nil
}

func (s *reportRepo) Store(r *report.Report) error {
	i, err := dynamodbattribute.MarshalMap(r)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      i,
		TableName: aws.String(os.Getenv("AWS_DYNAMODB_REPORTS")),
	}

	// Insert item into db
	_, err = s.db.PutItem(input)

	if err != nil {
		//fmt.Println("Got error calling PutItem:")
		//fmt.Println(err.Error())
		//os.Exit(1)
		return err
	}

	// Store report in cache for future use..
	if err = s.cache.Store(r); err != nil {
		return err
	}

	// TODO: Remove useless printf statement
	fmt.Printf("Successfully added Report %s to Reports table \n", r.ReportID)

	// No errors!
	return nil
}

// NewReportRepo returns a report repository
func NewReportRepo(cache report.ReportRepository) *reportRepo {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("Error creating session -> %s", err)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return &reportRepo{
		db:    svc,
		cache: cache,
	}
}
