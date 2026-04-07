# Lesson 7: What Are WebSockets?

## The Problem with HTTP

HTTP works great for loading web pages — you ask for a page, the server sends it, done. But for a **chat app**, HTTP has a big problem:

With HTTP, the **client** always has to ask. The server can't just *send* you a message — it has to wait for you to ask "any new messages?"

```
HTTP (Request/Response):

Bret: "Any new messages?"     → Server: "No."
Bret: "Any new messages?"     → Server: "No."
Bret: "Any new messages?"     → Server: "Yes! Dave said hi."
Bret: "Any new messages?"     → Server: "No."
...forever
```

This is called **polling** and it's wasteful — most of the time the answer is "no."

## WebSockets: A Persistent Connection

**WebSockets** solve this problem. Instead of repeatedly asking, the browser opens a **persistent connection** that stays open. Either side can send a message at any time.

```
WebSocket (Persistent Connection):

Browser ←→ Server    [connection stays open]

Server → Browser:    "Dave says: Hi!"        (instant)
Browser → Server:    "Bret says: Hey Dave!"   (instant)
Server → Browser:    "Dave says: What's up?"  (instant)
```

Think of HTTP like sending letters (you send one, wait for a reply). WebSockets are like a phone call (both sides talk whenever they want).

## How WebSockets Work

1. The browser sends a normal HTTP request saying "I'd like to upgrade to a WebSocket"
2. The server says "OK, let's upgrade"
3. The connection stays open — both sides can send messages freely
4. Either side can close the connection when done

The initial "upgrade" is called the **handshake**. After that, it's a two-way channel.

## The gorilla/websocket Package

Go's standard library doesn't include WebSocket support, so we'll use a very popular package called `gorilla/websocket`. This is our first **external dependency** — code someone else wrote that we import into our project.

### Installing It

In your project folder (where `go.mod` is), run:

```
go get github.com/gorilla/websocket
```

This downloads the package and adds it to your `go.mod` file.

### The Upgrader

The key piece from `gorilla/websocket` is the **Upgrader** — it handles the handshake that converts an HTTP connection to a WebSocket:

```go
import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true  // Allow connections from any origin
    },
}
```

### Upgrading a Connection

In your HTTP handler, you upgrade the connection:

```go
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    // Upgrade HTTP → WebSocket
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Upgrade error:", err)
        return
    }
    defer conn.Close()

    fmt.Println("New client connected!")

    // Now we can read and write messages on conn
}
```

**New concept: `defer`**

`defer conn.Close()` means "run `conn.Close()` when this function ends." It's Go's way of making sure you clean up resources. Think of it like saying "before I leave, remember to lock the door" — no matter how you exit the function, the deferred call runs.

### Reading Messages

```go
for {
    _, messageBytes, err := conn.ReadMessage()
    if err != nil {
        fmt.Println("Read error:", err)
        break
    }
    fmt.Println("Received:", string(messageBytes))
}
```

This is an **infinite loop** (`for` with no condition) that keeps reading messages until an error occurs (like the client disconnecting).

### Writing Messages

```go
message := []byte("Hello from the server!")
err := conn.WriteMessage(websocket.TextMessage, message)
if err != nil {
    fmt.Println("Write error:", err)
}
```

## The Full Server (Preview)

Here's what our WebSocket server will look like when it's done:

```go
http.HandleFunc("/ws", handleWebSocket)  // WebSocket route

fs := http.FileServer(http.Dir("./static"))
http.Handle("/", fs)                     // Static files

http.ListenAndServe(":8080", nil)
```

The browser connects to `ws://localhost:8080/ws` for WebSocket and loads the page from `http://localhost:8080/` for the HTML.

## What About the Frontend?

JavaScript has a built-in `WebSocket` class (no install needed!):

```javascript
let ws = new WebSocket("ws://localhost:8080/ws");

ws.onopen = function() {
    console.log("Connected!");
};

ws.onmessage = function(event) {
    console.log("Received:", event.data);
};

ws.send("Hello from the browser!");
```

We'll cover this in detail in Lessons 9-10.

## Try It!

This lesson is mostly conceptual, but the exercise has a basic echo server — it receives a message and sends it right back. You'll test it using the browser console.

## Key Takeaways

- **HTTP** = request/response (client always asks first)
- **WebSocket** = persistent two-way connection (either side sends anytime)
- WebSockets start with an HTTP "upgrade" handshake
- We use `gorilla/websocket` for WebSocket support in Go
- Install with: `go get github.com/gorilla/websocket`
- `upgrader.Upgrade()` converts HTTP to WebSocket
- `conn.ReadMessage()` and `conn.WriteMessage()` for sending/receiving
- `defer conn.Close()` ensures cleanup when the function ends
- JavaScript has built-in `WebSocket` support — no libraries needed
