package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// ==========================================
// STEP 1: Update the Message struct
// Add a Type field (string, json tag "type")
// This lets us distinguish between "chat",
// "system", and "usercount" messages.
// ==========================================

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
	Time   string `json:"time"`
	// YOUR CODE HERE: add Type field
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]string) // conn → username

// ==========================================
// STEP 2: Add a mutex
// Create a sync.Mutex to protect the clients map.
//
// var mutex = sync.Mutex{}
// ==========================================

// YOUR CODE HERE:
var mutex = sync.Mutex{}

func currentTime() string {
	return time.Now().Format("3:04 PM")
}

// ==========================================
// STEP 3: Update broadcast to use the mutex
// Lock before looping, unlock when done.
// ==========================================

func broadcast(msg Message) {
	// YOUR CODE HERE: add mutex.Lock() and defer mutex.Unlock()

	for conn := range clients {
		err := conn.WriteJSON(msg)
		if err != nil {
			conn.Close()
			delete(clients, conn)
		}
	}
}

// Helper: broadcast current user count
func broadcastUserCount() {
	countMsg := Message{
		Type: "usercount",
		Text: fmt.Sprintf("%d", len(clients)),
		Time: currentTime(),
	}
	// Note: don't lock here — the caller should hold the lock
	// or we should broadcast outside the lock
	for conn := range clients {
		conn.WriteJSON(countMsg)
	}
}

// ==========================================
// STEP 4: Update the WebSocket handler
//
// When a client connects:
//   a. Lock, add to clients (with empty username for now)
//   b. Unlock
//
// When a message is received:
//   a. If this is the first message from the client,
//      save their username and broadcast a join message
//   b. Set msg.Type = "chat" and broadcast
//
// When a client disconnects:
//   a. Lock, get their username, delete from clients
//   b. Broadcast a leave message
//   c. Broadcast updated user count
//   d. Unlock
// ==========================================

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Register client
	mutex.Lock()
	clients[conn] = "" // username unknown until first message
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

			// Broadcast leave message if we knew their name
			if username != "" {
				leaveMsg := Message{
					Type:   "system",
					Sender: "System",
					Text:   username + " left the chat",
					Time:   currentTime(),
				}
				for c := range clients {
					c.WriteJSON(leaveMsg)
				}
			}
			broadcastUserCount()
			mutex.Unlock()

			fmt.Printf("%s disconnected. Total: %d\n", username, len(clients))
			break
		}

		// ==========================================
		// STEP 5: Handle first message (join notification)
		// If clients[conn] is empty, this is the first
		// message — save their username and broadcast
		// a "joined" system message.
		//
		// Then set msg.Type = "chat" and broadcast.
		// ==========================================

		mutex.Lock()
		if clients[conn] == "" {
			clients[conn] = msg.Sender
			joinMsg := Message{
				Type:   "system",
				Sender: "System",
				Text:   msg.Sender + " joined the chat",
				Time:   currentTime(),
			}
			for c := range clients {
				c.WriteJSON(joinMsg)
			}
			broadcastUserCount()
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

	fmt.Println("Chat server starting on http://localhost:8080")
	fmt.Println("Open multiple tabs to test broadcasting!")
	http.ListenAndServe(":8080", nil)
}
