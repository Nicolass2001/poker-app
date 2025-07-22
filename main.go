package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	roomsMu sync.Mutex
	rooms   = map[string]*Room{}
)

type Room struct {
	Code        string
	Password    string
	Connections map[*websocket.Conn]bool
	Mu          sync.Mutex
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
	pass := r.FormValue("password")

	roomsMu.Lock()
	rooms[code] = &Room{
		Code:        code,
		Password:    pass,
		Connections: make(map[*websocket.Conn]bool),
	}
	roomsMu.Unlock()

	http.Redirect(w, r, "/room/"+code, http.StatusSeeOther)
}

func joinHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	pass := r.FormValue("password")

	roomsMu.Lock()
	room, ok := rooms[code]
	roomsMu.Unlock()

	if !ok || room.Password != pass {
		http.Error(w, "Invalid code or password", http.StatusForbidden)
		return
	}

	http.Redirect(w, r, "/room/"+code, http.StatusSeeOther)
}

func roomHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[len("/room/"):]
	roomsMu.Lock()
	room, ok := rooms[code]
	roomsMu.Unlock()
	if !ok {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	tmpl.ExecuteTemplate(w, "room.html", room)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[len("/ws/"):]
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
	room.Connections[conn] = true
	room.Mu.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		room.broadcast(msg)
	}

	room.Mu.Lock()
	delete(room.Connections, conn)
	room.Mu.Unlock()
	conn.Close()
}

func (r *Room) broadcast(msg []byte) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	for c := range r.Connections {
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
