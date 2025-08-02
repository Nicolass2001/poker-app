package handlers

import "net/http"

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

	http.Redirect(w, r, "/room/"+code+"?nickname="+nickname, http.StatusSeeOther)
}
