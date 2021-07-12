package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetRecords(table string) ArtistRecords {
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

	return artistsJSON
}

func main() {
	GetRecords("Artists")
}
