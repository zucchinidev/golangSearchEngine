package search

import (
	"os"
	"encoding/json"
	"path/filepath"
	"runtime"
	"path"
	"log"
)

const dataFile = "data/feeds.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func GetFeeds() ([]*Feed, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	absolutePath, err := filepath.Abs(path.Dir(filename) + "/../" + dataFile)
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
