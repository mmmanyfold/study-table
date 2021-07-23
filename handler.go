package main

import (
	"fmt"
	"net/http"

	"github.com/mmmanyfold/study-table-service/pkg/airtable"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200 - OK"))
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - only GET is allowed"))
	}

	artists := airtable.GetRecords("Artists")
	tags := airtable.ExtractTags(artists)

	a := airtable.ArtistAndTagsPayload{
		Tags:    tags,
		Records: artists,
	}

	fmt.Printf("%+v", a)
	// 4. store in S3

}
