package handlers

import (
	"sync"

	"github.com/gorilla/websocket"

	"poker-app/poker"
)

var (
	roomsMu sync.Mutex
	rooms   = map[string]*Room{}
)

type Room struct {
	Code          string
	Game          *poker.Game
	StartingChips int
	PlayersNames  []string
	Connections   map[string]*websocket.Conn
	Mu            sync.Mutex
}
