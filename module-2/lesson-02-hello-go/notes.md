# Lesson 2: Hello Go

## Installing Go

Before we write any code, we need to install Go:

1. Go to the official site: **https://go.dev/dl/**
2. Download the installer for your operating system (Windows/Mac/Linux)
3. Run the installer — accept all the defaults
4. Open a **new terminal** (Git Bash, Command Prompt, or Terminal)
5. Type: `go version`
6. You should see something like: `go version go1.22.0 windows/amd64`

If you see a version number, you're good to go!

## Your First Go Program

Create a file called `main.go` in the exercise folder and type this:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Running It

Open your terminal, navigate to the folder with your `main.go` file, and run:

```
go run main.go
```

You should see: `Hello, World!`

### Breaking It Down

```go
package main          // Every Go file belongs to a "package"
                      // "main" is special — it means "this is a runnable program"

import "fmt"          // Import the "fmt" package (for printing/formatting)
                      // Like <script src="..."> in HTML

func main() {         // The main function — where the program starts
                      // Like how your <script> code runs when the page loads
    fmt.Println("Hello, World!")  // Print a line to the terminal
}
```

## Go vs JavaScript: A Side-by-Side

You already know these concepts from JavaScript! Here's how they translate:

### Variables

```javascript
// JavaScript
let name = "Bret";
let age = 25;
let isLearning = true;
```

```go
// Go
name := "Bret"
age := 25
isLearning := true
```

The `:=` operator in Go is like `let` in JavaScript — it creates a variable and figures out the type automatically. Notice: **no semicolons** in Go!

You can also declare the type explicitly:

```go
var name string = "Bret"
var age int = 25
var isLearning bool = true
```

But `:=` is shorter and preferred when Go can figure out the type.

### Printing

```javascript
// JavaScript
console.log("Hello");
console.log("Name:", name);
console.log(`I am ${age} years old`);
```

```go
// Go
fmt.Println("Hello")
fmt.Println("Name:", name)
fmt.Printf("I am %d years old\n", age)
```

- `fmt.Println()` — print a line (like `console.log`)
- `fmt.Printf()` — formatted print (use `%s` for strings, `%d` for numbers, `%v` for anything)

### Data Types

| Type | JavaScript | Go |
|------|-----------|-----|
| Text | `"hello"` (string) | `"hello"` (string) |
| Whole number | `42` (number) | `42` (int) |
| Decimal | `3.14` (number) | `3.14` (float64) |
| True/false | `true` (boolean) | `true` (bool) |

Go is **stricter** about types than JavaScript. You can't accidentally add a string and a number — Go will tell you it's an error before the program even runs. This is actually a good thing — it catches mistakes early!

### Constants

```javascript
// JavaScript
const appName = "ChatApp";
```

```go
// Go
const appName = "ChatApp"
```

Almost identical!

## Comments

```go
// This is a single-line comment (same as JavaScript!)

/*
   This is a multi-line
   comment (also the same!)
*/
```

## A Note on Errors

Go is strict. If you create a variable and don't use it, Go won't compile your program. If you import a package and don't use it, Go won't compile. This feels annoying at first, but it keeps your code clean.

```go
package main

import "fmt"

func main() {
    name := "Bret"  // ERROR: "name declared but not used"
}
```

Fix: either use the variable or remove it.

## Try It!

Open the `exercise` folder. There's a `main.go` file with guided exercises. Run it with `go run main.go` in your terminal.

## Key Takeaways

- Install Go from **go.dev/dl**, verify with `go version`
- Run programs with `go run filename.go`
- Every program needs `package main` and `func main()`
- `:=` creates variables (like `let` in JavaScript)
- `fmt.Println()` prints to the terminal (like `console.log`)
- Go is stricter than JavaScript — it won't let you have unused variables or imports
- The syntax is very similar to JavaScript — you already know the concepts!
