package day11

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{"Message": "Hello json"}
	json.NewEncoder(w).Encode(response)
}

func SendJsons() {
	http.HandleFunc("/json:", SendJson)
	fmt.Println("Server runnimg on http://Localhost:8080/json")
	http.ListenAndServe(":8080", nil)
}
