# Lesson 6: Making Decisions (If/Else and Functions)

## If Statements: Making Choices

Programs need to make decisions. "If the user typed something, send it. If not, do nothing." This is done with **if statements**:

```javascript
let age = 18;

if (age >= 18) {
  console.log("You are an adult.");
}
```

The code inside the `{ }` only runs **if the condition is true**.

## If/Else: Two Paths

```javascript
let temperature = 35;

if (temperature > 30) {
  console.log("It's hot outside!");
} else {
  console.log("It's not that hot.");
}
```

- If the condition is true → the first block runs
- If the condition is false → the `else` block runs

## If/Else If/Else: Multiple Paths

```javascript
let score = 75;

if (score >= 90) {
  console.log("A - Excellent!");
} else if (score >= 80) {
  console.log("B - Good job!");
} else if (score >= 70) {
  console.log("C - Not bad.");
} else {
  console.log("Keep studying!");
}
```

JavaScript checks each condition from top to bottom and runs the **first one** that's true.

## Comparison Operators

These are used in conditions:

| Operator | Meaning | Example |
|----------|---------|---------|
| `===` | Equals (strict) | `name === "Bret"` |
| `!==` | Not equals | `status !== "offline"` |
| `>` | Greater than | `age > 18` |
| `<` | Less than | `score < 50` |
| `>=` | Greater than or equal | `count >= 10` |
| `<=` | Less than or equal | `temp <= 0` |

**Important**: Use `===` (three equals), not `==` (two equals). The triple equals is stricter and avoids weird bugs. Same goes for `!==` instead of `!=`.

## Combining Conditions

Use `&&` (AND) and `||` (OR):

```javascript
let age = 25;
let hasTicket = true;

// Both must be true
if (age >= 18 && hasTicket) {
  console.log("You can enter.");
}

// At least one must be true
if (age < 13 || age >= 65) {
  console.log("You get a discount.");
}
```

- `&&` — AND: both sides must be true
- `||` — OR: at least one side must be true
- `!` — NOT: flips true to false and vice versa

```javascript
let isEmpty = true;
if (!isEmpty) {
  console.log("Not empty!");
}
```

## Truthy and Falsy

In JavaScript, some values are considered "falsy" (treated as false in an if statement):
- `false`
- `0`
- `""` (empty string)
- `null`
- `undefined`

Everything else is "truthy." This is useful for checking if a string has content:

```javascript
let message = "";

if (message) {
  console.log("Message has content");
} else {
  console.log("Message is empty");
}
// prints: "Message is empty"
```

This will be handy in our chat app — we won't send empty messages!

## Functions: Reusable Blocks of Code

A **function** is a named block of code that you can run whenever you want:

```javascript
function sayHello() {
  console.log("Hello there!");
}

// Call the function (run it)
sayHello();  // prints: Hello there!
sayHello();  // prints it again!
```

### Functions with Parameters

Functions can accept **inputs** (called parameters):

```javascript
function greet(name) {
  console.log(`Hello, ${name}!`);
}

greet("Bret");   // Hello, Bret!
greet("Dave");   // Hello, Dave!
```

Multiple parameters:

```javascript
function add(a, b) {
  console.log(a + b);
}

add(5, 3);  // 8
add(10, 20); // 30
```

### Functions That Return Values

Functions can send a value back using `return`:

```javascript
function add(a, b) {
  return a + b;
}

let result = add(5, 3);
console.log(result);  // 8

// Or use it directly
console.log(add(10, 20));  // 30
```

### Why Use Functions?

1. **Reuse** — Write once, use many times
2. **Organization** — Break code into logical pieces
3. **Readability** — Give meaningful names to blocks of code

### Real-World Example

```javascript
function createMessageElement(text) {
  let p = document.createElement("p");
  p.textContent = text;
  p.style.padding = "8px";
  p.style.backgroundColor = "#f0f0f0";
  p.style.borderRadius = "8px";
  p.style.marginBottom = "5px";
  return p;
}

// Now creating styled messages is one line:
let msg1 = createMessageElement("Hey there!");
let msg2 = createMessageElement("How's it going?");
document.body.appendChild(msg1);
document.body.appendChild(msg2);
```

This is exactly the kind of function we'll use in our chat app!

## Try It!

Open `exercise.html` for practice with if/else and functions.

## Key Takeaways

- `if/else` lets your code make decisions
- Use `===` and `!==` for comparisons (not `==` or `!=`)
- `&&` = AND, `||` = OR, `!` = NOT
- Empty strings are "falsy" — `if (message)` checks if it has content
- **Functions** are reusable blocks of code: `function name() { ... }`
- Functions can take **parameters** (inputs) and **return** values (outputs)
- Functions help organize code and avoid repetition
