package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// ==========================================
// STEP 1: Define the Message struct
// Fields: Sender, Text, Time (all strings)
// Don't forget the json tags!
// ==========================================

// YOUR CODE HERE:

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// ==========================================
// STEP 2: Create a map to track connected clients
//
// var clients = make(map[*websocket.Conn]bool)
// ==========================================

// YOUR CODE HERE:

// ==========================================
// STEP 3: Write the broadcast function
// Loop through all clients and send the message
// using conn.WriteJSON(msg)
// If there's an error, close and delete that client
// ==========================================

// YOUR CODE HERE:

// ==========================================
// STEP 4: Write the WebSocket handler
// This function should:
//   a. Upgrade the connection
//   b. defer conn.Close()
//   c. Add the client to the map
//   d. Print how many clients are connected
//   e. Loop reading JSON messages (conn.ReadJSON)
//   f. For each message, print it and call broadcast
//   g. On error, delete the client and break
// ==========================================

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// YOUR CODE HERE:

}

func main() {
	// WebSocket route
	http.HandleFunc("/ws", handleWebSocket)

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Chat server starting on http://localhost:8080")
	fmt.Println("Open multiple browser tabs to test!")
	http.ListenAndServe(":8080", nil)
}

// ==========================================
// TO RUN:
//   1. go mod init chatserver
//   2. go get github.com/gorilla/websocket
//   3. go run main.go
//   4. Open http://localhost:8080
//   5. Open a SECOND tab at http://localhost:8080
//   6. Send messages from both tabs!
// ==========================================
