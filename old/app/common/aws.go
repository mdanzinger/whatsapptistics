package common

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"log"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"

	//"github.com/aws/aws-sdk-go/aws"
	//"bytes"
	//"net/http"

	"mime/multipart"
	"os"

	"fmt"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/nu7hatch/gouuid"
)

var (
	AWS_SESS      *session.Session
	AWS_S3MANAGER *s3manager.Uploader
	S3_BUCKET     = os.Getenv("S3_BUCKET")
	SNS_ARN       = os.Getenv("SNS_ARN")
)

// Init initializes the AWS session
func Init() {
	s, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	AWS_SESS = s
	svc := s3manager.NewUploader(s)
	AWS_S3MANAGER = svc

}

func AddFileToS3(s *session.Session, file multipart.File, fileheader *multipart.FileHeader) (key string, err error) {

	// Get file size and read the file content into a buffer
	var size int64 = fileheader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	// Generate uuid
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.

	result, err := AWS_S3MANAGER.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(id.String()),
		//Body:   bytes.NewReader(buffer),
		Body: file,
	})
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Printf("Sucessfully uploaded %s to %s \n", fileheader.Filename, result.Location)
	//_, err = s3.New(s).PutObject(&s3.PutObjectInput{
	//	Bucket:               aws.String(S3_BUCKET),
	//	Key:                  aws.String(id.String()),
	//	ACL:                  aws.String("private"),
	//	Body:                 bytes.NewReader(buffer),
	//	ContentLength:        aws.Int64(size),
	//	ContentType:          aws.String(http.DetectContentType(buffer)),
	//	ContentDisposition:   aws.String("attachment"),
	//	ServerSideEncryption: aws.String("AES256"),
	//})
	return id.String(), err
}

func SendSNSMessage(s *session.Session, key string, email string) (err error) {
	svc := sns.New(s)
	params := &sns.PublishInput{
		Message:  aws.String(key),
		TopicArn: aws.String(SNS_ARN),
	}

	if len(email) > 0 {
		params.MessageAttributes = map[string]*sns.MessageAttributeValue{
			"EmailAddress": {
				DataType:    aws.String("String"), // Required
				StringValue: aws.String(email),
			},
		}
	}

	fmt.Println(svc)
	fmt.Println(params)
	//
	resp, err := svc.Publish(params)
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

func DownloadFile(s *session.Session, key *string) {
	file, err := os.Create(*key)
	if err != nil {
		fmt.Println("Unable to open file %q, %v", err)
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(s)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(S3_BUCKET),
			Key:    key,
		})

	if err != nil {
		fmt.Println("Unable to download item %q, %v", file, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}
