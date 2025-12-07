package project

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User2 struct {
	ID   int
	Name string
	Age  int
}

var users = []User2{
	{ID: 1, Name: "Akash", Age: 20},
	{ID: 2, Name: "Bunty", Age: 19},
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func Adduser(w http.ResponseWriter, r *http.Request) {
	var u User2
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
func updateUser(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	var updated User2
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, emp := range users {
		if emp.ID == id {
			users[i].Name = updated.Name
			users[i].ID = updated.ID
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}
func Getuserbyid(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}
	for _, u := range users {
		if u.ID == id {
			w.Header().Set("content-Type", "application/json")
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}
func Crudmethod2() {
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
		case http.MethodGet:
			Getuserbyid(w, r)
		case http.MethodPut:
			updateUser(w, r)
		case http.MethodDelete:
			DeleteUser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
