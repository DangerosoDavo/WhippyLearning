# Lesson 5: Serving Your Chat App

## From Hardcoded Text to Real Files

In the last lesson, we wrote HTML directly in our Go code as strings. That works for tiny examples, but for a real app, you want to serve **actual HTML files** — like the chat app we built in Module 1.

## Go's File Server

Go has a built-in way to serve files from a folder:

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Serve all files from the "static" folder
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    fmt.Println("Server starting on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

That's it! If you have a `static` folder with an `index.html` file, visiting `http://localhost:8080` will show that page.

### How It Works

- `http.Dir("./static")` — "look for files in the `static` folder"
- `http.FileServer(...)` — "serve those files to anyone who asks"
- `http.Handle("/", fs)` — "for any URL, check the static folder"

The server automatically maps URLs to files:
- `/` → `static/index.html`
- `/style.css` → `static/style.css`
- `/app.js` → `static/app.js`

### Handle vs HandleFunc

Notice we used `http.Handle` (not `http.HandleFunc`). The difference:
- `HandleFunc` — takes a **function** you write
- `Handle` — takes a **handler object** (like `FileServer` returns)

Don't worry about the distinction too much. Just know `FileServer` uses `Handle`.

## Setting Up the Folder Structure

For our chat app, we'll organize things like this:

```
lesson-05/exercise/
├── main.go              (the Go server)
├── go.mod               (Go module file)
└── static/              (files served to the browser)
    └── index.html       (our chat app)
```

All the frontend code (HTML, CSS, JS) goes in `static/`. The Go server just hands these files to the browser.

## Why Not Just Double-Click the HTML File?

Good question! For our Module 1 chat app, double-clicking worked fine. But once we add WebSockets (coming soon), the browser needs to connect **back** to the server. That only works over HTTP, not from a `file://` URL.

Serving files through Go gives us:
- A proper `http://` URL
- The ability to add server-side features (like WebSockets)
- A setup that works just like a real website

## Separating the Frontend File

We're going to take the chat app from Module 1 (Lesson 10) and put it all in a single `index.html` file inside `static/`. The CSS and JavaScript are embedded in the same file using `<style>` and `<script>` tags — exactly like we've been doing.

Later, in a bigger project, you might split things into separate `.css` and `.js` files. But for now, one file keeps things simple.

## Mixing File Serving with Custom Routes

You can serve files AND have custom routes:

```go
func main() {
    // Custom API route
    http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Server is running!")
    })

    // Serve static files for everything else
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    fmt.Println("Server starting on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

**Important**: Register specific routes (like `/api/status`) **before** the catch-all file server. Go checks routes in order — the first match wins.

This pattern is exactly what we'll use: static files for the frontend + a WebSocket route for real-time messaging.

## Try It!

Open the `exercise` folder. It has a `main.go` and a `static/index.html`. Complete the exercises to serve the chat app from Go.

## Key Takeaways

- `http.FileServer(http.Dir("./static"))` serves files from a folder
- Put frontend files (HTML/CSS/JS) in the `static` folder
- The server maps URLs to files automatically (`/` → `index.html`)
- Serving via Go gives us `http://` URLs, needed for WebSockets later
- You can mix file serving with custom routes
- Register specific routes before the catch-all file server
