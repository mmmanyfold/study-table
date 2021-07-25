package airtable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		if r.Fields.Name != "" && r.Fields.Info != "" && !r.Fields.Delete {
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
