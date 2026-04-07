# Lesson 7: Lists of Things (Arrays and Loops)

## Why Arrays?

So far, our variables have held one thing at a time — one name, one number. But what about a **list** of things? A chat app needs to store multiple messages. A todo app needs multiple tasks. That's where **arrays** come in.

## What is an Array?

An array is an ordered list of values:

```javascript
let fruits = ["apple", "banana", "cherry"];
let scores = [95, 82, 67, 90];
let mixed = ["hello", 42, true];
```

- Square brackets `[ ]` define an array
- Items are separated by commas
- Arrays can hold any type of data

## Accessing Array Items

Each item has a **position** (called an index), starting from **0** (not 1!):

```javascript
let fruits = ["apple", "banana", "cherry"];
//             index 0   index 1   index 2

console.log(fruits[0]);  // "apple"
console.log(fruits[1]);  // "banana"
console.log(fruits[2]);  // "cherry"
```

Yes, the first item is index 0. This is confusing at first but you'll get used to it!

## Array Length

```javascript
let fruits = ["apple", "banana", "cherry"];
console.log(fruits.length);  // 3
```

## Adding and Removing Items

```javascript
let fruits = ["apple", "banana"];

// Add to the END
fruits.push("cherry");
// fruits is now ["apple", "banana", "cherry"]

// Remove from the END
let removed = fruits.pop();
// removed is "cherry", fruits is ["apple", "banana"]

// Add to the BEGINNING
fruits.unshift("mango");
// fruits is now ["mango", "apple", "banana"]
```

For our chat app, we'll mainly use `.push()` to add new messages to the list.

## Loops: Doing Things Repeatedly

A **loop** runs the same code multiple times. The most common type is a `for` loop:

```javascript
for (let i = 0; i < 5; i++) {
  console.log(i);
}
// prints: 0, 1, 2, 3, 4
```

Let's break it down:
- `let i = 0` — start at 0
- `i < 5` — keep going while i is less than 5
- `i++` — add 1 to i after each round (`i++` is shorthand for `i = i + 1`)

## Looping Through an Array

The most common use of loops — doing something with each item in an array:

```javascript
let fruits = ["apple", "banana", "cherry"];

for (let i = 0; i < fruits.length; i++) {
  console.log(fruits[i]);
}
// prints: apple, banana, cherry
```

### The Easier Way: for...of

There's a simpler syntax when you just need each item:

```javascript
let fruits = ["apple", "banana", "cherry"];

for (let fruit of fruits) {
  console.log(fruit);
}
// prints: apple, banana, cherry
```

Much cleaner! Use `for...of` when you don't need the index number.

## forEach: Another Common Pattern

Arrays have a built-in `forEach` method:

```javascript
let fruits = ["apple", "banana", "cherry"];

fruits.forEach(function(fruit) {
  console.log(fruit);
});
```

All three approaches do the same thing. Use whichever feels most natural.

## Practical Example: Rendering a List

Here's something close to what our chat app will do — take an array of data and display it on the page:

```javascript
let messages = [
  "Hey, how are you?",
  "I'm good, thanks!",
  "Want to grab lunch?"
];

let chatBox = document.getElementById("chat");

for (let message of messages) {
  let p = document.createElement("p");
  p.textContent = message;
  chatBox.appendChild(p);
}
```

This loops through all messages and creates a `<p>` element for each one. This is essentially how our chat app will display messages!

## Finding Items

```javascript
let fruits = ["apple", "banana", "cherry"];

// Check if something exists
console.log(fruits.includes("banana"));  // true
console.log(fruits.includes("grape"));   // false

// Find the position
console.log(fruits.indexOf("cherry"));  // 2
console.log(fruits.indexOf("grape"));   // -1 (not found)
```

## Objects: A Quick Preview

Arrays are great for lists, but sometimes you need to group related information together. That's what **objects** are for:

```javascript
let message = {
  text: "Hello!",
  sender: "Bret",
  time: "2:30 PM"
};

console.log(message.text);    // "Hello!"
console.log(message.sender);  // "Bret"
```

Objects use `{ }` and store data as **key: value** pairs. You access values with a dot: `object.key`.

You can have an **array of objects** — which is exactly how we'll store chat messages:

```javascript
let messages = [
  { text: "Hey!", sender: "Bret", time: "2:30 PM" },
  { text: "Hi there!", sender: "Dave", time: "2:31 PM" }
];

console.log(messages[0].text);    // "Hey!"
console.log(messages[1].sender);  // "Dave"
```

## Try It!

Open `exercise.html` for hands-on practice.

## Key Takeaways

- **Arrays** store ordered lists: `let items = ["a", "b", "c"]`
- Access items by index (starting at 0): `items[0]`
- `.push()` adds to the end, `.length` gives the count
- **Loops** repeat code: `for (let item of array) { ... }`
- **Objects** group related data: `{ key: value, key: value }`
- Arrays of objects are perfect for storing chat messages
