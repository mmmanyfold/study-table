package server

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AppServer struct {
	Sess     *session.Session
	Uploader *s3manager.Uploader
}
