package main

import (
	"fmt"
	"sync"
)

// Define the Singleton structure
type Singleton struct {
	// You can store shared system state here
	data string
}

// Private variable to hold the Singleton instance
var instance *Singleton
var once sync.Once

// Function to get the Singleton instance
func GetInstance() *Singleton {
	// Ensure the instance is created only once
	once.Do(func() {
		instance = &Singleton{
			data: "Some shared data", // default value
		}
	})
	return instance
}

// Method to display the data of the Singleton
func (s *Singleton) ShowData() {
	fmt.Println(s.data)
}

func main() {
	// Attempt to get multiple instances
	instance1 := GetInstance()
	instance2 := GetInstance()

	// Display the shared data
	instance1.ShowData()

	// Check if instance1 and instance2 are the same instance
	if instance1 == instance2 {
		fmt.Println("instance1 and instance2 are the same instance (Singleton Pattern)")
	} else {
		fmt.Println("instance1 and instance2 are different instances.")
	}
}
