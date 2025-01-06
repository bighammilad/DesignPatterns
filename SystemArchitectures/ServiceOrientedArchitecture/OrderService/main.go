package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Order struct represents an order in the system.
type Order struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
}

// In-memory order database
var orders = make(map[int]Order)
var nextOrderID = 1

// Check if the user exists by calling the User Service API
func userExists(userID int) bool {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8081/users/%d", userID))
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	defer resp.Body.Close()
	return true
}

// CreateOrder creates a new order.
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	// Parse JSON data from the request body
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check if the user exists using the User Service
	if !userExists(order.UserID) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Store the order
	order.ID = nextOrderID
	orders[nextOrderID] = order
	nextOrderID++

	// Respond with the created order
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func main() {
	http.HandleFunc("/orders", CreateOrder) // Handle POST /orders

	// Start the server
	fmt.Println("Order Service is running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
