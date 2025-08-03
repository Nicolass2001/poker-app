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
	room.broadcastPersonalInfoToPlayers()

	for {
		_, msg, err := conn.ReadMessage()
		println("Received message:", string(msg))
		if err != nil {
			break
		}
		room.handleIncomingMessage(nickname, msg)
		room.broadcastGameState()
		room.broadcastPersonalInfoToPlayers()
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

func (r *Room) sendPlayerInfo(nickname string, msg any) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	r.Connections[nickname].WriteJSON(msg)
}

func (r *Room) sendErrorMessage(nickname string, errMsg string) {
	r.sendPlayerInfo(nickname, message{
		Type: "error",
		Data: errMsg,
	})
}
