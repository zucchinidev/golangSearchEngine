package main

import (
	"log"
	"os"
	_ "github.com/zucchinidev/golangSearchEngine/matchers"
	"github.com/zucchinidev/golangSearchEngine/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var searchTerm = "Angola"
	search.Run(searchTerm)
}
