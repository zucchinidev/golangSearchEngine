package main

import (
	_ "github.com/zucchinidev/golangSearchEngine/matchers"
	"github.com/zucchinidev/golangSearchEngine/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var searchTerm = "Angola"
	search.Run(searchTerm)
}
