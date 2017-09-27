package search

import "log"

type Result struct {
	Field   string
	Content string
}

func DisplayResults(results chan *Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
