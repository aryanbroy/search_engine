package utils

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func SaveDocs(docs []document, fileName string) {
	log.Println("Saving docs...")
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("Error creating file: ", err.Error())
		return
	}

	defer file.Close()

	enc := gob.NewEncoder(file)
	if err := enc.Encode(docs); err != nil {
		log.Fatalln("Error encoding docs: ", err.Error())
		return
	}

	log.Println("Docs saved to a file")
}

func LoadDocs(fileName string, docs *[]document) error {
	log.Println("Opening file...")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Error opening file: ", err.Error())
		return fmt.Errorf("Error opening file: %v", err.Error())
	}

	log.Println("Finished opening the file")

	defer file.Close()

	log.Println("Decoding file...")
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&docs); err != nil {
		log.Fatalln("Error decoding file: ", err.Error())
		return fmt.Errorf("Error decoding file: %v", err.Error())
	}
	log.Println("Successfully decoded the file")
	return nil
}
