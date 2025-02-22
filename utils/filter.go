package utils

import "strings"

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
