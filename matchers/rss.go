package matchers

import (
	"github.com/zucchinidev/golangSearchEngine/search"
	"log"
	"errors"
	"encoding/xml"
	"net/http"
	"fmt"
	"regexp"
	"reflect"
)

const RssMatcherTypeName = "rss"
var searchFields = [2]string{"Title", "Description"}

type RSSMatcher struct{}

type (
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}

	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}
)

func init() {
	var matcher RSSMatcher
	search.Register(RssMatcherTypeName, matcher)
}

func (rssMatcher RSSMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)
	document, err := rssMatcher.getFeed(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		for _, fieldName := range searchFields {
			result, err := findTermInField(searchTerm, &channelItem, fieldName)
			if err != nil {
				return nil, err
			}
			if result != nil {
				results = append(results, result)
			}
		}
	}

	return results, nil
}

func findTermInField(searchTerm string, item *item, fieldName string) (*search.Result, error) {
	reflectItem := reflect.ValueOf(item)
	fieldValue := reflect.Indirect(reflectItem).FieldByName(fieldName)
	matched, err := regexp.MatchString(searchTerm, fieldValue.String())
	if err != nil {
		return nil, err
	}

	if matched {
		return &search.Result{
			Field:   fieldName,
			Content: fieldValue.String(),
		}, nil
	}
	return nil, nil
}

func (rssMatcher RSSMatcher) getFeed(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("no rss feed uri provider")
	}
	response, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", response.StatusCode)
	}

	var document rssDocument
	err = xml.NewDecoder(response.Body).Decode(&document)
	return &document, nil
}
