package types

type Response struct {
	Query       string `json:"query"`
	MatchedDocs []int  `json:"matchedDocs"`
	TimeTook    string `json:"timeTook"`
}
