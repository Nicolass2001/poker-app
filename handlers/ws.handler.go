package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[len("/ws/"):]
	nickname := r.URL.Query().Get("nickname")
	roomsMu.Lock()
	room, ok := rooms[code]
	roomsMu.Unlock()
	if !ok {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	room.Mu.Lock()
	room.Connections[nickname] = conn
	room.PlayersNames = append(room.PlayersNames, nickname)
	room.Mu.Unlock()

	room.broadcastGameState()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		room.broadcast(msg)
	}

	room.Mu.Lock()
	delete(room.Connections, nickname)
	room.Mu.Unlock()
	conn.Close()
}

func (r *Room) broadcast(msg []byte) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	for _, c := range r.Connections {
		c.WriteMessage(websocket.TextMessage, msg)
	}
}

func (r *Room) broadcastGameState() {
	r.Mu.Lock()
	playersNames := make([]string, len(r.PlayersNames))
	copy(playersNames, r.PlayersNames)
	r.Mu.Unlock()

	r.broadcast([]byte("Current players: " + strings.Join(playersNames, ", ")))
}
