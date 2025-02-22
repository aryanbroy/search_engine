package utils

import (
	"compress/gzip"
	"encoding/xml"
	"log"
	"os"
)

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocuments(path string) ([]document, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err.Error())
		return nil, err
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		log.Fatalf("Error reading zip file: %v", err.Error())
		return nil, err
	}
	defer gz.Close()

	doc := xml.NewDecoder(gz)
	dump := struct {
		Document []document `xml:"doc"`
	}{}
	if err := doc.Decode(&dump); err != nil {
		log.Fatalf("Error decoding xml: %v", err.Error())
		return nil, err
	}

	docs := dump.Document

	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
