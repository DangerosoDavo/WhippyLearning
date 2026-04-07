# Lesson 10: Sending & Receiving Messages

## Wiring Up the Chat

In Lesson 9, we established the WebSocket connection. Now we'll:
1. **Send** messages from the browser to the server as JSON
2. **Receive** messages from the server and display them in the chat

By the end of this lesson, you'll be able to chat with yourself across two browser tabs!

## Sending a Message

Remember Module 1's `sendMessage()` function? It created a message object and displayed it locally. Now we need to **send it to the server** instead:

```javascript
function sendMessage() {
    let text = messageInput.value.trim();
    let sender = usernameInput.value.trim() || "Anonymous";

    if (!text) return;
    if (ws.readyState !== WebSocket.OPEN) return;

    // Create the message object
    let message = {
        sender: sender,
        text: text,
        time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
    };

    // Convert to JSON and send over WebSocket
    ws.send(JSON.stringify(message));

    // Clear the input
    messageInput.value = "";
    messageInput.focus();
}
```

Key changes from Module 1:
- We check `ws.readyState` before sending (can't send if disconnected)
- We use `JSON.stringify(message)` to convert the object to a JSON string
- We use `ws.send()` to send it over the WebSocket
- We **don't** display the message ourselves — the server will broadcast it back to us!

### Why Not Display It Locally?

You might wonder: "Why not just display it immediately and also send it?" The reason is that the server broadcasts messages to **everyone**, including the sender. If we displayed it locally AND received it from the server, we'd see duplicates.

By letting the server be the single source of truth, every message goes through the same path: send → server → broadcast → display. This keeps things consistent.

## Receiving a Message

When the server broadcasts a message, our `ws.onmessage` fires. We need to:
1. Parse the JSON string back into an object
2. Display it in the chat

```javascript
ws.onmessage = function(event) {
    // Parse the JSON string into an object
    let message = JSON.parse(event.data);

    // Display it in the chat
    displayMessage(message);
};
```

`event.data` contains the raw text the server sent. Since we know it's JSON, we parse it with `JSON.parse()`.

## The Display Function

This is the same `displayMessage` function from Module 1, updated to work with our chat:

```javascript
function displayMessage(message) {
    let currentUser = usernameInput.value.trim() || "Anonymous";
    let isSent = (message.sender === currentUser);

    // Sender name
    let senderDiv = document.createElement("div");
    senderDiv.classList.add("sender-name");
    senderDiv.textContent = message.sender;
    if (isSent) senderDiv.classList.add("sent-name");

    // Message bubble
    let messageDiv = document.createElement("div");
    messageDiv.classList.add("message");
    messageDiv.textContent = message.text;
    messageDiv.classList.add(isSent ? "sent" : "received");

    // Add to chat
    chatMessages.appendChild(senderDiv);
    chatMessages.appendChild(messageDiv);

    // Auto-scroll
    chatMessages.scrollTop = chatMessages.scrollHeight;
}
```

The `isSent` check determines if this is a message **we** sent or one from someone else, and applies the right CSS class.

## Hooking Up the Button and Enter Key

Same as Module 1:

```javascript
sendBtn.addEventListener("click", function() {
    sendMessage();
});

messageInput.addEventListener("keydown", function(event) {
    if (event.key === "Enter") {
        sendMessage();
    }
});
```

## The Complete Flow

Let's trace what happens when Bret sends "Hello!":

```
1. Bret types "Hello!" and presses Enter
2. sendMessage() creates: {sender: "Bret", text: "Hello!", time: "2:30 PM"}
3. JSON.stringify converts it to: '{"sender":"Bret","text":"Hello!","time":"2:30 PM"}'
4. ws.send() sends the JSON string to the Go server
5. Go server receives it with conn.ReadJSON()
6. Go server calls broadcast() — sends to ALL connected clients
7. Bret's browser receives the message via ws.onmessage
8. JSON.parse converts it back to a JavaScript object
9. displayMessage() creates the HTML and adds it to the chat
10. Dave's browser also receives and displays the same message
```

## Testing with Two Tabs

This is where it gets exciting! After completing the exercise:

1. Start the server: `go run main.go`
2. Open `http://localhost:8080` in a browser tab
3. Open `http://localhost:8080` in a **second** tab
4. Change the username in one tab to "Dave"
5. Send messages from both tabs
6. Watch them appear in **both** tabs in real-time!

## Try It!

Open `exercise/static/index.html` — it builds on Lesson 9's connection code. Complete the exercises to add sending and receiving.

## Key Takeaways

- **Sending**: `ws.send(JSON.stringify(message))` — convert to JSON and send
- **Receiving**: `JSON.parse(event.data)` in `ws.onmessage` — parse and display
- Don't display messages locally — let the server broadcast them back to you
- The server is the **single source of truth** for all messages
- Test with **two browser tabs** to see real-time chat in action!
