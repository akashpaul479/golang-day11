package day11

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User1 struct {
	ID   int
	Name string
}

var users = []User1{
	{ID: 1, Name: "Akash"},
	{ID: 2, Name: "Bunty"},
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func Adduser(w http.ResponseWriter, r *http.Request) {
	var u User1
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u.ID = len(users) + 1
	users = append(users, u)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintf(w, "Deleted user with id %d", id)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}
func Crudmethod1() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetUser(w, r)
		case http.MethodPost:
			Adduser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			DeleteUser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
