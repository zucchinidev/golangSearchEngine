package search

import "log"

var availableMatchers = make(map[string]Matcher)

func Register(matcherType string, matcher Matcher)  {
	if _, exists := availableMatchers[matcherType]; exists {
		log.Fatalf(matcherType, "Matcher already registered")
	}

	log.Println("Register", matcherType, "matcher")
	availableMatchers[matcherType] = matcher
}

func GetAvailableMatchers() map[string]Matcher {
	return availableMatchers
}