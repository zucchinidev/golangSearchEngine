package search

import (
	"testing"
)

func TestGetFeeds(t *testing.T) {
	feeds, err := GetFeeds()
	if err != nil {
		t.Error("Should retrieve the feeds but throuth an Error")
	}
	if len(feeds) == 0 {
		t.Error("For len(feeds)",
			"a value greater than 0 is expected but find", len(feeds))
	}
}
