package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"

	"poker-app/poker"
)

var (
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	roomsMu sync.Mutex
	rooms   = map[string]*Room{}
)

type Room struct {
	Code          string
	Game          *poker.Game
	StartingChips int
	Connections   map[string]*websocket.Conn
	Mu            sync.Mutex
}

var upgrader = websocket.Upgrader{}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/join", joinHandler)
	http.HandleFunc("/room/", roomHandler)
	http.HandleFunc("/ws/", wsHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
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

func joinHandler(w http.ResponseWriter, r *http.Request) {
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

func roomHandler(w http.ResponseWriter, r *http.Request) {
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

func wsHandler(w http.ResponseWriter, r *http.Request) {
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
	var playersNames []string
	for p := range room.Connections {
		playersNames = append(playersNames, p)
	}
	room.Mu.Unlock()

	room.broadcast([]byte(nickname + " has joined the room."))
	room.broadcast([]byte("Current players: " + strings.Join(playersNames, ", ")))

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

func (r *Room) broadcast(msg []byte) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	for _, c := range r.Connections {
		c.WriteMessage(websocket.TextMessage, msg)
	}
}

func generateCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 4)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
