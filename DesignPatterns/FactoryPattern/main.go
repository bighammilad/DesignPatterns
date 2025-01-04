package main

import "fmt"

// Define an interface for Car
type Car interface {
	Drive() string // Method that all car types must implement
}

// Define a structure for SportsCar
type SportsCar struct{}

// Implement the Drive method for SportsCar
func (s *SportsCar) Drive() string {
	return "Driving a sports car!"
}

// Define a structure for SUVCar
type SUVCar struct{}

// Implement the Drive method for SUVCar
func (s *SUVCar) Drive() string {
	return "Driving an SUV car!"
}

// Factory function that returns a Car based on the car type
func CarFactory(carType string) Car {
	// Depending on the carType, return the appropriate car type
	switch carType {
	case "sports":
		return &SportsCar{} // Return a SportsCar if the type is "sports"
	case "suv":
		return &SUVCar{} // Return an SUVCar if the type is "suv"
	default:
		return nil // Return nil if the carType is unknown
	}
}

func main() {
	// Create a SportsCar using the Factory
	car1 := CarFactory("sports")
	fmt.Println(car1.Drive()) // Output: "Driving a sports car!"

	// Create an SUVCar using the Factory
	car2 := CarFactory("suv")
	fmt.Println(car2.Drive()) // Output: "Driving an SUV car!"
}
