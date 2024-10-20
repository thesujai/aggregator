package scrapper

import "testing"

func TestFetchFeed(t *testing.T) {
	_, err := fetchFeed("https://blog.boot.dev/index.xml")
	if err != nil {
		t.Fail()
	}
}
