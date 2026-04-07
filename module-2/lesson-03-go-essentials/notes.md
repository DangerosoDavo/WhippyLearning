# Lesson 3: Go Essentials

## Functions

You already know what functions are from JavaScript. Go functions look very similar:

```javascript
// JavaScript
function add(a, b) {
    return a + b;
}
```

```go
// Go
func add(a int, b int) int {
    return a + b
}
```

Key differences:
- `func` instead of `function`
- Parameter types come **after** the name: `a int` not `int a`
- The **return type** comes after the parameters: `func add(...) int`
- If parameters share a type, you can shorten: `func add(a, b int) int`

### Functions with no return value

```go
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

No return type needed if you don't return anything.

### Multiple return values

Go can return **multiple values** — something JavaScript can't do easily:

```go
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// Using it:
q, r := divide(10, 3)
fmt.Println(q, r)  // 3, 1
```

This is used a lot in Go, especially for returning a result AND an error.

## If/Else

Almost identical to JavaScript, but **no parentheses** around the condition:

```javascript
// JavaScript
if (age >= 18) {
    console.log("Adult");
} else {
    console.log("Minor");
}
```

```go
// Go
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

The only difference: no `( )` around the condition. The `{ }` are always required.

## For Loops

Go has **only one loop keyword**: `for`. But it can do everything:

### Standard for loop (like JavaScript)

```javascript
// JavaScript
for (let i = 0; i < 5; i++) {
    console.log(i);
}
```

```go
// Go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Nearly identical! Just no parentheses and `let` becomes `:=`.

### While-style loop

Go doesn't have `while` — you just use `for` with only a condition:

```go
count := 0
for count < 10 {
    fmt.Println(count)
    count++
}
```

### Looping over a collection (range)

```javascript
// JavaScript
let fruits = ["apple", "banana", "cherry"];
for (let fruit of fruits) {
    console.log(fruit);
}
```

```go
// Go
fruits := []string{"apple", "banana", "cherry"}
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```

`range` gives you two values: the **index** and the **value**. The `_` means "I don't need the index." If you want it:

```go
for i, fruit := range fruits {
    fmt.Printf("%d: %s\n", i, fruit)
}
```

## Slices: Go's Version of Arrays

In Go, the most common list type is a **slice** (think of it as a flexible array):

```go
// Create a slice
names := []string{"Bret", "Dave", "Alex"}

// Access by index
fmt.Println(names[0])  // "Bret"

// Length
fmt.Println(len(names))  // 3

// Add an item (like JavaScript's .push())
names = append(names, "Sam")
```

Key differences from JavaScript arrays:
- You specify the **type** of items: `[]string`, `[]int`, `[]bool`
- All items must be the **same type** (no mixing strings and numbers)
- Use `append()` instead of `.push()`
- Use `len()` instead of `.length`

## Structs: Go's Version of Objects

In JavaScript, you create objects on the fly:

```javascript
// JavaScript
let message = {
    sender: "Bret",
    text: "Hello!",
    time: "2:30 PM"
};
```

In Go, you first **define** the structure, then create instances of it:

```go
// Go - Step 1: Define the struct (the blueprint)
type Message struct {
    Sender string
    Text   string
    Time   string
}

// Step 2: Create an instance
message := Message{
    Sender: "Bret",
    Text:   "Hello!",
    Time:   "2:30 PM",
}

// Access fields with a dot (same as JavaScript)
fmt.Println(message.Sender)  // "Bret"
fmt.Println(message.Text)    // "Hello!"
```

**Important**: In Go, field names that start with a **capital letter** are "exported" (visible outside the package). We'll always capitalize our struct fields since we'll need them for JSON later.

Structs will be crucial for our chat app — we'll use a `Message` struct to define what a chat message looks like.

## Error Handling

Go doesn't have `try/catch` like JavaScript. Instead, functions return an **error** as a second value:

```go
import (
    "fmt"
    "strconv"
)

// strconv.Atoi converts a string to an int
number, err := strconv.Atoi("42")
if err != nil {
    fmt.Println("That's not a number!")
} else {
    fmt.Println("The number is:", number)
}
```

- `nil` in Go is like `null` in JavaScript — it means "nothing"
- If `err != nil`, something went wrong
- This pattern (value, error) is everywhere in Go

You don't need to master this pattern right now. Just know that when you see `something, err :=`, the second value is an error check.

## Putting It All Together

Here's a mini program using everything from this lesson:

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func greet(p Person) string {
    if p.Age >= 18 {
        return fmt.Sprintf("Hello %s, welcome!", p.Name)
    }
    return fmt.Sprintf("Hi %s, nice to meet you!", p.Name)
}

func main() {
    people := []Person{
        {Name: "Bret", Age: 25},
        {Name: "Dave", Age: 30},
        {Name: "Alex", Age: 15},
    }

    for _, person := range people {
        message := greet(person)
        fmt.Println(message)
    }
}
```

This creates a list of people, loops through them, and greets each one differently based on age. Notice how readable it is — even if Go is new, you can follow the logic!

## Try It!

Open `exercise/main.go` and complete the exercises.

## Key Takeaways

- Go functions declare types after names: `func add(a, b int) int`
- `if/else` and `for` loops are like JavaScript but **without parentheses**
- Go only has `for` — no `while` keyword needed
- **Slices** are Go's arrays: `[]string{"a", "b"}`, use `append()` to add
- **Structs** are Go's objects: define the shape, then create instances
- Errors are returned as values, not thrown: `value, err := something()`
- Capital letter fields/functions = accessible from other packages
