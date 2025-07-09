package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"encoding/json"
	"log"
)

var (
	clients   = make(map[net.Conn]bool)
	mu        sync.Mutex
	broadcast = make(chan string, 100) // Buffered channel
)

type Point struct {
    X float32 `json:"x"`
    Y float32 `json:"y"`
}

func main() {
	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	go handleMessages()

	fmt.Println("Chat server started on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		mu.Lock()
		clients[conn] = true
		mu.Unlock()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
	}()
	fmt.Println("client joined")
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}
		broadcast <- message
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		fmt.Print(msg)
msg = strings.ReplaceAll(msg, "\x00", "")
//		msg = strings.TrimSuffix(msg, "\n")
/*
		// Split the string by the comma
		parts := strings.Split(msg, ",")
		newposition := ""
		// Check if we have exactly two parts
		if len(parts) == 2 {
			var num1, num2 int
			// Parse the two numbers
			str := parts[0][1:]
			_, err1 := fmt.Sscanf(str, "%d", &num1)
			_, err2 := fmt.Sscanf(parts[1], "%d", &num2)

			if err1 == nil && err2 == nil {
				//fmt.Printf("Extracted numbers: %d and %d\n", num1, num2)
			} else {
				fmt.Println("Error parsing numbers")
				fmt.Println("1: ", err1, " ", str)
				fmt.Println("2: ", err2)
			}
			newposition = fmt.Sprint(num1, ",", num2,",\n")
		} else {
			fmt.Println("Invalid input format")
		}
*/

var point Point

    // Parse the JSON string
    err := json.Unmarshal([]byte(msg), &point)
    if err != nil {
        log.Fatalf("Error parsing JSON: %v", err)
    }

    // Use the data
    fmt.Printf("X: %d, Y: %d\n", point.X, point.Y)
newposition := fmt.Sprint(point.X, ",", point.Y,",\n")
		mu.Lock()
		for client := range clients {
			_, err := client.Write([]byte(newposition))
			if err != nil {
				fmt.Println("Error sending message:", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}

