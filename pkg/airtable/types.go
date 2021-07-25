package airtable

type TagRecord struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ArtistRecord struct {
	CreatedTime string `json:"createdTime"`
	Fields      struct {
		Info  string `json:"Info"`
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

type Meta struct {
	LastUpdateAt string `json:"last_updated_at"`
	Version      string `json:"version"`
}

type ArtistRecords struct {
	Records []ArtistRecord `json:"records"`
}

type ArtistAndTagsPayload struct {
	Meta    Meta           `json:"meta"`
	Tags    []TagRecord    `json:"tags"`
	Records []ArtistRecord `json:"records"`
}
