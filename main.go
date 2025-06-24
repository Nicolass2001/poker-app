package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocket Upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow connections from any origin
}

// Room struct to manage users and messages
type Room struct {
	clients map[*websocket.Conn]bool // Active connections
	mu      sync.Mutex               // Mutex to handle concurrent access
	broadcast chan []byte            // Channel for broadcasting messages
}

func newRoom() *Room {
	return &Room{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
}

func (room *Room) run() {
	for {
		// Broadcast messages to all clients
		message := <-room.broadcast
		room.mu.Lock()
		for client := range room.clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println("Error writing to client:", err)
				client.Close()
				delete(room.clients, client)
			}
		}
		room.mu.Unlock()
	}
}

func (room *Room) addClient(conn *websocket.Conn) {
	room.mu.Lock()
	room.clients[conn] = true
	room.mu.Unlock()
}

func (room *Room) removeClient(conn *websocket.Conn) {
	room.mu.Lock()
	delete(room.clients, conn)
	room.mu.Unlock()
	conn.Close()
}

func main() {
	room := newRoom()
	go room.run() // Start the room's broadcast loop

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error upgrading connection:", err)
			return
		}
		room.addClient(conn)
		defer room.removeClient(conn)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				break
			}
			room.broadcast <- message // Send the message to all clients
		}
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
