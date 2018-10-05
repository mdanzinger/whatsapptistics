package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mdanzinger/whatsapp/internal/pkg/cache"
)

type reportStore struct {
	db    *dynamodb.DynamoDB
	cache cache.Cacher
}

func (s *reportStore)

//func (s *ReportStore) Create(r *report.Report) (bool, error) {
//	av, err := dynamodbattribute.MarshalMap(r)
//
//	if err != nil {
//		fmt.Errorf("Error creating marhsal map -> %s", err)
//	}
//
//	input := &dynamodb.PutItemInput{
//		Item:      av,
//		TableName: aws.String("mywhatsapp_reports"),
//	}
//
//	_, err = s.db.PutItem(input)
//
//	if err != nil {
//		fmt.Println("Got error calling PutItem:")
//		fmt.Println(err.Error())
//		os.Exit(1)
//	}
//
//	fmt.Printf("Successfully added Report %s to Reports table \n", r.ReportID)
//
//	return true, nil
//}
//
//func (s *ReportStore) Read(r int) (*report.Report, error) {
//	r, err := s.Cache.get(r)
//	return &report.Report{}, nil
//}
//
//func (s *ReportStore) Update(r *report.Report) (bool, error) {
//	fmt.Printf("Updating report ID: %v ", r.ReportID)
//	return true, nil
//}
//
//func (s *ReportStore) Delete(r *report.Report) (bool, error) {
//	fmt.Printf("Deleting report ID: %v ", r.ReportID)
//	return true, nil
//}
//
func NewReportStore(sess *session.Session, cache cache.Cacher) *reportStore {
	//sess, err := session.NewSession()
	//if err != nil {
	//	fmt.Errorf("Error creating session -> %s", err)
	//}

	// Create DynamoDB client
	svc := dynamodb.New(sess)



	return &reportStore{
		db: svc,
		cache: cache,
	}
}
