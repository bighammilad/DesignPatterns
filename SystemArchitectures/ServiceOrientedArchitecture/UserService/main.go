package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// User struct represents a user in the system.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// In-memory user database (just for demonstration)
var users = make(map[int]User)
var nextID = 1

// CreateUser creates a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	// Parse JSON data from the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	// Assign a new ID and store the user in the map
	user.ID = nextID
	users[nextID] = user
	nextID++

	// Respond with the created user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser retrieves a user by ID.
func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Respond with the user
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/users", CreateUser) // Handle POST /users
	http.HandleFunc("/users/", GetUser)   // Handle GET /users/{id}

	// Start the server
	fmt.Println("User Service is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
