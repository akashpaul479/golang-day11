package day11

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello user!")
}
func Routesserver() {
	http.HandleFunc("/Hello:", Hello)
	http.ListenAndServe(":8080", nil)
}
