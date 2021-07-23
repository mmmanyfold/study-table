package airtable

type TagRecord struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ArtistRecord struct {
	CreatedTime string `json:"createdTime"`
	Fields      struct {
		Bio   string `json:"Bio"`
		Image []struct {
			Filename   string `json:"filename"`
			ID         string `json:"id"`
			Size       int64  `json:"size"`
			Thumbnails struct {
				Full struct {
					Height int64  `json:"height"`
					URL    string `json:"url"`
					Width  int64  `json:"width"`
				} `json:"full"`
				Large struct {
					Height int64  `json:"height"`
					URL    string `json:"url"`
					Width  int64  `json:"width"`
				} `json:"large"`
				Small struct {
					Height int64  `json:"height"`
					URL    string `json:"url"`
					Width  int64  `json:"width"`
				} `json:"small"`
			} `json:"thumbnails"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"Image"`
		Name    string   `json:"Name"`
		Publish bool     `json:"PUBLISH"`
		Tags    []string `json:"Tags"`
	} `json:"fields"`
	ID string `json:"id"`
}

type ArtistRecords struct {
	Records []ArtistRecord `json:"records"`
}

type ArtistAndTagsPayload struct {
	Tags    []TagRecord     `json:"tags"`
	Records []ArtistRecord `json:"records"`
}
