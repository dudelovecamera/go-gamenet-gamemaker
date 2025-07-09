package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

const (
	// Define the server address and port
	ADDRESS    = ":8080"
	START_PORT = 8080
	END_PORT   = 8090
)

var (
	// Mutex to protect shared resources
	mu      sync.Mutex
	clients = make(map[string]net.UDPAddr)
)

func main() {

	// Start listening on multiple ports
	for port := START_PORT; port <= END_PORT; port++ {
		go startUDPServer(port)
		fmt.Println("Chat server started on", port)
	}

	// Keep the main goroutine alive
	select {}
}

func startUDPServer(port int) {
	// Resolve UDP address
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	// Start listening for messages
	for {
		handleConnection(conn)
	}
}

func handleConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading from UDP:", err)
		return
	}

	message := string(buffer[:n])
	fmt.Printf("Received message from %s: %s\n", addr.String(), message)

	// Register client
	mu.Lock()
	clients[addr.String()] = *addr
	mu.Unlock()

	// Broadcast message to all clients
	broadcastMessage(conn, message, addr)
}

func broadcastMessage(conn *net.UDPConn, message string, sender *net.UDPAddr) {
	mu.Lock()
	defer mu.Unlock()

	msg := strings.ReplaceAll(message, "\x00", "")
	var point Point

	// Parse the JSON string
	err := json.Unmarshal([]byte(msg), &point)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Use the data
	fmt.Printf("X: %d, Y: %d\n", point.X, point.Y)
	newposition := fmt.Sprint(point.X, ",", point.Y, ",\n")

	for addrStr, addr := range clients {
		_ = addrStr
		//		if addrStr != sender.String() {
		_, err := conn.WriteToUDP([]byte(newposition), &addr)
		if err != nil {
			fmt.Println("Error broadcasting message:", err)
		}
		//		}
	}
}
