package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// Define the server address
	serverAddr := "ws://127.0.0.1:8080" // Adjust the path if necessary

	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server:", serverAddr)

	// Send a message to the server
	message := `{"x": 10, "y": 20}` // Example JSON message
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}
	fmt.Println("Message sent:", message)

	// Listen for messages from the server
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Received message: %s\n", msg)
	}

	// Optional: Sleep for a while before closing the connection
	time.Sleep(1 * time.Second)
}
