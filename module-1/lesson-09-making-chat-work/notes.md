# Lesson 9: Making Chat Work

## The Moment of Truth!

In the last lesson, we built the **look** of our chat app. Now we add the **brain** — JavaScript that actually sends and displays messages. By the end of this lesson, you'll have a working chat app!

## The Plan

Here's what our JavaScript needs to do:

1. Keep a **list of messages** (array of objects)
2. When the user clicks Send (or presses Enter):
   a. **Read** the message from the input
   b. **Validate** it (don't send empty messages)
   c. **Add** it to our messages array
   d. **Display** it on the page
   e. **Clear** the input
   f. **Scroll** to the bottom

We've practiced every single one of these skills in earlier lessons!

## Step 1: Set Up the Data

```javascript
// Our list of messages (starts empty)
let messages = [];

// The current user's name
let currentUser = "Bret";
```

Each message will be an object:
```javascript
{
  sender: "Bret",
  text: "Hello!",
  time: "2:30 PM"
}
```

## Step 2: Get References to Elements

```javascript
let chatMessages = document.getElementById("chat-messages");
let messageInput = document.getElementById("message-input");
let sendBtn = document.getElementById("send-btn");
```

## Step 3: Create the Send Function

```javascript
function sendMessage() {
  // Get the text from the input
  let text = messageInput.value.trim();

  // Don't send empty messages
  if (!text) {
    return;  // "return" exits the function immediately
  }

  // Create a message object
  let message = {
    sender: currentUser,
    text: text,
    time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
  };

  // Add to our array
  messages.push(message);

  // Display it
  displayMessage(message);

  // Clear the input
  messageInput.value = "";

  // Scroll to the bottom
  chatMessages.scrollTop = chatMessages.scrollHeight;
}
```

### New Things Here:

- `.trim()` — removes extra spaces from the start and end of a string. `"  hello  ".trim()` becomes `"hello"`. `"   ".trim()` becomes `""` (empty).
- `return;` — exits the function immediately, skipping the rest of the code
- `new Date().toLocaleTimeString(...)` — gets the current time as a readable string like "2:30 PM"
- `.scrollTop = .scrollHeight` — scrolls a container to the very bottom

## Step 4: Display a Message

```javascript
function displayMessage(message) {
  // Create the sender name label
  let senderDiv = document.createElement("div");
  senderDiv.classList.add("sender-name");
  senderDiv.textContent = message.sender;

  // Create the message bubble
  let messageDiv = document.createElement("div");
  messageDiv.classList.add("message");
  messageDiv.textContent = message.text;

  // Add "sent" or "received" class for styling
  if (message.sender === currentUser) {
    messageDiv.classList.add("sent");
  } else {
    messageDiv.classList.add("received");
  }

  // Add to the chat
  chatMessages.appendChild(senderDiv);
  chatMessages.appendChild(messageDiv);
}
```

This function takes a message object and creates the HTML elements to display it. It uses `classList.add` to apply the right styling — purple for sent messages, gray for received.

## Step 5: Wire Up the Events

```javascript
// Click the Send button
sendBtn.addEventListener("click", function() {
  sendMessage();
});

// Press Enter in the input
messageInput.addEventListener("keydown", function(event) {
  if (event.key === "Enter") {
    sendMessage();
  }
});
```

## That's It!

The complete JavaScript is about 50 lines. Each piece uses a concept you learned in a previous lesson:

| Concept | Where It's Used |
|---------|----------------|
| Variables | Storing messages, user name, element references |
| DOM (getElementById) | Finding the chat elements |
| Events (addEventListener) | Click and Enter key detection |
| Functions | sendMessage(), displayMessage() |
| If statements | Validating input, sent vs received styling |
| Arrays | Storing the messages list |
| Objects | Each message's data (sender, text, time) |
| createElement + appendChild | Building message bubbles |

## Try It!

Open `exercise.html` — it has the complete UI from Lesson 8 with guided steps to add the JavaScript. Follow the numbered comments.

## Key Takeaways

- `.trim()` removes whitespace from strings — great for input validation
- `return;` exits a function early (useful for validation)
- Break your code into focused functions (sendMessage, displayMessage)
- The pattern: **get input → validate → update data → update display → cleanup**
- `scrollTop = scrollHeight` auto-scrolls to the bottom
