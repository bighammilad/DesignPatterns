package main

import "fmt"

// Component interface defines the method that both the concrete component and decorators will implement
type Component interface {
	Operation() string
}

// ConcreteComponent is a basic implementation of the Component interface
type ConcreteComponent struct{}

// Operation method for ConcreteComponent
func (c *ConcreteComponent) Operation() string {
	return "Basic Operation"
}

// Decorator struct holds a reference to the Component and is used to add additional behavior
type Decorator struct {
	component Component
}

// NewDecorator creates a new Decorator instance
func NewDecorator(c Component) *Decorator {
	return &Decorator{component: c}
}

// Operation method for Decorator, which adds extra behavior
func (d *Decorator) Operation() string {
	return d.component.Operation() + " with extra functionality"
}

// ConcreteDecoratorA is a specific decorator that extends the behavior of the component
type ConcreteDecoratorA struct {
	Decorator
}

// Operation method for ConcreteDecoratorA, which adds a specific behavior
func (c *ConcreteDecoratorA) Operation() string {
	return c.Decorator.Operation() + " from ConcreteDecoratorA"
}

// ConcreteDecoratorB is another specific decorator
type ConcreteDecoratorB struct {
	Decorator
}

// Operation method for ConcreteDecoratorB, which adds a different specific behavior
func (c *ConcreteDecoratorB) Operation() string {
	return c.Decorator.Operation() + " from ConcreteDecoratorB"
}

func main() {
	// Create the basic component
	component := &ConcreteComponent{}
	fmt.Println(component.Operation()) // Output: Basic Operation

	// Decorate with ConcreteDecoratorA
	decoratorA := &ConcreteDecoratorA{Decorator{component}}
	fmt.Println(decoratorA.Operation()) // Output: Basic Operation with extra functionality from ConcreteDecoratorA

	// Decorate with ConcreteDecoratorB
	decoratorB := &ConcreteDecoratorB{Decorator{component}}
	fmt.Println(decoratorB.Operation()) // Output: Basic Operation with extra functionality from ConcreteDecoratorB

	// Combine multiple decorators
	combinedDecorator := &ConcreteDecoratorB{Decorator{&ConcreteDecoratorA{Decorator{component}}}}
	fmt.Println(combinedDecorator.Operation()) // Output: Basic Operation with extra functionality from ConcreteDecoratorA with extra functionality from ConcreteDecoratorB
}
