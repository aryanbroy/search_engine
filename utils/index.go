package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/aryanbroy/search_engine/types"
)

type Index map[string][]int

func (idx Index) Add(docs []document) {
	for _, doc := range docs {
		for _, token := range Analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && doc.ID == ids[len(ids)-1] {
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func Intersection(a []int, b []int) []int {

	var i, j int
	maxLen := int(math.Max(float64(len(a)), float64(len(b))))

	r := make([]int, 0, maxLen)

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func (idx Index) Search(text string) []int {
	var r []int

	for _, token := range Analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = Intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}

func (idx Index) HandleSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	query := r.URL.Query().Get("q")

	searchTime := time.Now()

	log.Printf("Searching for `%v`", query)

	matchedIds := idx.Search(query)

	timeTook := fmt.Sprintf("%v μs", time.Since(searchTime).Microseconds())

	res := types.Response{
		Query:       query,
		MatchedDocs: matchedIds,
		TimeTook:    timeTook,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)
}
