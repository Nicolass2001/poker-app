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

type actionMessage struct {
	Type string     `json:"type"`
	Data actionData `json:"data"`
}

type actionData struct {
	Action string `json:"action"`
	Amount int    `json:"amount,omitempty"`
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
	err := json.Unmarshal(msg, &m)
	if err != nil {
		log.Println("Error unmarshalling message:", err)
		return
	}

	switch m.Type {
	case "action":
		var actionMsg actionMessage
		err := json.Unmarshal(msg, &actionMsg)
		if err != nil {
			log.Println("Error unmarshalling action message:", err)
			return
		}

		r.Mu.Lock()
		player := r.Game.GetCurrentPlayer()
		r.Mu.Unlock()
		if player.Id != nickname {
			log.Println("Player", nickname, "is not the current player")
			return
		}

		r.handleAction(actionMsg.Data)

	case "startGame":
		r.Mu.Lock()
		err := r.Game.StartGame()
		r.Mu.Unlock()
		if err != nil {
			log.Println("Error starting game:", err)
			return
		}

	default:
		log.Println("Unknown message type:", m.Type)

	}
}

func (r *Room) handleAction(actionData actionData) {
	switch actionData.Action {
	case "call":
		r.Mu.Lock()
		err := r.Game.MakeAction(poker.ActionCall, 0)
		r.Mu.Unlock()
		if err != nil {
			log.Println("Error making call action:", err)
			return
		}
	case "raise":
		r.Mu.Lock()
		err := r.Game.MakeAction(poker.ActionRaise, actionData.Amount)
		r.Mu.Unlock()
		if err != nil {
			log.Println("Error making raise action:", err)
			return
		}
	case "fold":
		r.Mu.Lock()
		err := r.Game.MakeAction(poker.ActionFold, 0)
		r.Mu.Unlock()
		if err != nil {
			log.Println("Error making fold action:", err)
			return
		}
	case "check":
		r.Mu.Lock()
		err := r.Game.MakeAction(poker.ActionCheck, 0)
		r.Mu.Unlock()
		if err != nil {
			log.Println("Error making check action:", err)
			return
		}
	case "allin":
		r.Mu.Lock()
		err := r.Game.MakeAction(poker.ActionAllIn, actionData.Amount)
		r.Mu.Unlock()
		if err != nil {
			log.Println("Error making all-in action:", err)
			return
		}
	default:
		log.Println("Unknown action:", actionData.Action)
		return
	}
}
