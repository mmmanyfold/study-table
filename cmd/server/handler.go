package server

import (
	"bytes"
	"encoding/json"
	"github.com/mmmanyfold/study-table-service/pkg/airtable"
	"github.com/mmmanyfold/study-table-service/pkg/aws"
	"log"
	"net/http"
)

func (app *AppConfig) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200 - OK"))
}

func (app *AppConfig) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - only GET is allowed"))
	}

	payload := airtable.GetAirtable()

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("err 1: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - failed to encode JSON payload"))
	}

	body := bytes.NewReader(jsonData)

	if err := aws.UploadFile(app.Sess, body); err != nil {
		log.Printf("err 2: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - failed to upload JSON to AWS S3"))
	}

	w.Write([]byte("file successfully uploaded to S3"))
}
