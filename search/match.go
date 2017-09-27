package search

import "log"

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, resultChannel chan<- *Result) {
	foundResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range foundResults {
		resultChannel <- result
	}
}
