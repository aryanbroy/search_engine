package utils

import (
	"encoding/gob"
	"log"
	"os"
)

func SaveIndex(filename string, idx Index) {
	log.Println("Saving indexes...")

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln("Error creating file: ", err.Error())
		return
	}

	defer file.Close()

	enc := gob.NewEncoder(file)
	if err := enc.Encode(idx); err != nil {
		log.Fatalln("Error encoding file: ", err.Error())
		return
	}
	log.Println("Indexes saved successfully!")
}

func LoadIndex(filename string, idx *Index) {
	log.Println("Loading indexes...")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Error opening file: ", err.Error())
		return
	}

	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&idx); err != nil {
		log.Fatalln("Error decoding file: ", err.Error())
		return
	}

	log.Println("Loaded indexes from an existing file")

}
