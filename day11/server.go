package day11

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello from go web server!")
	})

	port := ":5000"
	fmt.Println("server is runnimg on port" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
