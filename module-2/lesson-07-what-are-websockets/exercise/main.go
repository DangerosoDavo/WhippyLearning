package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// The upgrader handles the HTTP → WebSocket handshake
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections for development
	},
}

// ==========================================
// EXERCISE 1: Echo handler
// Complete this WebSocket handler so that it:
//   1. Upgrades the connection
//   2. Loops forever reading messages
//   3. Prints each message to the terminal
//   4. Sends the same message BACK to the client (echo)
//
// Hint for the loop:
//   for {
//       _, msg, err := conn.ReadMessage()
//       if err != nil {
//           fmt.Println("Connection closed:", err)
//           break
//       }
//       fmt.Println("Received:", string(msg))
//       conn.WriteMessage(websocket.TextMessage, msg)
//   }
// ==========================================

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Step 1: Upgrade the connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("New client connected!")

	// Step 2: Read and echo messages in a loop
	// YOUR CODE HERE:

}

func main() {
	// Register the WebSocket handler
	http.HandleFunc("/ws", handleWebSocket)

	// Serve the test page
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Open http://localhost:8080 to test the WebSocket")
	http.ListenAndServe(":8080", nil)
}

// ==========================================
// TO RUN:
//   1. In this folder, run:
//      go mod init chatserver
//      go get github.com/gorilla/websocket
//   2. Run: go run main.go
//   3. Open http://localhost:8080 in your browser
//   4. Type messages and see them echo back!
//   5. Watch the terminal for server-side logs
// ==========================================
