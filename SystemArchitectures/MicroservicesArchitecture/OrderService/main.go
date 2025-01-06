package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Order struct {
	OrderID string `json:"order_id"`
	UserID  string `json:"user_id"`
	Product string `json:"product"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	// Simulate reading an order from the request body
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Call User Service to get user details
	userResp, err := http.Get(fmt.Sprintf("http://localhost:8081/user?id=%s", order.UserID))
	if err != nil {
		http.Error(w, "Failed to call User Service", http.StatusInternalServerError)
		return
	}
	defer userResp.Body.Close()

	if userResp.StatusCode != http.StatusOK {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userBody, err := io.ReadAll(userResp.Body)
	if err != nil {
		http.Error(w, "Failed to read User Service response", http.StatusInternalServerError)
		return
	}

	var user User
	err = json.Unmarshal(userBody, &user)
	if err != nil {
		http.Error(w, "Failed to parse user data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Order ID: %s\nProduct: %s\nUser: %s", order.OrderID, order.Product, user.Name)
}

func main() {
	http.HandleFunc("/order", createOrder)

	fmt.Println("Order Service running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
