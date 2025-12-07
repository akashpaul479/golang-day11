package day11

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	fmt.Fprintf(w, "Search query: %s\n ", query)

	tags := r.URL.Query()["tags"]
	fmt.Fprintf(w, "Tags :%v\n", tags)
}
func Queryparameters1() {
	http.HandleFunc("/search", Handler)
	http.ListenAndServe(":8080", nil)
}
