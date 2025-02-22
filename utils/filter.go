package utils

import (
	"log"
	"strings"

	"github.com/kljensen/snowball"
)

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, word := range tokens {
		r[i] = strings.ToLower(word)
	}
	return r
}

func stopWordsFilter(tokens []string) []string {
	stopWords := map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}

	r := make([]string, 0, len(tokens))

	for _, word := range tokens {
		if _, found := stopWords[word]; !found {
			r = append(r, word)
		}
	}
	return r
}

func stemFilter(tokens []string) []string {
    r := make([]string, len(tokens))
    for i, word := range tokens {
        stemWord, err := snowball.Stem(word, "english", true)
        if err != nil {
            log.Fatalf("Error stemming %v, error: %v", word, err.Error())
            return nil
        }
        r[i] = stemWord
    }
    return r
}