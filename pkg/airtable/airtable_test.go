package airtable

import (
	"net/http"
	"testing"
)

// Custom type that allows setting the func that our Mock Do func will run instead
type MockDoType func(req *http.Request) (*http.Response, error)

type MockClient struct {
	MockDo MockDoType
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestGetRecords(t *testing.T) {
	//jsonResponse := `
	//[{
	//  "fields": {
	//    "Name": "Gedi Sidony"
	//  }
	//},
	//{
	//  "fields": {
	//    "Name": "Carl Craig"
	//  }
	//}]
	//`

	//r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	//
	//client = &MockClient{
	//	MockDo: func(req *http.Request) (*http.Response, error) {
	//		return &http.Response{
	//			StatusCode: 200,
	//			Body: r,
	//		}, nil
	//	},
	//}

	//records := GetRecords("Artists")
	//ExtractTags(records)
	//fmt.Println("records", records)
	//t.Errorf("no records received")
	//if len(records.Records) != 0 {
	//}
}

func TestFilterTag(t *testing.T) {
	testTags := []TagRecord{{
		Id:   1,
		Name: "abc",
	}, {
		Id:   2,
		Name: "1970s",
	}}

	validTags := []string{"abc", "1970s"}
	for _, tagName := range validTags {
		if !filterTag(testTags, tagName) {
			t.Errorf("tag: %s, should be included", tagName)
		}

	}
}
