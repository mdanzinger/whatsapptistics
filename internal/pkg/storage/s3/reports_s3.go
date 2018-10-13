package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/mdanzinger/mywhatsapp/internal/pkg/report"
	"log"
)

type storage struct {
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func (s *storage) Store(ctx context.Context, report *report.Report) error {
	// Remove dashes
	//idProcessed := strings.Replace(id.String(), "-", "", -1)

	result, err := s.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String("whatsappchats"),
		Key:    aws.String(report.ReportID),
		Body:   bytes.NewReader(report.Content),
	})
	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	fmt.Printf("Sucessfully uploaded to %s \n", result.Location)
	return nil
}

func (s *storage) Get(ctx context.Context, key string) (*report.Report, error) {
	reportBuf := aws.NewWriteAtBuffer([]byte{})

	_, err := s.downloader.DownloadWithContext(ctx, reportBuf,
		&s3.GetObjectInput{
			Bucket: aws.String("whatsappchats"),
			Key:    aws.String(key),
		})

	if err != nil {
		fmt.Println("Unable to download item :  %v", err)
		return nil, err
	}

	// create reader
	reportReader := bytes.NewReader(reportBuf.Bytes())

	return report.NewReport(reportReader), nil

}

func NewReportStorage() *storage {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	u := s3manager.NewUploader(sess)
	d := s3manager.NewDownloader(sess)
	return &storage{
		uploader:   u,
		downloader: d,
	}
}