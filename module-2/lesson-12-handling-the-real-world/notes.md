# Lesson 12: Handling the Real World

## Making It Robust

Our chat app works — but real-world apps need to handle things that go wrong. In this final lesson, we'll add:

1. **Graceful disconnection handling** on both sides
2. **Input validation** on the server
3. **Reconnection with message history** on the frontend
4. **A "typing" indicator** for fun
5. The **complete, finished chat app**

## Server-Side Validation

Never trust data from the client. What if someone sends an empty message? Or a message with 10,000 characters?

```go
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    // ... connection setup ...

    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            // handle disconnect
            break
        }

        // Validate!
        msg.Text = strings.TrimSpace(msg.Text)
        if msg.Text == "" {
            continue  // skip empty messages
        }
        if len(msg.Text) > 500 {
            msg.Text = msg.Text[:500]  // truncate long messages
        }
        if msg.Sender == "" {
            msg.Sender = "Anonymous"
        }

        // Set the time on the server (don't trust the client's clock)
        msg.Time = currentTime()

        msg.Type = "chat"
        broadcast(msg)
    }
}
```

Key practices:
- **Trim whitespace** — don't send blank messages
- **Limit length** — prevent abuse
- **Default values** — handle missing fields
- **Server-set timestamps** — the server's clock is the source of truth

## Reconnection That Works

Our current reconnect logic has a problem: when you reconnect, you lose all previous messages. Let's fix that by keeping a message history on the server and sending it to new connections:

```go
var messageHistory []Message
const maxHistory = 50

func broadcast(msg Message) {
    mutex.Lock()
    defer mutex.Unlock()

    // Save to history
    messageHistory = append(messageHistory, msg)
    if len(messageHistory) > maxHistory {
        messageHistory = messageHistory[1:]  // remove oldest
    }

    // Send to all clients
    for conn := range clients {
        err := conn.WriteJSON(msg)
        if err != nil {
            conn.Close()
            delete(clients, conn)
        }
    }
}
```

When a new client connects, send them the history:

```go
// After registering the client:
mutex.Lock()
for _, msg := range messageHistory {
    conn.WriteJSON(msg)
}
mutex.Unlock()
```

Now when someone reconnects (or joins late), they see the last 50 messages!

## Frontend: Clearing Chat on Reconnect

When reconnecting, clear old messages first to avoid duplicates (the server sends history):

```javascript
ws.onopen = function() {
    updateStatus("connected");
    chatMessages.innerHTML = "";  // clear old messages
    addSystemMessage("Connected to server!");
};
```

## Typing Indicators

A fun feature: show when someone is typing. This uses a new message type:

### Frontend:
```javascript
let typingTimeout = null;

messageInput.addEventListener("input", function() {
    if (ws.readyState !== WebSocket.OPEN) return;

    // Send "typing" notification
    ws.send(JSON.stringify({
        type: "typing",
        sender: usernameInput.value.trim() || "Anonymous"
    }));

    // Clear previous timeout
    clearTimeout(typingTimeout);

    // After 2 seconds of no typing, send "stopped"
    typingTimeout = setTimeout(function() {
        ws.send(JSON.stringify({
            type: "stopped_typing",
            sender: usernameInput.value.trim() || "Anonymous"
        }));
    }, 2000);
});
```

### Server:
The server broadcasts typing notifications to everyone except the sender:

```go
if msg.Type == "typing" || msg.Type == "stopped_typing" {
    mutex.Lock()
    for c := range clients {
        if c != conn {  // don't send back to the typer
            c.WriteJSON(msg)
        }
    }
    mutex.Unlock()
    continue  // don't save typing events to history
}
```

### Frontend display:
```javascript
let typingDisplay = document.getElementById("typing-indicator");

if (message.type === "typing") {
    typingDisplay.textContent = message.sender + " is typing...";
} else if (message.type === "stopped_typing") {
    typingDisplay.textContent = "";
}
```

## The Complete Application

The exercise for this lesson is the **finished chat app** — both server and frontend, with everything we've built across all 12 lessons. It includes:

- Go WebSocket server with broadcasting
- Mutex-protected concurrent access
- Message validation and server-side timestamps
- Message history (last 50 messages)
- Join/leave notifications with user count
- Auto-reconnection with history reload
- Typing indicators
- Clean, styled UI from Module 1

## Testing It Properly

1. Start the server: `go run main.go`
2. Open `http://localhost:8080` in **Tab 1** (set name to "Bret")
3. Open `http://localhost:8080` in **Tab 2** (set name to "Dave")
4. Send messages back and forth — they appear in both tabs instantly
5. Stop the server (`Ctrl+C`) — both tabs show "Disconnected"
6. Start the server again — both tabs auto-reconnect
7. Open a **third tab** — it should see the recent message history
8. Close a tab — the others see a "left the chat" notification

## Congratulations!

You've built a **real-time, multi-user chat application** with:
- A **Go backend** that handles WebSocket connections
- A **JavaScript frontend** that connects and displays messages
- **JSON** as the data format between them
- **Broadcasting** so everyone sees every message
- **Production-quality** features like validation, reconnection, and history

### What You've Learned in Module 2

| Lesson | Skills |
|--------|--------|
| 1 | Client-server model, HTTP, ports |
| 2 | Go basics — variables, types, printing |
| 3 | Go functions, loops, structs, slices |
| 4 | Building an HTTP server in Go |
| 5 | Serving static files |
| 6 | JSON encoding/decoding in Go and JS |
| 7 | WebSocket concepts, gorilla/websocket |
| 8 | WebSocket server, goroutines, broadcast |
| 9 | JS WebSocket API, connection lifecycle |
| 10 | Sending/receiving JSON over WebSocket |
| 11 | Concurrency safety, message types, user tracking |
| 12 | Validation, history, typing indicators, polish |

## Where to Go Next

- **Databases** — Store messages permanently (SQLite, PostgreSQL)
- **Authentication** — User accounts and passwords
- **Multiple rooms/channels** — Like Discord or Slack
- **File sharing** — Send images and files
- **Deploy** — Put your server on the internet for anyone to use
- **React/Vue frontend** — Build the UI with a modern framework

You've proven you can learn any of these. The fundamentals you've built here — client-server communication, data formats, real-time connections — are the same whether you're building a chat app, a game, or a full SaaS product.

Keep building!
