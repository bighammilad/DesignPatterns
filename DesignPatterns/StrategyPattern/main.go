package main

import "fmt"

// Strategy interface defines the common behavior for all strategies
type Strategy interface {
	Execute(a, b int) int
}

// AddStrategy implements the Strategy interface for addition
type AddStrategy struct{}

// Execute method for AddStrategy
func (a AddStrategy) Execute(x, y int) int {
	return x + y
}

// SubtractStrategy implements the Strategy interface for subtraction
type SubtractStrategy struct{}

// Execute method for SubtractStrategy
func (s SubtractStrategy) Execute(x, y int) int {
	return x - y
}

// MultiplyStrategy implements the Strategy interface for multiplication
type MultiplyStrategy struct{}

// Execute method for MultiplyStrategy
func (m MultiplyStrategy) Execute(x, y int) int {
	return x * y
}

// Context holds a reference to a Strategy and can switch between strategies dynamically
type Context struct {
	strategy Strategy
}

// SetStrategy changes the strategy of the context
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// ExecuteStrategy calls the Execute method of the current strategy
func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func main() {
	// Create a context and set a strategy
	context := &Context{}

	// Using the AddStrategy
	context.SetStrategy(AddStrategy{})
	fmt.Println("Addition: 3 + 2 =", context.ExecuteStrategy(3, 2)) // Output: 3 + 2 = 5

	// Using the SubtractStrategy
	context.SetStrategy(SubtractStrategy{})
	fmt.Println("Subtraction: 3 - 2 =", context.ExecuteStrategy(3, 2)) // Output: 3 - 2 = 1

	// Using the MultiplyStrategy
	context.SetStrategy(MultiplyStrategy{})
	fmt.Println("Multiplication: 3 * 2 =", context.ExecuteStrategy(3, 2)) // Output: 3 * 2 = 6
}
