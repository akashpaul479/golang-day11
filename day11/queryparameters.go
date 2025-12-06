package day11

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Welcome, %s!", name)
}
func Queryparameter() {
	http.HandleFunc("Welcome:", Welcome)
	http.ListenAndServe(":8080", nil)
}
