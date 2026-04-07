# Lesson 8: Building the Chat UI

## Time to Build!

Over the last 7 lessons, you've learned everything you need:
- **HTML** to create structure
- **CSS** to make it look good
- **JavaScript** to make it interactive
- **Events** to respond to user actions
- **Functions** to organize code
- **Arrays and objects** to manage data

Now we're putting it all together to build the **user interface** (UI) of our chat app. In this lesson, we focus on the HTML and CSS — making it look like a real chat app.

## What Does a Chat App Look Like?

Think about any chat app you've used (iMessage, WhatsApp, Discord). They all have:

1. A **header** — shows the app name or who you're chatting with
2. A **message area** — where messages appear (scrollable)
3. An **input area** — where you type and send messages

## The HTML Structure

```html
<div id="chat-app">
  <!-- Header -->
  <div id="chat-header">
    <h1>ChatApp</h1>
  </div>

  <!-- Messages Area -->
  <div id="chat-messages">
    <!-- Messages will be added here by JavaScript -->
  </div>

  <!-- Input Area -->
  <div id="chat-input-area">
    <input type="text" id="message-input" placeholder="Type a message...">
    <button id="send-btn">Send</button>
  </div>
</div>
```

That's it! The structure is simple. The CSS is what makes it look good.

## Key CSS Concepts for the Chat Layout

### Flexbox: Arranging Things in a Line

**Flexbox** is a CSS layout system that makes arranging elements easy. We'll use it for:
- Making the whole app a vertical column (header, messages, input)
- Making the input area a horizontal row (text input + button)

```css
#chat-app {
  display: flex;
  flex-direction: column;  /* stack children vertically */
  height: 100vh;           /* fill the full screen height */
}
```

- `display: flex` — turns on flexbox
- `flex-direction: column` — stack children top to bottom
- `height: 100vh` — `vh` means "viewport height" (100vh = full screen)

For the input area:
```css
#chat-input-area {
  display: flex;
  flex-direction: row;  /* place children side by side */
}
```

### Making the Message Area Scrollable

```css
#chat-messages {
  flex: 1;            /* take up all available space */
  overflow-y: auto;   /* add scrollbar when content overflows */
}
```

- `flex: 1` — "grow to fill whatever space is left"
- `overflow-y: auto` — "if there's more content than fits, add a vertical scrollbar"

### Making the Input Stretch

```css
#message-input {
  flex: 1;  /* take up all available width */
}
```

## Styling Messages

Messages in chat apps usually look like speech bubbles:

```css
.message {
  max-width: 70%;
  padding: 10px 15px;
  margin: 5px 10px;
  border-radius: 18px;
  word-wrap: break-word;
}
```

- `max-width: 70%` — messages don't stretch the full width
- `border-radius: 18px` — rounded bubble shape
- `word-wrap: break-word` — long words wrap instead of overflowing

## Try It!

Open `exercise.html` — it has a partially built chat UI. Your job is to complete the CSS to make it look like a real chat app. The instructions guide you step by step.

## Key Takeaways

- A chat app has three sections: **header**, **messages area**, **input area**
- **Flexbox** (`display: flex`) arranges elements in rows or columns
- `flex: 1` makes an element grow to fill available space
- `overflow-y: auto` makes a container scrollable
- `100vh` = full screen height
- Good CSS turns simple HTML into a professional-looking interface
