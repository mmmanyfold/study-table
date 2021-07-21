package main

import (
	"net/http"

	"github.com/mmmanyfold/study-table-service/pkg/airtable"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200 - OK"))
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - only GET is allowed"))
	}

	artistsRecords := airtable.GetRecords("Artists")
	airtable.ExtractTags(artistsRecords)

	// 1. down the data
	// 2. extract tags from payload
	// 3. format data
	//	a. {
	//	     tags: [{id: 1, name: "xyz"}],
	//       artists: [{
	//                  id: 1,
	// 					name: 'Simone Forti',
	// 					image: (get the first image)
	// 					  'https://dl.airtable.com/.attachmentThumbnails/0936f073c25da2b872272632d75e696c/911b163c',
	// 					tags: ['Performance', 'Sculpture'],
	//                 }]
	//     }
	// 4. store in S3

}
