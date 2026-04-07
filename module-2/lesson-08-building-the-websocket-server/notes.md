# Lesson 8: Building the WebSocket Server

## From Echo to Chat

In the last lesson, we built an **echo server** — it receives a message and sends it back to the same person. But a chat server needs to do something more: **receive a message from one user and send it to everyone else.**

In this lesson, we focus on the **Go server side**. We'll build the server that can:
1. Accept WebSocket connections
2. Receive JSON messages
3. Store connected clients
4. Prepare for broadcasting (next lessons)

## The Message Struct

First, let's define what a chat message looks like in Go:

```go
type Message struct {
    Sender string `json:"sender"`
    Text   string `json:"text"`
    Time   string `json:"time"`
}
```

This matches our JavaScript message objects exactly, thanks to the JSON tags.

## Goroutines: Handling Multiple Users

Here's a challenge: when Bret connects, our server sits in a loop reading his messages. If Dave connects, we need to read *his* messages too — at the same time. How do we do two things at once?

Go has an amazing feature called **goroutines** — lightweight threads that let you run functions concurrently (at the same time).

```go
// Normal function call (blocks, waits until done)
doSomething()

// Goroutine (runs in the background, doesn't block)
go doSomething()
```

Just add `go` in front! The function runs in the background while the rest of your code continues.

### Why This Matters for Chat

```go
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    defer conn.Close()

    // This loop reads messages from ONE client
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
        fmt.Println("Got:", string(msg))
    }
}
```

When `handleWebSocket` is called, Go automatically runs it in its own goroutine for each connection. So if Bret and Dave both connect, there are two goroutines running — one reading Bret's messages, one reading Dave's. Go handles this for you!

## Tracking Connected Clients

To broadcast messages, we need to know who's connected. We'll use a **map** (Go's version of a dictionary/lookup table):

```go
// A map from WebSocket connection → true/false
var clients = make(map[*websocket.Conn]bool)
```

When someone connects, we add them:
```go
clients[conn] = true
```

When they disconnect, we remove them:
```go
delete(clients, conn)
```

## Reading JSON Messages

Instead of reading raw text, we'll read and parse JSON:

```go
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    defer conn.Close()

    // Register this client
    clients[conn] = true
    fmt.Println("Client connected. Total:", len(clients))

    // Read messages
    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            fmt.Println("Client disconnected")
            delete(clients, conn)
            break
        }
        fmt.Printf("From %s: %s\n", msg.Sender, msg.Text)
    }
}
```

`conn.ReadJSON(&msg)` is a shortcut that reads a message AND parses the JSON into your struct — two steps in one!

Similarly, `conn.WriteJSON(msg)` converts a struct to JSON and sends it.

## The Broadcast Function

Here's the heart of the chat server — sending a message to **every** connected client:

```go
func broadcast(msg Message) {
    for conn := range clients {
        err := conn.WriteJSON(msg)
        if err != nil {
            fmt.Println("Write error:", err)
            conn.Close()
            delete(clients, conn)
        }
    }
}
```

This loops through every connected client and sends them the message. If writing fails (they disconnected), we clean up.

## The Complete Server

Here's everything together:

```go
package main

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
```

That's the **entire backend** for a chat server — about 50 lines of Go!

## Try It!

Open `exercise/main.go` — build the server step by step. The test page from Lesson 7 is included so you can test it.

## Key Takeaways

- `conn.ReadJSON(&msg)` reads a WebSocket message and parses JSON in one step
- `conn.WriteJSON(msg)` converts to JSON and sends in one step
- **Goroutines** (`go functionName()`) let Go handle multiple connections at once
- A **map** (`map[*websocket.Conn]bool`) tracks connected clients
- **Broadcasting** = looping through all clients and sending the message to each
- The complete chat server is about 50 lines of Go!
