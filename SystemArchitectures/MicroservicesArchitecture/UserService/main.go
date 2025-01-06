package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = map[string]User{
	"1": {ID: "1", Name: "John Doe"},
	"2": {ID: "2", Name: "Jane Smith"},
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user, exists := users[id]

	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/user", getUser)

	fmt.Println("User Service running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
