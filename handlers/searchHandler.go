package handlers

import (
	"net/http"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello there"))
}
