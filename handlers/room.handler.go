package handlers

import "net/http"

func RoomHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[len("/room/"):]
	nickname := r.URL.Query().Get("nickname")
	roomsMu.Lock()
	room, ok := rooms[code]
	roomsMu.Unlock()
	if !ok {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	tmpl.ExecuteTemplate(w, "room.html", struct {
		Code     string
		Nickname string
	}{
		Code:     room.Code,
		Nickname: nickname,
	})
}
