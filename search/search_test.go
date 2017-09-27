package search

import (
	"testing"
)

type fakeMatcher struct{}

func (f fakeMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

func TestGetAvailableMatchers(t *testing.T) {
	hasMatcher(DefaultMatcherTypeName, t)
}

func TestRegister(t *testing.T) {
	var fake fakeMatcher
	matcherType := "fakeMatcher"
	Register(matcherType, fake)
	hasMatcher(matcherType, t)
}

func hasMatcher(matcherType string, t *testing.T) {
	availableMatchers := GetAvailableMatchers()
	_, exists := availableMatchers[matcherType]
	if !exists {
		t.Error(
			"For",
			matcherType,
			"expected a valid matcher but it don't exists")
	}
}
