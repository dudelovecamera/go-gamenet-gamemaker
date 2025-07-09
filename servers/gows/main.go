package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	mu        sync.Mutex
	broadcast = make(chan string, 100) // Buffered channel
	upgrader  = websocket.Upgrader{}   // Upgrader to upgrade HTTP connection to WebSocket
)

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func main() {
	http.HandleFunc("/", handleConnections)

	go handleMessages()

	fmt.Println("Chat server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	fmt.Println("Client joined")
	for {
		_, msg, err := conn.ReadMessage() // Correctly read the message
		if err != nil {
			log.Println("Error reading message:", err)
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			return
		}
		broadcast <- string(msg) // Convert byte slice to string
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		fmt.Print(msg)
		msg = strings.ReplaceAll(msg, "\x00", "")

		var point Point

		// Parse the JSON string
		err := json.Unmarshal([]byte(msg), &point)
		if err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}

		// Use the data
		fmt.Printf("X: %f, Y: %f\n", point.X, point.Y)
		newposition := fmt.Sprintf("%f,%f,\n", point.X, point.Y)

		mu.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(newposition))
			if err != nil {
				fmt.Println("Error sending message:", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}
