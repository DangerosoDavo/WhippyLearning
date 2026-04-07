package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ==========================================
// EXERCISE 1: Define a Message struct
// Create a struct called "Message" with:
//   - Sender (string) with json tag "sender"
//   - Text (string) with json tag "text"
//   - Time (string) with json tag "time"
//
// Example:
//   type Message struct {
//       Sender string `json:"sender"`
//       ...
//   }
// ==========================================

// YOUR CODE HERE:

func main() {
	// ==========================================
	// EXERCISE 2: Struct to JSON
	// Create a Message struct with some values,
	// convert it to JSON using json.Marshal,
	// and print the result.
	//
	// msg := Message{Sender: "Bret", Text: "Hello!", Time: "2:30 PM"}
	// jsonBytes, err := json.Marshal(msg)
	// if err != nil { fmt.Println("Error:", err); return }
	// fmt.Println(string(jsonBytes))
	// ==========================================

	fmt.Println("--- Struct to JSON ---")
	// YOUR CODE HERE:

	// ==========================================
	// EXERCISE 3: JSON to Struct
	// Take this JSON string and convert it to a
	// Message struct using json.Unmarshal.
	// Print the Sender and Text fields.
	//
	// jsonStr := `{"sender":"Dave","text":"How are you?","time":"3:00 PM"}`
	// var received Message
	// err := json.Unmarshal([]byte(jsonStr), &received)
	// ...
	// fmt.Println("From:", received.Sender)
	// fmt.Println("Says:", received.Text)
	// ==========================================

	fmt.Println("--- JSON to Struct ---")
	// YOUR CODE HERE:

	// ==========================================
	// EXERCISE 4: JSON array
	// Create a slice of 3 Message structs.
	// Convert the whole slice to JSON and print it.
	//
	// messages := []Message{
	//     {Sender: "Bret", Text: "Hey!", Time: "2:30 PM"},
	//     ...
	// }
	// ==========================================

	fmt.Println("--- JSON Array ---")
	// YOUR CODE HERE:

	// ==========================================
	// EXERCISE 5: JSON API endpoint
	// Create a web server with a route "/api/messages"
	// that returns a JSON array of messages.
	//
	// In your handler:
	//   1. Set the content type to JSON:
	//      w.Header().Set("Content-Type", "application/json")
	//   2. Create a slice of Messages
	//   3. Use json.NewEncoder(w).Encode(messages)
	//      to write JSON directly to the response
	//
	// Then visit http://localhost:8080/api/messages
	// in your browser to see the JSON!
	// ==========================================

	http.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// YOUR CODE HERE:
		// Create messages slice and encode it

	})

	fmt.Println("\nServer starting on http://localhost:8080")
	fmt.Println("Try: http://localhost:8080/api/messages")
	http.ListenAndServe(":8080", nil)
}
