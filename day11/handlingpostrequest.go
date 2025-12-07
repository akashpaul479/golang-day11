package day11

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name string
	Age  int
}

func Createuser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post allowed", http.StatusMethodNotAllowed)
		return
	}
	var u user
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
func Handlingpostrequest() {
	http.HandleFunc("/createuser", Createuser)
	fmt.Println("Server on http://localhost:8080/Createuser")
	http.ListenAndServe(":8080", nil)
}
