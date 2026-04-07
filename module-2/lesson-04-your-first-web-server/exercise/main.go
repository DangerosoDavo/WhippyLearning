package main

import (
	"fmt"
	"net/http"
)

// ==========================================
// EXERCISE 1: Basic handler
// Write a handler function called "homePage" that
// writes "Welcome to my server!" to the response.
//
// func homePage(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Welcome to my server!")
// }
// ==========================================

// YOUR CODE HERE:

// ==========================================
// EXERCISE 2: HTML handler
// Write a handler called "aboutPage" that sends
// back HTML (not just plain text).
// Include an <h1>, a <p>, and a <ul> with a few items.
//
// Don't forget to set the content type:
//   w.Header().Set("Content-Type", "text/html")
// ==========================================

// YOUR CODE HERE:

// ==========================================
// EXERCISE 3: Dynamic handler
// Write a handler called "greetPage" that reads
// a name from the URL query string and greets them.
//
// The URL will be: /greet?name=Bret
// Read it with: r.URL.Query().Get("name")
//
// If no name is provided, greet "stranger".
//
// Example:
//   func greetPage(w http.ResponseWriter, r *http.Request) {
//       name := r.URL.Query().Get("name")
//       if name == "" {
//           name = "stranger"
//       }
//       fmt.Fprintf(w, "Hello, %s!", name)
//   }
// ==========================================

// YOUR CODE HERE:

func main() {
	// ==========================================
	// EXERCISE 4: Register your routes
	// Use http.HandleFunc to connect:
	//   "/" -> homePage
	//   "/about" -> aboutPage
	//   "/greet" -> greetPage
	//
	// Then start the server on port 8080.
	// ==========================================

	// YOUR CODE HERE:

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// ==========================================
// TO RUN:
//   1. Open terminal in this folder
//   2. Run: go mod init chatserver
//      (only needed the first time)
//   3. Run: go run main.go
//   4. Open browser to http://localhost:8080
//   5. Try http://localhost:8080/about
//   6. Try http://localhost:8080/greet?name=Bret
//   7. Press Ctrl+C in terminal to stop
// ==========================================
