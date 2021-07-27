package airtable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/mmmanyfold/study-table-service/pkg/aws"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func GetRecords(airtable Response) Response {
	baseURL := "https://api.airtable.com/v0/appgnbNAyXRTziPYF/Artists"

	if airtable.Offset != "" {
		baseURL = fmt.Sprintf("%s?offset=%s", baseURL, airtable.Offset)
	}

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("AIRTABLE_API_KEY")))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var airtableJSON Response

	if err = json.Unmarshal([]byte(body), &airtableJSON); err != nil {
		log.Fatal(err)
	}

	if airtableJSON.Offset != "" {
		airtable.Offset = airtableJSON.Offset
		airtable.Records = append(airtable.Records, airtableJSON.Records...)
		return GetRecords(airtable)
	}

	airtable.Records = append(airtable.Records, airtableJSON.Records...)

	return airtable
}

func filterDeletedAndPublishedArtists(records []ArtistRecord) []ArtistRecord {
	var filtered []ArtistRecord

	for _, r := range records {
		if r.Fields.Name != "" && !r.Fields.Delete && r.Fields.Publish == true {
			filtered = append(filtered, r)
		}
	}

	return filtered
}

func GetAirtable() ArtistAndTagsPayload {
	var response Response

	airtable := GetRecords(response)
	filtered := filterDeletedAndPublishedArtists(airtable.Records)
	tags := extractTags(filtered)
	now := time.Now()

	return ArtistAndTagsPayload{
		Meta: Meta{
			LastUpdateAt: now.String(),
			Version:      os.Getenv("COMMIT"),
		},
		Tags:    tags,
		Records: filtered,
	}
}

func ScheduleAirtableSync(sess *session.Session) {
	every30Minutes := time.NewTicker(30 * time.Minute)

	for {
		select {
		case <-every30Minutes.C:
			log.Println("syncing airtable records")

			payload := GetAirtable()

			jsonData, err := json.Marshal(payload)
			if err != nil {
				log.Printf("err: %v", err)
				return
			}

			body := bytes.NewReader(jsonData)

			if err := aws.UploadFile(sess, body); err != nil {
				log.Printf("err: %v", err)
				return
			}

			log.Println("file successfully uploaded to S3")
		}
	}
}

func filterTag(tags []TagRecord, tag string) bool {
	for _, t := range tags {
		if t.Name == tag {
			return true
		}
	}

	return false
}

func extractTags(artists []ArtistRecord) []TagRecord {
	var tags []TagRecord
	var tagCount = 1

	for _, r := range artists {
		if len(r.Fields.Tags) > 0 {
			for _, t := range r.Fields.Tags {
				if !filterTag(tags, t) {
					tags = append(tags, TagRecord{
						Id:   tagCount,
						Name: t,
					})
					tagCount += 1
				}
			}
		}
	}

	return tags
}
