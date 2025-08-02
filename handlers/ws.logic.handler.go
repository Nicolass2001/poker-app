package handlers

import (
	"encoding/json"
	"log"
	"poker-app/poker"
)

type message struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type gameStateData struct {
	Players []poker.PlayerInfo `json:"players"`
}

func (r *Room) broadcastGameState() {
	r.Mu.Lock()
	players := r.Game.GetPlayersInfo()
	r.Mu.Unlock()

	r.broadcast(message{
		Type: "gameState",
		Data: gameStateData{
			Players: players,
		},
	})
}

func (r *Room) handleIncomingMessage(nickname string, msg []byte) {
	var m message
	if err := json.Unmarshal(msg, &m); err != nil {
		log.Println("Error unmarshalling message:", err)
		return
	}

	switch m.Type {
	case "action":
		// TODO: Handle actions like call, raise, fold, etc.
	case "startGame":
		r.Game.StartGame()
	default:
		log.Println("Unknown message type:", m.Type)
	}

	r.broadcastGameState()
}
