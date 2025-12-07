package day11

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	ID   int
	Name string
}

var employees1 = []Employee{
	{ID: 1, Name: "Akash"},
	{ID: 2, Name: "kushal"},
}

func Getemployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(employees1)
}
func Addemployees(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	emp.ID = len(employees1) + 1
	employees1 = append(employees1, emp)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}
func updateemployee(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/employees/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	var updated Employee
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, emp := range employees1 {
		if emp.ID == id {
			employees1[i].Name = updated.Name
			employees1[i].ID = updated.ID
			json.NewEncoder(w).Encode(employees1[i])
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}
func Deleteemployee(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/employees/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	for i, emp := range employees1 {
		if emp.ID == id {
			employees1 = append(employees1[:i], employees1[i+1:]...)
			fmt.Fprintf(w, "Deleted employee with id %d", id)
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}
func Crudmethod() {
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Getemployees(w, r)
		case http.MethodPost:
			Addemployees(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateemployee(w, r)
		case http.MethodDelete:
			Deleteemployee(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
