package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/aryanbroy/search_engine/utils"
)

func main() {
	dumpPath := flag.String("p", "enwiki-latest-abstract1.xml.gz", "path to file")
	// query := flag.String("q", "Small wild cat", "search query")
	flag.Parse()

	startTime := time.Now()
	docs, err := utils.LoadDocuments(*dumpPath)
	if err != nil {
		log.Fatalf("Error loading documents: %v", err.Error())
	}

	fmt.Printf("Time to load %v documents: %v\n", len(docs), time.Since(startTime))
}
