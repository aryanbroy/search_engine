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
	query := flag.String("q", "Small wild cat", "search query")
	flag.Parse()

	startTime := time.Now()
    fmt.Println("Loading documents...")
	docs, err := utils.LoadDocuments(*dumpPath)
	if err != nil {
		log.Fatalf("Error loading documents: %v", err.Error())
	}

	fmt.Printf("Time to load %v documents: %v\n", len(docs), time.Since(startTime))

    startTime = time.Now()
    fmt.Println("Indexing documents...")
    idx := make(utils.Index)
    idx.Add(docs)
    fmt.Printf("Indexed %v docs in %v\n", len(docs), time.Since(startTime)) 

    startTime = time.Now()
    fmt.Println("Searching...")
    matchedIds := idx.Search(*query)
	fmt.Printf("Found %v document(s) that contains the query: %v, time took: %v\n", len(matchedIds), *query, time.Since(startTime))
}