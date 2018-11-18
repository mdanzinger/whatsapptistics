package s3

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/mdanzinger/whatsapptistics/internal/chat"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type chatRepo struct {
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func (s *chatRepo) Upload(ctx context.Context, chat *chat.Chat) error {
	result, err := s.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String("whatsappchats"),
		Key:    aws.String(chat.ChatID),
		Body:   bytes.NewReader(chat.Content),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Successfully uploaded to %s \n", result.Location)
	return nil
}

func (s *chatRepo) Download(id string) (*chat.Chat, error) {
	reportBuf := aws.NewWriteAtBuffer([]byte{})

	_, err := s.downloader.Download(reportBuf,
		&s3.GetObjectInput{
			Bucket: aws.String("whatsappchats"),
			Key:    aws.String(id),
		})

	if err != nil {
		return nil, err
	}

	c := &chat.Chat{
		ChatID:  id,
		Content: reportBuf.Bytes(),
	}

	return c, nil

}

// NewChatRepo returns a chat repository
func NewChatRepo() *chatRepo {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	u := s3manager.NewUploader(sess)
	d := s3manager.NewDownloader(sess)
	return &chatRepo{
		uploader:   u,
		downloader: d,
	}
}
