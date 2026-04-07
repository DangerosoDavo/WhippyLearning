# Lesson 11: Broadcasting to Everyone

## The Server Gets Smarter

In Lesson 8 we wrote a basic broadcast function. Now we're going to enhance the **server** to handle multiple users properly. This lesson focuses on:

1. **Concurrency safety** — what happens when two users send messages at the exact same time
2. **Server-generated messages** — the server itself sending notifications
3. **User tracking** — knowing who's connected, not just their connection

## The Problem: Concurrent Access

Here's a subtle bug in our current server. Imagine:

- Bret sends a message → the server loops through `clients` to broadcast
- At the same moment, Dave disconnects → we try to `delete(clients, conn)`

Both things are happening **at the same time** in different goroutines, and both are touching the `clients` map. In Go, this can crash your program! Maps are not safe for concurrent access.

## The Solution: sync.Mutex

A **mutex** (mutual exclusion) is a lock. Only one goroutine can hold the lock at a time:

```go
import "sync"

var clients = make(map[*websocket.Conn]bool)
var mutex = sync.Mutex{}

// When modifying or reading the clients map:
mutex.Lock()
clients[conn] = true    // safe!
mutex.Unlock()
```

While one goroutine has the lock, any other goroutine that tries to `Lock()` will wait until the first one calls `Unlock()`.

Think of it like a bathroom lock — only one person at a time.

## Updated Broadcast Function

```go
func broadcast(msg Message) {
    mutex.Lock()
    defer mutex.Unlock()

    for conn := range clients {
        err := conn.WriteJSON(msg)
        if err != nil {
            conn.Close()
            delete(clients, conn)
        }
    }
}
```

`defer mutex.Unlock()` ensures the lock is released no matter what happens in the function — even if there's an error.

## Server-Generated Messages

The server can create and broadcast its own messages. This is great for notifications like "Bret joined the chat":

```go
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    defer conn.Close()

    // Add client
    mutex.Lock()
    clients[conn] = true
    mutex.Unlock()

    // We'll learn who they are from their first message
    // (or we could add a "join" message type — see below)

    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            // Broadcast that someone left
            mutex.Lock()
            delete(clients, conn)
            mutex.Unlock()

            leaveMsg := Message{
                Sender: "System",
                Text:   "A user disconnected",
                Time:   currentTime(),
            }
            broadcast(leaveMsg)
            break
        }

        fmt.Printf("[%s] %s: %s\n", msg.Time, msg.Sender, msg.Text)
        broadcast(msg)
    }
}

func currentTime() string {
    return time.Now().Format("3:04 PM")
}
```

## Message Types

Right now, all messages are chat messages. But we might want different types — join notifications, leave notifications, system messages. We can add a `Type` field:

```go
type Message struct {
    Type   string `json:"type"`
    Sender string `json:"sender"`
    Text   string `json:"text"`
    Time   string `json:"time"`
}
```

Then on the server:
```go
// User joined
joinMsg := Message{
    Type:   "system",
    Sender: "System",
    Text:   msg.Sender + " joined the chat",
    Time:   currentTime(),
}
broadcast(joinMsg)
```

And on the frontend, you can style system messages differently:
```javascript
ws.onmessage = function(event) {
    let message = JSON.parse(event.data);

    if (message.type === "system") {
        addSystemMessage(message.text);
    } else {
        displayMessage(message);
    }
};
```

## User Count

Let's show how many people are connected:

```go
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    // ... after adding client:

    // Tell everyone the new count
    countMsg := Message{
        Type:   "usercount",
        Text:   fmt.Sprintf("%d", len(clients)),
        Time:   currentTime(),
    }
    broadcast(countMsg)

    // ... in the disconnect section:
    countMsg := Message{
        Type:   "usercount",
        Text:   fmt.Sprintf("%d", len(clients)),
        Time:   currentTime(),
    }
    broadcast(countMsg)
}
```

Frontend:
```javascript
if (message.type === "usercount") {
    document.getElementById("user-count").textContent = message.text + " online";
}
```

## Try It!

The exercise enhances the server with a mutex, system messages, and message types. The frontend is updated to handle the different message types.

## Key Takeaways

- **Concurrent access** to maps is unsafe — use `sync.Mutex` to protect them
- `mutex.Lock()` / `defer mutex.Unlock()` — lock before, unlock after
- The server can **create its own messages** (join/leave notifications, user counts)
- Adding a `Type` field lets you handle different kinds of messages
- Always use `defer` with `Unlock()` to prevent forgetting to release the lock
