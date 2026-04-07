# Lesson 5: Clicks and Events

## What Are Events?

An **event** is something that happens on the page — a user clicks a button, types in a field, moves their mouse, presses a key, etc. JavaScript can **listen** for these events and run code when they happen.

This is what makes web pages interactive!

## Adding an Event Listener

The main way to respond to events is `addEventListener`:

```javascript
let button = document.getElementById("myButton");

button.addEventListener("click", function() {
  console.log("Button was clicked!");
});
```

Let's break this down:
1. **Find the element** — `getElementById("myButton")`
2. **Tell it to listen** — `.addEventListener()`
3. **What event** — `"click"` (first argument)
4. **What to do** — the `function() { ... }` (second argument). This is called a **callback function** — code that runs *when the event happens*, not immediately.

## Common Events

| Event | When It Fires |
|-------|--------------|
| `"click"` | User clicks the element |
| `"input"` | User types in an input/textarea (fires on every keystroke) |
| `"change"` | Input value changes and element loses focus |
| `"keydown"` | User presses a key |
| `"keyup"` | User releases a key |
| `"submit"` | A form is submitted |
| `"mouseover"` | Mouse moves over the element |
| `"mouseout"` | Mouse leaves the element |

## Practical Example: Click Counter

```html
<p>Count: <span id="count">0</span></p>
<button id="add-btn">Add One</button>

<script>
  let count = 0;
  let countDisplay = document.getElementById("count");
  let addButton = document.getElementById("add-btn");

  addButton.addEventListener("click", function() {
    count = count + 1;
    countDisplay.textContent = count;
  });
</script>
```

Each click increases the count and updates the display. This pattern — **listen for event, update data, update display** — is the foundation of interactive apps.

## Responding to Input

The `"input"` event fires every time someone types a character:

```html
<input type="text" id="name-input" placeholder="Type your name">
<p id="greeting">Hello, !</p>

<script>
  let nameInput = document.getElementById("name-input");
  let greeting = document.getElementById("greeting");

  nameInput.addEventListener("input", function() {
    greeting.textContent = "Hello, " + nameInput.value + "!";
  });
</script>
```

As the user types, the greeting updates in real-time!

## Keyboard Events

You can detect specific keys:

```javascript
let input = document.getElementById("chat-input");

input.addEventListener("keydown", function(event) {
  if (event.key === "Enter") {
    console.log("User pressed Enter!");
    console.log("They typed: " + input.value);
  }
});
```

The `event` parameter (you can call it anything, but `event` or `e` is conventional) contains details about what happened. `event.key` tells you which key was pressed.

This will be super important for our chat app — sending a message when the user presses Enter!

## Preventing Default Behavior

Some elements have built-in behaviors. For example, clicking a submit button refreshes the page. You can prevent this:

```javascript
form.addEventListener("submit", function(event) {
  event.preventDefault();  // stops the page from refreshing
  // your code here
});
```

## Putting It Together: A Mini App

Here's a small example that combines what we've learned:

```html
<input type="text" id="todo-input" placeholder="Add a task...">
<button id="add-btn">Add</button>
<ul id="todo-list"></ul>

<script>
  let input = document.getElementById("todo-input");
  let addBtn = document.getElementById("add-btn");
  let list = document.getElementById("todo-list");

  addBtn.addEventListener("click", function() {
    // Get the text from the input
    let taskText = input.value;

    // Create a new list item
    let li = document.createElement("li");
    li.textContent = taskText;

    // Add it to the list
    list.appendChild(li);

    // Clear the input
    input.value = "";
  });
</script>
```

Look at that — a working todo list in about 15 lines! This is very close to how our chat app will work.

## Try It!

Open `exercise.html` in this folder for hands-on practice with events.

## Key Takeaways

- **Events** are things that happen on the page (clicks, typing, key presses)
- `element.addEventListener("event", function() { ... })` listens for events
- The function inside is a **callback** — it runs when the event happens, not immediately
- `"click"` for button clicks, `"input"` for typing, `"keydown"` for key presses
- `event.key === "Enter"` detects the Enter key
- `event.preventDefault()` stops default browser behavior
- Pattern: **listen for event → update data → update the page**
