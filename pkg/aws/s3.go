package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
)

func InitSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID)},
	)
	if err != nil {
		return nil, err
	}

	return sess, nil
}
