package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ==========================================
	// EXERCISE 1: Serve static files
	// Create a file server that serves from the
	// "./static" folder and register it on "/".
	//
	// Hint:
	//   fs := http.FileServer(http.Dir("./static"))
	//   http.Handle("/", fs)
	// ==========================================

	// YOUR CODE HERE:

	// ==========================================
	// EXERCISE 2: Add a status route
	// Before the file server, add a route at
	// "/api/status" that responds with the text:
	// "Chat server is running!"
	//
	// This is the start of our server API.
	// We'll add a WebSocket route here later.
	// ==========================================

	// YOUR CODE HERE:

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Open http://localhost:8080 in your browser")
	http.ListenAndServe(":8080", nil)
}

// ==========================================
// TO RUN:
//   1. Open terminal in this folder
//   2. Run: go mod init chatserver
//   3. Run: go run main.go
//   4. Open http://localhost:8080 in your browser
//   5. You should see the chat app!
//   6. Try http://localhost:8080/api/status
//   7. Ctrl+C to stop
// ==========================================
