package utils

import (
	"compress/gzip"
	"encoding/xml"
	"log"
	"os"

	"github.com/aryanbroy/search_engine/checkers"
)

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocuments(path string, isIndexed *bool) ([]document, error) {
	cachePath := "example.gob"

	if checkers.FileExists(cachePath) {
		*isIndexed = true
		docs := struct {
			Document []document `xml:"doc"`
		}{}
		log.Println("Loading indexes from cache...")
		err := LoadDocs(cachePath, &docs.Document)
		if err != nil {
			log.Fatalf("Error loading indexes: %v", err.Error())
			return nil, err
		}
		log.Println("Successfully loaded docs from cache!")
		return docs.Document, nil
	}
	log.Println("Loading documents for the first time")
	log.Println("Opeing file...", path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err.Error())
		return nil, err
	}
	log.Println("Successfully opened file")
	defer file.Close()

	log.Println("Reading the zip file")

	gz, err := gzip.NewReader(file)
	if err != nil {
		log.Fatalf("Error reading zip file: %v", err.Error())
		return nil, err
	}

	log.Println("Reading zip file completed")
	defer gz.Close()

	log.Println("Decoding xml...")

	doc := xml.NewDecoder(gz)
	dump := struct {
		Document []document `xml:"doc"`
	}{}
	if err := doc.Decode(&dump); err != nil {
		log.Fatalf("Error decoding xml: %v", err.Error())
		return nil, err
	}
	log.Println("Sucessfully decoded xml file")

	docs := dump.Document

	for i := range docs {
		docs[i].ID = i
	}

	log.Println("No cache file found, caching now!!")
	SaveDocs(docs, cachePath)
	return docs, nil
}
