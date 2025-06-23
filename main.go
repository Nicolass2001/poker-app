package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true }, // Permite conexiones desde cualquier origen
}

func main() {
    // Serve static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Routes
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/ws", wsHandler)

    // Start server
    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/index.html")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil) // Actualiza la conexión HTTP a WebSocket
    if err != nil {
        fmt.Println("Error upgrading connection:", err)
        return
    }
    defer conn.Close()

    for {
        // Lee mensajes desde el cliente
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Error reading message:", err)
            break
        }
        fmt.Printf("Received: %s\n", message)

        // Envía una respuesta al cliente
        err = conn.WriteMessage(messageType, []byte(fmt.Sprintf("Echo: %s", message)))
        if err != nil {
            fmt.Println("Error writing message:", err)
            break
        }
    }
}
