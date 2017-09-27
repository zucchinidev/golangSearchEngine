package search

type defaultMatcher struct {}

const DefaultMatcherTypeName = "default"

func init() {
	var matcher defaultMatcher
	Register(DefaultMatcherTypeName, matcher)
}

func (d defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
