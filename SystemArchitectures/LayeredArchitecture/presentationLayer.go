package main

import (
	"fmt"
)

func main() {
	// Initialize the layers.
	repo := NewUserRepository()
	service := NewUserService(repo)

	// Simulate creating users.
	service.CreateUser("Alice", 25)
	service.CreateUser("Bob", 30)

	// Simulate listing all users.
	fmt.Println("List of Users:")
	users := service.ListAllUsers()
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

	// Simulate retrieving a user by ID.
	fmt.Println("\nRetrieve User with ID 1:")
	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

	// Simulate an invalid user lookup.
	fmt.Println("\nRetrieve User with ID 99:")
	_, err = service.GetUser(99)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
