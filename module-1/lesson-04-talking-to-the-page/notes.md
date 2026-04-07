# Lesson 4: Talking to the Page (The DOM)

## The Big Idea

So far, our JavaScript has only talked to the **console**. But the real power of JavaScript is that it can **read and change what's on the page** — add text, change colors, show and hide things, and more.

## What is the DOM?

When your browser loads an HTML page, it creates a behind-the-scenes structure called the **DOM** (Document Object Model). Think of it as a tree of all the elements on your page that JavaScript can interact with.

You don't need to fully understand the DOM right now. Just know this: **JavaScript can find any element on your page and change it.**

## Finding Elements

Before you can change something, you need to **find** it. The most common way is `document.getElementById()`:

```html
<p id="greeting">Hello!</p>

<script>
  let greetingElement = document.getElementById("greeting");
  console.log(greetingElement);  // shows the <p> element
</script>
```

- `document` — represents the entire page
- `.getElementById("greeting")` — finds the element with `id="greeting"`
- The result is stored in a variable so you can work with it

### Other Ways to Find Elements

```javascript
// Find by class name (returns a list of all matches)
let boxes = document.getElementsByClassName("box");

// Find using CSS selector syntax (returns the FIRST match)
let firstBox = document.querySelector(".box");

// Find ALL matches using CSS selector syntax
let allBoxes = document.querySelectorAll(".box");
```

`querySelector` is the most flexible — it uses the same selectors as CSS:
- `"#myId"` — find by ID
- `".myClass"` — find by class
- `"p"` — find by tag name
- `"div.box"` — find a div with class "box"

## Changing Text Content

Once you've found an element, you can change its text:

```html
<p id="message">Original text</p>

<script>
  let msg = document.getElementById("message");
  msg.textContent = "New text!";
</script>
```

When this runs, the paragraph will show "New text!" instead of "Original text".

### textContent vs innerHTML

- `textContent` — sets plain text (safe, use this by default)
- `innerHTML` — sets HTML (can include tags like `<strong>`)

```javascript
msg.textContent = "<strong>Bold?</strong>";
// Shows literally: <strong>Bold?</strong>

msg.innerHTML = "<strong>Bold!</strong>";
// Shows: Bold!  (actually bold)
```

## Changing Styles

You can change CSS styles directly from JavaScript:

```javascript
let title = document.getElementById("title");
title.style.color = "red";
title.style.fontSize = "48px";
title.style.backgroundColor = "yellow";
```

**Note**: CSS properties that use hyphens (`font-size`, `background-color`) become **camelCase** in JavaScript (`fontSize`, `backgroundColor`).

## Changing Classes

Instead of changing individual styles, you can add/remove CSS classes:

```javascript
let box = document.getElementById("mybox");
box.classList.add("highlight");     // adds a class
box.classList.remove("highlight");  // removes a class
box.classList.toggle("highlight");  // adds if missing, removes if present
```

This is usually better than changing `style` directly because it keeps your styles in CSS where they belong.

## Creating New Elements

You can create entirely new HTML elements with JavaScript:

```javascript
// Step 1: Create the element
let newParagraph = document.createElement("p");

// Step 2: Give it content
newParagraph.textContent = "I was created by JavaScript!";

// Step 3: Add it to the page
document.body.appendChild(newParagraph);
```

`appendChild` adds the new element as the **last child** of whatever element you call it on. In this case, it gets added to the end of `<body>`.

## Getting Values from Input Fields

This is crucial for our chat app later! You can read what a user types into an `<input>`:

```html
<input type="text" id="nameInput">

<script>
  let input = document.getElementById("nameInput");
  let whatTheyTyped = input.value;
  console.log(whatTheyTyped);
</script>
```

The `.value` property gives you the current text inside the input field.

## Try It!

Open `exercise.html` in this folder. This time, you'll see changes happen right on the page, not just in the console!

## Key Takeaways

- The **DOM** is how JavaScript sees and interacts with your page
- `document.getElementById("id")` finds an element by its ID
- `document.querySelector(".class")` finds elements using CSS selectors
- `.textContent` changes the text inside an element
- `.style.propertyName` changes CSS styles from JavaScript
- `.classList.add/remove/toggle` manages CSS classes
- `document.createElement()` + `.appendChild()` creates new elements
- `.value` reads what's typed in an input field
