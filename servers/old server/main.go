package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type Player struct {
	ID       string
	Position string // e.g., "x,y"
}

var (
	players = make(map[net.Conn]*Player)
	mu      sync.Mutex
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Game server started on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	//reader := bufio.NewReader(conn)

	// Assign a new player ID (for simplicity, using the connection address)
	playerID := conn.RemoteAddr().String()
	mu.Lock()
	players[conn] = &Player{ID: playerID, Position: "0,0"}
	mu.Unlock()

	// Notify other players about the new player
	broadcast(fmt.Sprintf("0 %s %s\n", playerID, "0,0"))

	for {
		/*
			msg, err := reader.ReadString('|')
			message := strings.ReplaceAll(msg, "|", "")
			fmt.Println(message)
			if err != nil {
				fmt.Println("Error reading message:", err)
				break
			}
		*/
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			//remove_client(conn)
			return
		}
		index := 0
		mm := buffer_get_string(buf, &index, reqLen)
		mm = strings.ReplaceAll(mm, "|", "")
		fmt.Println("I will send this to everyone now:", mm)
		fmt.Println("rarar")
		// Update player position
		//mu.Lock()
		//if player, exists := players[conn]; exists {
		//	player.Position = message
		//		broadcast(fmt.Sprintf("1 %s %s \n", playerID, message))
		//_ = message
		broadcast(fmt.Sprintf("1 %s %s", "ramin", mm))
		//}
		//mu.Unlock()
	}

	// Remove player on disconnect
	mu.Lock()
	delete(players, conn)
	mu.Unlock()
	broadcast(fmt.Sprintf("2 %s 0,0,\n", playerID))
}

func broadcast(message string) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(message)
	for conn := range players {
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			conn.Close()
			delete(players, conn)
		}
	}
}

func buffer_get_string(buff []byte, index *int, l int) string {
	str := ""
	start := *index
	for i := start; i < l; i++ {
		*index = i
		if buff[i] == 0 {
			*index += 1
			break
		} else {
			str += string(buff[i])
		}
	}
	return str
}
