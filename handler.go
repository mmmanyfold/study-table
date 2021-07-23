package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/mmmanyfold/study-table-service/pkg/airtable"
	"log"
	"net/http"
)

const bucket = "study-table-service-assets"
const filename = "airtable.json"

type AppServer struct {
	awsSess *session.Session
	uploader *s3manager.Uploader
}

func (a *AppServer) HealthHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("200 - OK"))
}

func (a *AppServer) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - only GET is allowed"))
	}

	artists := airtable.GetRecords("Artists")
	tags := airtable.ExtractTags(artists)

	payload := airtable.ArtistAndTagsPayload{
		Tags:    tags,
		Records: artists,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - failed to encode JSON payload"))
	}

	body := bytes.NewReader(jsonData)

	if _, err := a.uploader.Upload(&s3manager.UploadInput{
		ACL:                       aws.String("public-read"),
		Body:                      body,
		Bucket:                    aws.String(bucket),
		ContentType:               aws.String("application/json"),
		Key:                       aws.String(filename),
	}); err != nil {
		log.Printf("err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - failed to upload JSON to AWS S3"))
	}

	w.Write([]byte("file successfully uploaded to S3"))
}
