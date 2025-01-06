package main

import (
	"fmt"
	"net"
	"time"
)

// Event represents the event structure that will be sent by the producer.
type Event struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// produceEvent generates events and sends them to the consumer
func produceEvent(conn net.Conn) {
	for {
		// Simulate producing an event every 3 seconds
		event := Event{
			ID:      fmt.Sprintf("event-%d", time.Now().Unix()),
			Message: "New Event Created",
		}

		// Send the event to the consumer (connected via TCP)
		_, err := fmt.Fprintf(conn, "%v\n", event)
		if err != nil {
			fmt.Println("Error sending event:", err)
			return
		}
		fmt.Printf("Produced Event: %v\n", event)

		// Wait for 3 seconds before producing the next event
		time.Sleep(3 * time.Second)
	}
}

func main() {
	// Set up the TCP server (Producer)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Producer is running and waiting for consumer...")

	// Accept a connection from the consumer
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	// Start producing events and sending them to the consumer
	produceEvent(conn)
}
