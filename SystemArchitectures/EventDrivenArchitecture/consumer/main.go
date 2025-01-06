package main

import (
	"bufio"
	"fmt"
	"net"
)

// Event represents the event structure received from the producer.
type Event struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// consumeEvent listens for events from the producer
func consumeEvent(conn net.Conn) {
	// Create a buffered reader to read events from the connection
	reader := bufio.NewReader(conn)
	for {
		// Read the event from the producer
		eventString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading event:", err)
			return
		}

		// Process the event (in this case, just print it)
		fmt.Printf("Consumed Event: %v\n", eventString)
	}
}

func main() {
	// Set up the TCP connection to the producer
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to producer:", err)
		return
	}
	defer conn.Close()

	// Start consuming events from the producer
	consumeEvent(conn)
}
