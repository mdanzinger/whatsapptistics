package store

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"

	"github.com/mdanzinger/whatsapp/internal/pkg/reports"
)

type ReportStore struct {
	db *dynamodb.DynamoDB
}

func (s *ReportStore) Create(r *reports.Report) (bool, error) {
	av, err := dynamodbattribute.MarshalMap(r)

	if err != nil {
		fmt.Errorf("Error creating marhsal map -> %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("mywhatsapp_reports"),
	}

	_, err = s.db.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Successfully added Report %s to Reports table \n", r.ReportID)

	return true, nil
}

func (s *ReportStore) Read(r int) (*reports.Report, error) {
	fmt.Printf("Getting report ID: %v ", r)
	return &reports.Report{}, nil
}

func (s *ReportStore) Update(r *reports.Report) (bool, error) {
	fmt.Printf("Updating report ID: %v ", r.ReportID)
	return true, nil
}

func (s *ReportStore) Delete(r *reports.Report) (bool, error) {
	fmt.Printf("Deleting report ID: %v ", r.ReportID)
	return true, nil
}

func NewReportStore() *ReportStore {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Errorf("Error creating session -> %s", err)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	return &ReportStore{
		db: svc,
	}
}
