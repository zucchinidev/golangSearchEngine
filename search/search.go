package search

import (
	"log"
	"sync"
)

var availableMatchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := GetFeeds()
	if err != nil {
		log.Fatal(err)
	}

	resultChannel := make(chan *Result)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := availableMatchers[feed.Type]
		if !exists {
			matcher = availableMatchers[DefaultMatcherTypeName]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, resultChannel)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		close(resultChannel)
	}()

	DisplayResults(resultChannel)
}

func Register(matcherType string, matcher Matcher) {
	if _, exists := availableMatchers[matcherType]; exists {
		log.Fatalf(matcherType, "Matcher already registered")
	}

	log.Println("Register", matcherType, "matcher")
	availableMatchers[matcherType] = matcher
}

func GetAvailableMatchers() map[string]Matcher {
	return availableMatchers
}
