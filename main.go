package main

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/mmmanyfold/study-table-service/pkg/aws"
	"log"
	"net/http"
)

func main() {
	awsSess, err := aws.InitSession()
	if err != nil {
		log.Panic("failed to start aws session")
	}

	uploader := s3manager.NewUploader(awsSess)

	appServer := AppServer{
		awsSess:  awsSess,
		uploader: uploader,
	}

	http.HandleFunc("/", appServer.HealthHandler)
	http.HandleFunc("/webhook", appServer.WebhookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
