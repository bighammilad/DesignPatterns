package main

import "fmt"

// Observer interface defines the method that observers need to implement
type Observer interface {
	Update(message string)
}

// Subject interface defines the method for attaching, detaching, and notifying observers
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(message string)
}

// ConcreteSubject implements the Subject interface
type ConcreteSubject struct {
	observers []Observer
	state     string
}

// Attach adds an observer to the list
func (s *ConcreteSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

// Detach removes an observer from the list
func (s *ConcreteSubject) Detach(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// Notify sends a message to all observers
func (s *ConcreteSubject) Notify(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

// SetState changes the state and notifies the observers
func (s *ConcreteSubject) SetState(state string) {
	s.state = state
	s.Notify(state)
}

// ConcreteObserver implements the Observer interface
type ConcreteObserver struct {
	name string
}

// Update is called when the subject's state changes
func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("%s received update: %s\n", o.name, message)
}

func main() {
	// Creating the subject and observers
	subject := &ConcreteSubject{}
	observer1 := &ConcreteObserver{name: "Observer 1"}
	observer2 := &ConcreteObserver{name: "Observer 2"}

	// Attaching observers to the subject
	subject.Attach(observer1)
	subject.Attach(observer2)

	// Changing the state of the subject, which will notify the observers
	subject.SetState("State 1")
	subject.SetState("State 2")

	// Detach observer1 and update the state
	subject.Detach(observer1)
	subject.SetState("State 3")
}
