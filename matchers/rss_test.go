package matchers

import (
	"github.com/zucchinidev/golangSearchEngine/search"
	"testing"
)

func TestRSSMatcher_Search(t *testing.T) {
	matcher := new(RSSMatcher)
	feed := search.Feed{
		Name: "npr",
		URI:  "http://www.npr.org/rss/rss.php?id=1001",
		Type: "rss",
	}
	results, err := matcher.Search(&feed, "Between Blacks")
	if err != nil {
		t.Error("should retrieve a match string but retrieve an error", err)
	}
	if len(results) == 0 {
		t.Error("should retrieve a match string but it doesn't return anything")
	}
}
