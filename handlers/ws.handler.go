package handlers

import (
	"log"
	"net/http"

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

func (r *Room) broadcast(msg any) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	for _, c := range r.Connections {
		c.WriteJSON(msg)
	}
}

type message struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

func (r *Room) broadcastGameState() {
	r.Mu.Lock()
	players := r.Game.GetPlayersInfo()
	r.Mu.Unlock()

	r.broadcast(message{
		Type: "gameState",
		Data: players,
	})
}
