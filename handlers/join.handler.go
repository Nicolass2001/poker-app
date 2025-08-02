package handlers

import (
	"net/http"

	"poker-app/poker"
)

func JoinHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	nickname := r.FormValue("nickname")

	roomsMu.Lock()
	room, ok := rooms[code]
	roomsMu.Unlock()

	if !ok {
		http.Error(w, "Invalid code", http.StatusForbidden)
		return
	}

	room.Mu.Lock()
	_, exists := room.Connections[nickname]
	room.Mu.Unlock()
	if exists {
		http.Error(w, "Nickname already taken", http.StatusConflict)
		return
	}

	player := poker.NewPlayer(nickname, nickname, room.StartingChips)

	room.Mu.Lock()
	err := room.Game.AddPlayer(player)
	room.Mu.Unlock()
	if err != nil {
		http.Error(w, "Error adding player: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/room/"+code+"?nickname="+nickname, http.StatusSeeOther)
}
