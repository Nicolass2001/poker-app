package handlers

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"

	"poker-app/poker"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	code := generateCode()
	nickname := r.FormValue("nickname")
	smallBlind, err := strconv.Atoi(r.FormValue("small_blind"))
	if err != nil {
		http.Error(w, "Invalid small blind", http.StatusBadRequest)
		return
	}
	bigBlind, err := strconv.Atoi(r.FormValue("big_blind"))
	if err != nil {
		http.Error(w, "Invalid big blind", http.StatusBadRequest)
		return
	}
	startingChips, err := strconv.Atoi(r.FormValue("starting_chips"))
	if err != nil {
		http.Error(w, "Invalid starting chips", http.StatusBadRequest)
		return
	}

	game := poker.NewGame(smallBlind, bigBlind)

	roomsMu.Lock()
	rooms[code] = &Room{
		Code:          code,
		Game:          game,
		StartingChips: startingChips,
		Connections:   make(map[string]*websocket.Conn),
	}
	roomsMu.Unlock()

	http.Redirect(w, r, "/room/"+code+"?nickname="+nickname, http.StatusSeeOther)
}

func generateCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 4)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
