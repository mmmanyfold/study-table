package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
)

const bucket = "study-table-service-assets"
const filename = "airtable.json"

func InitSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		Region:                        aws.String(endpoints.UsEast1RegionID)},
	)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func UploadFile(session *session.Session, body io.Reader) error {
	uploader := s3manager.NewUploader(session)

	if _, err := uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Body:        body,
		Bucket:      aws.String(bucket),
		ContentType: aws.String("application/json"),
		Key:         aws.String(filename),
	}); err != nil {
		return err
	}

	return nil
}
