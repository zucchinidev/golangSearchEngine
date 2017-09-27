package search

import (
	"os"
	"encoding/json"
	"path/filepath"
)

const dataFile = "../data/feeds.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func GetFeeds() ([]*Feed, error) {
	absolutePath, err := filepath.Abs(dataFile)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(absolutePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, nil
}
