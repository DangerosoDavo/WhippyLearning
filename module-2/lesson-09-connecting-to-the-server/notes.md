# Lesson 9: Connecting to the Server

## Focusing on the Frontend

In the last two lessons, we built the Go server. Now we switch back to JavaScript and focus entirely on the **browser side** — connecting to the WebSocket server and understanding the connection lifecycle.

## The WebSocket API in JavaScript

JavaScript has a built-in `WebSocket` class. No libraries, no installs — it's just there:

```javascript
let ws = new WebSocket("ws://localhost:8080/ws");
```

That one line starts a connection to your Go server. The URL uses `ws://` instead of `http://` (just like `https://` has `wss://` for secure WebSockets).

## Connection Lifecycle

A WebSocket connection goes through stages, and JavaScript gives you **event callbacks** for each one:

```javascript
let ws = new WebSocket("ws://localhost:8080/ws");

// 1. Connection opened successfully
ws.onopen = function() {
    console.log("Connected to server!");
};

// 2. A message arrived from the server
ws.onmessage = function(event) {
    console.log("Got message:", event.data);
};

// 3. Connection was closed
ws.onclose = function(event) {
    console.log("Disconnected from server");
};

// 4. Something went wrong
ws.onerror = function(error) {
    console.log("WebSocket error:", error);
};
```

These should feel familiar — they're just like the event listeners from Module 1! The difference is you assign them with `=` instead of `addEventListener`.

### The Lifecycle Flow

```
1. new WebSocket("ws://...")  →  Connecting...
2. ws.onopen fires            →  Connected!
3. ws.onmessage fires         →  Messages arrive (can happen many times)
4. ws.onclose fires           →  Connection closed
```

`onerror` can fire at any point if something goes wrong.

## Checking Connection State

The `ws.readyState` property tells you the current state:

```javascript
WebSocket.CONNECTING  // 0 — Still connecting
WebSocket.OPEN        // 1 — Connected and ready
WebSocket.CLOSING     // 2 — Connection is closing
WebSocket.CLOSED      // 3 — Connection is closed
```

Use this before sending to avoid errors:

```javascript
if (ws.readyState === WebSocket.OPEN) {
    ws.send("Hello!");
} else {
    console.log("Not connected yet!");
}
```

## Building a Connection Status Indicator

A great UX feature is showing the user whether they're connected:

```javascript
let statusBar = document.getElementById("connection-status");

ws.onopen = function() {
    statusBar.textContent = "Connected";
    statusBar.style.backgroundColor = "#d4edda";
    statusBar.style.color = "#155724";
};

ws.onclose = function() {
    statusBar.textContent = "Disconnected — trying to reconnect...";
    statusBar.style.backgroundColor = "#f8d7da";
    statusBar.style.color = "#721c24";
};
```

Our chat app already has a `#connection-status` bar from Lesson 5 — now we'll make it actually work!

## Reconnection

WebSocket connections can drop (server restart, network hiccup). A good chat app tries to **reconnect automatically**:

```javascript
function connect() {
    let ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = function() {
        console.log("Connected!");
    };

    ws.onclose = function() {
        console.log("Disconnected. Reconnecting in 3 seconds...");
        setTimeout(function() {
            connect();  // Try again
        }, 3000);
    };

    // Store ws somewhere so other code can use it
    return ws;
}

let ws = connect();
```

When the connection closes, we wait 3 seconds and try again. The function calls itself — this is called **recursion** and it works perfectly here.

## What We're NOT Doing Yet

This lesson is just about the **connection**. We're not wiring up the send button or displaying messages yet — that's Lesson 10. For now, focus on:
- Opening the connection
- Handling the lifecycle events
- Showing connection status
- Auto-reconnecting

## Try It!

The exercise has the chat app frontend with spots to implement the WebSocket connection, status indicator, and reconnection logic. The server from Lesson 8 is included.

## Key Takeaways

- `new WebSocket("ws://localhost:8080/ws")` opens a connection
- Four lifecycle events: `onopen`, `onmessage`, `onclose`, `onerror`
- `ws.readyState === WebSocket.OPEN` checks if connected
- Always show the user whether they're connected or not
- Auto-reconnect by calling your connect function again after a delay
- Separate concerns: this lesson = connection; next lesson = messages
