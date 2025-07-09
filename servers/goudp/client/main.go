package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	// Define the server address and port
	ADDRESS = "185.94.99.107:8080"
)

func main() {
	// Resolve UDP address
	addr, err := net.ResolveUDPAddr("udp", ADDRESS)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error dialing UDP:", err)
		return
	}
	defer conn.Close()

	// Handle incoming messages in a separate goroutine
	go listenForMessages(conn)

	// Handle user input
	handleUserInput(conn)
}

func listenForMessages(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			return
		}
		fmt.Printf("Message: %s\n", string(buffer[:n]))
	}
}

func handleUserInput(conn *net.UDPConn) {
	// Capture interrupt signal to gracefully exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nExiting chat client...")
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim the newline character from the message
		message = message[:len(message)-1]

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}
