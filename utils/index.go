package utils

import "math"

type Index map[string][]int

func (idx Index) Add(docs []document) {
    for _, doc := range docs {
        for _, token := range Analyze(doc.Text) {
            ids := idx[token]
            if ids != nil && doc.ID == ids[len(ids) - 1] {
                continue
            }
            idx[token] = append(ids, doc.ID)
        }
    }
}

func Intersection(a []int, b []int) []int{

    var i, j int
    maxLen := int(math.Max(float64(len(a)), float64(len(b))))
    
    r := make([]int, maxLen)

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