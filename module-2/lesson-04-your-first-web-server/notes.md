# Lesson 4: Your First Web Server

## From Terminal to Browser

So far, our Go programs have printed to the terminal. Now we're going to build something you can see **in your browser** — a web server!

## What is a Web Server?

A web server is a program that:
1. Listens on a **port** (like a door number)
2. Waits for **HTTP requests** from browsers
3. Sends back **HTTP responses** (HTML, text, files, etc.)

## The Simplest Possible Server

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from my Go server!")
    })

    fmt.Println("Server starting on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

Run it with `go run main.go`, then open **http://localhost:8080** in your browser. You should see "Hello from my Go server!"

### Breaking It Down

```go
import (
    "fmt"
    "net/http"     // Go's built-in HTTP package — no install needed!
)
```

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from my Go server!")
})
```

This says: "When someone visits the path `/` (the root/homepage), run this function."

- `w` (ResponseWriter) — where you write your response (what the browser sees)
- `r` (Request) — information about the incoming request
- `fmt.Fprintf(w, ...)` — write text to the response (instead of to the terminal)

```go
http.ListenAndServe(":8080", nil)
```

This starts the server on port 8080 and **keeps it running** forever (until you press `Ctrl+C` to stop it).

## Multiple Routes

You can handle different URLs:

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the homepage!")
})

http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the about page!")
})

http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello there, visitor!")
})
```

Now:
- `http://localhost:8080/` → "Welcome to the homepage!"
- `http://localhost:8080/about` → "This is the about page!"
- `http://localhost:8080/hello` → "Hello there, visitor!"

Each URL path is called a **route**.

## Sending HTML

You're not limited to plain text — you can send HTML:

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, "<h1>Hello!</h1><p>This is <strong>HTML</strong> from Go!</p>")
})
```

`w.Header().Set("Content-Type", "text/html")` tells the browser "this is HTML, not plain text."

## Using Named Handler Functions

Instead of writing everything inline, you can use named functions:

```go
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome!")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "About us!")
}

func main() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/about", aboutPage)

    fmt.Println("Server starting on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

This is cleaner when your handlers get longer.

## Important: Stopping the Server

Your server runs until you stop it. Press `Ctrl+C` in the terminal to stop it. You'll need to restart it every time you change your Go code:

1. Press `Ctrl+C` to stop
2. Run `go run main.go` again

(This is different from the browser where you just refresh!)

## Setting Up a Go Module

Before running the exercise, you need to initialize a **Go module**. This tells Go about your project:

```
cd module-2/lesson-04-your-first-web-server/exercise
go mod init chatserver
```

This creates a `go.mod` file. You only need to do this once per project.

## Try It!

Open `exercise/main.go` and build a server with multiple routes. Run it and visit each one in your browser!

## Key Takeaways

- `net/http` is Go's built-in web server package — batteries included
- `http.HandleFunc("/path", handlerFunc)` maps a URL to a function
- The handler gets `w` (to write responses) and `r` (the request info)
- `http.ListenAndServe(":8080", nil)` starts the server
- `fmt.Fprintf(w, ...)` writes to the browser instead of the terminal
- Stop the server with `Ctrl+C`, restart with `go run main.go`
- Initialize with `go mod init projectname`
