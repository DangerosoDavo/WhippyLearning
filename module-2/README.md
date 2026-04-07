# Module 2: Building a Real-Time Backend with Go

## Prerequisites

- Complete **Module 1** first — you'll need HTML, CSS, and JavaScript fundamentals
- **Go** installed on your machine (Lesson 2 walks you through this)
- A text editor (VS Code recommended)

## What You'll Build

You'll take the chat app from Module 1 and add a **Go backend server** with **WebSocket** support, turning it into a real-time multi-user chat application. Two (or more!) people can chat with each other live.

## What You'll Need to Install

- **Go** — download from https://go.dev/dl/ (covered in Lesson 2)
- The `gorilla/websocket` Go package (covered in Lesson 7)

## Lesson Plan

| # | Lesson | What You'll Learn |
|---|--------|-------------------|
| 1 | How the Internet Works | Client/server model, HTTP, ports, why we need a backend |
| 2 | Hello Go | Installing Go, your first program, variables, printing |
| 3 | Go Essentials | Functions, if/else, loops, structs, slices |
| 4 | Your First Web Server | Building an HTTP server with `net/http` |
| 5 | Serving Your Chat App | Serving static HTML/CSS/JS files from Go |
| 6 | Speaking JSON | JSON format, encoding/decoding in Go and JavaScript |
| 7 | What Are WebSockets? | HTTP vs WebSocket, persistent connections, gorilla/websocket |
| 8 | Building the WebSocket Server | Accepting connections, reading/writing, goroutines, broadcasting |
| 9 | Connecting to the Server | JS WebSocket API, connection lifecycle, status indicator, auto-reconnect |
| 10 | Sending & Receiving Messages | Wiring the chat UI to WebSocket, JSON.stringify/parse, the message flow |
| 11 | Broadcasting to Everyone | Concurrency with mutex, message types, join/leave notifications, user count |
| 12 | Handling the Real World | Input validation, message history, typing indicators, the finished app |

## How Each Lesson Works

- **notes.md** — Read this first. Explains concepts with code examples.
- **exercise/** folder — Contains a `main.go` (Go server) and often a `static/index.html` (frontend). Follow the guided comments to build each piece.

## Running Go Exercises

For lessons with Go code:

```
cd module-2/lesson-XX-name/exercise
go mod init chatserver          # first time only
go get github.com/gorilla/websocket  # lessons 7+ only
go run main.go                  # start the server
```

Then open `http://localhost:8080` in your browser. Press `Ctrl+C` to stop the server.

## The Journey

```
Lessons 1-3:    Learn Go basics (you already know the concepts from JS!)
Lessons 4-5:    Build a web server that serves your chat app
Lesson 6:       Learn JSON — the language between frontend and backend
Lessons 7-8:    Add WebSocket support to the server
Lessons 9-10:   Connect the frontend to the server
Lessons 11-12:  Handle multiple users, polish, and finish!
```

By Lesson 12, you'll have a fully working real-time chat app with a Go backend!
