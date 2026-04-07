package main

// Completed server — same as Lesson 9. Focus on the frontend.

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
	Time   string `json:"time"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

func broadcast(msg Message) {
	for conn := range clients {
		err := conn.WriteJSON(msg)
		if err != nil {
			conn.Close()
			delete(clients, conn)
		}
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	fmt.Println("Client connected. Total:", len(clients))

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Client disconnected. Total:", len(clients)-1)
			delete(clients, conn)
			break
		}

		fmt.Printf("[%s] %s: %s\n", msg.Time, msg.Sender, msg.Text)
		broadcast(msg)
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Chat server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
