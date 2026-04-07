package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Message represents a chat message exchanged between client and server
type Message struct {
	Type   string `json:"type"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
	Time   string `json:"time"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	clients        = make(map[*websocket.Conn]string) // conn → username
	mutex          = sync.Mutex{}
	messageHistory []Message
	maxHistory     = 50
)

func currentTime() string {
	return time.Now().Format("3:04 PM")
}

func broadcast(msg Message) {
	mutex.Lock()
	defer mutex.Unlock()

	// Save chat and system messages to history (not typing)
	if msg.Type == "chat" || msg.Type == "system" {
		messageHistory = append(messageHistory, msg)
		if len(messageHistory) > maxHistory {
			messageHistory = messageHistory[1:]
		}
	}

	for conn := range clients {
		err := conn.WriteJSON(msg)
		if err != nil {
			conn.Close()
			delete(clients, conn)
		}
	}
}

func broadcastExcept(msg Message, exclude *websocket.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	for conn := range clients {
		if conn != exclude {
			conn.WriteJSON(msg)
		}
	}
}

func sendUserCount() {
	countMsg := Message{
		Type: "usercount",
		Text: fmt.Sprintf("%d", len(clients)),
		Time: currentTime(),
	}
	for conn := range clients {
		conn.WriteJSON(countMsg)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Register client
	mutex.Lock()
	clients[conn] = ""

	// Send message history to new client
	for _, msg := range messageHistory {
		conn.WriteJSON(msg)
	}
	mutex.Unlock()

	fmt.Println("Client connected. Total:", len(clients))

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			// Client disconnected
			mutex.Lock()
			username := clients[conn]
			delete(clients, conn)

			if username != "" {
				leaveMsg := Message{
					Type:   "system",
					Sender: "System",
					Text:   username + " left the chat",
					Time:   currentTime(),
				}
				// Save to history and send to remaining clients
				messageHistory = append(messageHistory, leaveMsg)
				if len(messageHistory) > maxHistory {
					messageHistory = messageHistory[1:]
				}
				for c := range clients {
					c.WriteJSON(leaveMsg)
				}
				sendUserCount()
			}
			mutex.Unlock()

			fmt.Printf("%s disconnected. Total: %d\n", username, len(clients))
			break
		}

		// Handle typing indicators (forward to others, don't save)
		if msg.Type == "typing" || msg.Type == "stopped_typing" {
			broadcastExcept(msg, conn)
			continue
		}

		// Validate the message
		msg.Text = strings.TrimSpace(msg.Text)
		if msg.Text == "" {
			continue
		}
		if len(msg.Text) > 500 {
			msg.Text = msg.Text[:500]
		}
		if msg.Sender == "" {
			msg.Sender = "Anonymous"
		}

		// Server sets the timestamp
		msg.Time = currentTime()

		// Handle first message (join notification)
		mutex.Lock()
		if clients[conn] == "" {
			clients[conn] = msg.Sender

			joinMsg := Message{
				Type:   "system",
				Sender: "System",
				Text:   msg.Sender + " joined the chat",
				Time:   currentTime(),
			}
			messageHistory = append(messageHistory, joinMsg)
			if len(messageHistory) > maxHistory {
				messageHistory = messageHistory[1:]
			}
			for c := range clients {
				c.WriteJSON(joinMsg)
			}
			sendUserCount()
		}
		mutex.Unlock()

		msg.Type = "chat"
		fmt.Printf("[%s] %s: %s\n", msg.Time, msg.Sender, msg.Text)
		broadcast(msg)
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("=================================")
	fmt.Println("  ChatApp Server")
	fmt.Println("  http://localhost:8080")
	fmt.Println("=================================")
	http.ListenAndServe(":8080", nil)
}
