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

func GetRecords(table string) []ArtistRecord {
	baseURL := fmt.Sprintf("https://api.airtable.com/v0/appgnbNAyXRTziPYF/%s", table)

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

	var artistsJSON ArtistRecords

	err = json.Unmarshal([]byte(body), &artistsJSON)
	if err != nil {
		log.Fatal(err)
	}

	return artistsJSON.Records
}

func FilterDeletedAndPublishedArtists(records []ArtistRecord) []ArtistRecord {
	var filtered []ArtistRecord

	for _, r := range records {
		if r.Fields.Name != "" && !r.Fields.Delete && r.Fields.Publish == true {
			filtered = append(filtered, r)
		}
	}

	return filtered
}

func ExtractTags(artists []ArtistRecord) []TagRecord {
	var tags []TagRecord
	var tagCount = 1

	for _, r := range artists {
		if len(r.Fields.Tags) > 0 {
			for _, t := range r.Fields.Tags {
				if !FilterTag(tags, t) {
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

func FilterTag(tags []TagRecord, tag string) bool {
	for _, t := range tags {
		if t.Name == tag {
			return true
		}
	}

	return false
}

func ScheduleAirtableSync(sess *session.Session) {
	everyHour := time.NewTicker(1 * time.Hour)

	for {
		select {
		case <-everyHour.C:
			log.Println("syncing airtable records")

			artists := GetRecords("Artists")
			artists = FilterDeletedAndPublishedArtists(artists)
			tags := ExtractTags(artists)
			now := time.Now()

			payload := ArtistAndTagsPayload{
				Meta: Meta{
					LastUpdateAt: now.String(),
					Version:      os.Getenv("COMMIT"),
				},
				Tags:    tags,
				Records: artists,
			}

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
