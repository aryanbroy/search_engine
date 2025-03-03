package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aryanbroy/search_engine/checkers"
	"github.com/aryanbroy/search_engine/utils"
)

func main() {
	dumpPath := flag.String("p", "enwiki-latest-abstract1.xml.gz", "path to file")
	// query := flag.String("q", "Small wild cat", "search query")
	flag.Parse()

	router := http.NewServeMux()

	startTime := time.Now()
	isIndexed := true
	fmt.Println("Loading documents...")
	docs, err := utils.LoadDocuments(*dumpPath, &isIndexed)
	if err != nil {
		log.Fatalf("Error loading documents: %v", err.Error())
	}

	log.Printf("Time to load %v documents: %v\n", len(docs), time.Since(startTime))

	idx := make(utils.Index)
	startTime = time.Now()

	log.Println("Indexing documents...")

	indexFile := "index.gob"
	if !checkers.FileExists(indexFile) {
		log.Println("Indexing for the first time...")
		idx.Add(docs)
		utils.SaveIndex(indexFile, idx)
	} else {
		log.Println("Index file already exists")
		utils.LoadIndex(indexFile, &idx)
	}

	log.Printf("Indexed %v docs in %v\n", len(docs), time.Since(startTime))

	router.HandleFunc("/search", idx.HandleSearch)

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server started at port: ", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("Failed to start server")
	}

	// fmt.Println("Searching...")
	// searchTime := time.Now()
	// matchedIds := idx.Search(*query)

	// fmt.Printf("Found %v document(s),i.e. %v, that contains the query: %v, time took: %v\n", len(matchedIds), matchedIds, *query, time.Since(searchTime))
}
