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
	GameState      string             `json:"gameState"`
	Players        []poker.PlayerInfo `json:"players"`
	CommunityCards []poker.Card       `json:"communityCards"`
}

type playerInfo struct {
	Player poker.Player `json:"player"`
}

func (r *Room) broadcastGameState() {
	r.Mu.Lock()
	gameState := r.Game.GetGameState()
	players := r.Game.GetPlayersInfo()
	communityCards := r.Game.GetCommunityCards()
	r.Mu.Unlock()

	r.broadcast(message{
		Type: "gameState",
		Data: gameStateData{
			GameState:      gameState,
			Players:        players,
			CommunityCards: communityCards,
		},
	})
}

func (r *Room) broadcastPersonalInfoToPlayers() {
	messages := make(map[string]message)

	r.Mu.Lock()
	for nick := range r.Connections {
		player, err := r.Game.GetPlayerById(nick)
		if err != nil {
			log.Println("")
			return
		}
		messages[nick] = message{
			Type: "playerInfo",
			Data: playerInfo{
				Player: player,
			},
		}
	}
	r.Mu.Unlock()

	for nick, msg := range messages {
		r.sendPlayerInfo(nick, msg)
	}
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
		err := r.Game.StartGame()
		if err != nil {
			log.Println("Error starting game:", err)
			return
		}
	default:
		log.Println("Unknown message type:", m.Type)
	}

	r.broadcastGameState()
}
