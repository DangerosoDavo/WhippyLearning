# Lesson 3: Hello JavaScript

## What is JavaScript?

So far we've learned:
- **HTML** = what's on the page (structure)
- **CSS** = how it looks (style)

Now we add the third piece: **JavaScript (JS)** = what it **does** (behavior).

JavaScript makes your page **interactive**. It can respond to clicks, change content, do calculations, validate forms, and much more. Unlike HTML and CSS, JavaScript is a real **programming language** — it can make decisions, repeat actions, and remember information.

## Adding JavaScript to Your Page

Just like CSS, you put JavaScript inside a special tag in your HTML. Use `<script>`:

```html
<!DOCTYPE html>
<html>
  <head>
    <title>My JS Page</title>
  </head>
  <body>
    <h1>Hello!</h1>

    <script>
      // JavaScript code goes here
    </script>
  </body>
</html>
```

**Important**: Put your `<script>` tag at the **bottom** of the `<body>`, right before `</body>`. This ensures the HTML has loaded before your JavaScript tries to interact with it.

## The Console: Your Best Friend

The **console** is a hidden panel in your browser where JavaScript can print messages. It's incredibly useful for testing and debugging.

**How to open it:**
1. Right-click anywhere on the page
2. Click "Inspect" (or "Inspect Element")
3. Click the "Console" tab

To print something to the console:

```javascript
console.log("Hello, world!");
```

`console.log()` is like "print this message so I can see it." You'll use it constantly.

## Variables: Storing Information

A **variable** is a named container that stores a value. Think of it like a labeled box.

```javascript
let name = "Bret";
let age = 25;
let isLearning = true;
```

- `let` — the keyword that creates a variable
- `name` — the name you choose for the variable (the label on the box)
- `=` — means "store this value" (not "equals" in the math sense)
- `"Bret"` — the value stored in the variable

You can then use the variable by its name:

```javascript
let name = "Bret";
console.log(name);        // prints: Bret
console.log("Hi, " + name);  // prints: Hi, Bret
```

### Changing a Variable

After creating a variable with `let`, you can change its value (just don't use `let` again):

```javascript
let score = 0;
score = 10;       // now it's 10
score = score + 5; // now it's 15
```

### const: Variables That Don't Change

If you have a value that should **never change**, use `const` instead of `let`:

```javascript
const pi = 3.14159;
const appName = "ChatApp";
```

If you try to change a `const`, you'll get an error. Use `const` by default, and only use `let` when you know the value will change.

## Data Types

JavaScript has a few basic types of data:

| Type | What It Is | Examples |
|------|-----------|----------|
| **String** | Text (always in quotes) | `"hello"`, `'world'`, `"123"` |
| **Number** | A number (no quotes) | `42`, `3.14`, `-7` |
| **Boolean** | True or false | `true`, `false` |

```javascript
let greeting = "Hello";   // string
let count = 42;            // number
let isReady = false;       // boolean
```

**Important**: `"42"` (with quotes) is a **string**, not a number! The quotes matter.

## Basic Math

JavaScript can do math with numbers:

```javascript
let a = 10;
let b = 3;

console.log(a + b);   // 13  (addition)
console.log(a - b);   // 7   (subtraction)
console.log(a * b);   // 30  (multiplication)
console.log(a / b);   // 3.333...  (division)
console.log(a % b);   // 1   (remainder/modulo)
```

## Combining Strings (Concatenation)

You can glue strings together with `+`:

```javascript
let first = "Hello";
let second = "World";
console.log(first + " " + second);  // "Hello World"
```

### Template Literals (The Modern Way)

There's a nicer way to build strings using **backticks** (`` ` ``) and `${}`:

```javascript
let name = "Bret";
let age = 25;
console.log(`My name is ${name} and I am ${age} years old.`);
// prints: My name is Bret and I am 25 years old.
```

This is called a **template literal**. The `${}` parts get replaced with the variable values. Much cleaner than using `+` everywhere!

## Comments

Comments are notes in your code that JavaScript ignores:

```javascript
// This is a single-line comment

/*
  This is a
  multi-line comment
*/

let x = 5; // You can also put comments at the end of a line
```

Use comments to explain **why** you're doing something, not what (the code itself shows the what).

## Try It!

Open `exercise.html` in this folder. Open the browser's console (right-click > Inspect > Console tab) to see the output of your JavaScript code.

## Key Takeaways

- JavaScript makes pages **interactive** and **dynamic**
- `console.log()` prints messages to the browser console
- Variables store data: `let name = "Bret";`
- Use `const` for values that don't change, `let` for values that do
- Three basic types: **strings** (text), **numbers**, **booleans** (true/false)
- Template literals (`` `Hello ${name}` ``) are a clean way to build strings
