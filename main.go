package main

import (
	"log"
	"net/http"

	"poker-app/handlers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/create", handlers.CreateHandler)
	http.HandleFunc("/join", handlers.JoinHandler)
	http.HandleFunc("/room/", handlers.RoomHandler)
	http.HandleFunc("/ws/", handlers.WsHandler)

	log.Println("Server running at http://localhost:7777")
	log.Fatal(http.ListenAndServe(":7777", nil))
}
