package utils

import (
	"fmt"
)

type Index map[string][]int

func (idx Index) Add(docs []document) {
    for _, doc := range docs {
        for _, token := range Analyze(doc.Text) {
            ids := idx[token]
            fmt.Println(ids)
            if ids != nil && doc.ID == ids[len(ids) - 1] {
                continue
            }
            idx[token] = append(ids, doc.ID)
        }
    }
}