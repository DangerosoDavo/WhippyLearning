# Lesson 1: How the Internet Works

## Where We Left Off

In Module 1, you built a working chat app — but it only works for **one person**. Messages are saved in your own browser and nobody else can see them. To make a *real* chat app where multiple people can talk to each other, we need a **server**.

## The Client-Server Model

When you use any website or app, there are two sides:

- **Client** — the thing you interact with (your browser, your phone app)
- **Server** — a computer somewhere that the client talks to

Here's what happens when you visit a website:

```
1. You type "google.com" in your browser        (client)
2. Your browser sends a REQUEST to Google        (client → server)
3. Google's server processes the request          (server)
4. The server sends back a RESPONSE (the page)   (server → client)
5. Your browser displays it                      (client)
```

This **request → response** cycle is the foundation of the entire internet.

## What is HTTP?

**HTTP** (HyperText Transfer Protocol) is the language clients and servers speak. When your browser asks for a webpage, it sends an **HTTP request**. The server sends back an **HTTP response**.

You've actually seen HTTP in action — every URL starts with `http://` or `https://` (the "s" means secure/encrypted).

### Types of HTTP Requests

| Method | Purpose | Example |
|--------|---------|---------|
| **GET** | "Give me something" | Loading a webpage, fetching data |
| **POST** | "Here's some data" | Submitting a form, sending a message |

There are others (PUT, DELETE, etc.) but GET and POST are the two you'll see the most.

## What Does a Server Actually Do?

A server is just a program that:

1. **Listens** for incoming connections on a **port** (like a specific door number)
2. **Receives** requests from clients
3. **Processes** the request (looks up data, runs logic, etc.)
4. **Sends** a response back

When we say "server," we usually mean the **software**, not the physical computer (though the computer running it is also called a server).

## Ports: The Door Numbers

A single computer can run many services at once (a web server, an email server, etc.). Each one listens on a different **port number**.

- Port `80` — standard HTTP
- Port `443` — standard HTTPS
- Port `8080` — commonly used for development

When you visit `http://localhost:8080`, you're saying: "Connect to **my own computer** (`localhost`) on **port 8080**."

We'll use `localhost` throughout this module because we're running the server on the same machine as the browser.

## Why Do We Need a Server for Chat?

Think about it: if Bret sends a message, how does Dave see it?

**Without a server:**
```
Bret's Browser → saves to Bret's localStorage
Dave's Browser → has no idea a message was sent
```

**With a server:**
```
Bret's Browser → sends message to SERVER
SERVER → forwards message to Dave's Browser
Dave's Browser → displays the message
```

The server is the **middleman** that connects everyone together.

## What We're Building

Over the next 11 lessons, we're going to:

1. Learn **Go** — a programming language perfect for building servers
2. Build a **web server** that serves our chat app files
3. Add **WebSocket** support for real-time messaging
4. Connect our **frontend** (the Module 1 chat app) to the server
5. Enable **multiple users** to chat in real-time

By the end, you'll be able to open two browser tabs (or two computers on the same network) and have a real conversation!

## What is Go?

**Go** (also called **Golang**) is a programming language created by Google. It's great for building servers because:

- It's **simple and readable** — you'll see how similar it feels to JavaScript
- It has a huge **standard library** — you can build a web server without installing anything extra
- It's **fast** — much faster than JavaScript for server tasks
- It handles **multiple connections** easily — perfect for chat apps

## What's Different About Backend Code?

In Module 1, all your code ran **in the browser**. Backend code is different:

| Frontend (Browser) | Backend (Server) |
|--------------------|-----------------|
| Runs in the browser | Runs in a terminal/command line |
| Written in HTML/CSS/JS | Written in Go (or Python, Java, etc.) |
| The user sees it | The user never sees it |
| One user at a time | Handles many users at once |
| Files open with double-click | Files run with a command |

## Try It!

This lesson is conceptual — no exercise file. But here's something you can try:

1. Open your browser
2. Press `F12` to open Developer Tools
3. Click the **Network** tab
4. Visit any website
5. Watch all the HTTP requests fly by!

Each line is a request your browser made to a server, and each one got a response back.

## Key Takeaways

- **Client** = your browser. **Server** = a program that clients talk to.
- **HTTP** is the language of the web (request → response)
- A **port** is like a door number on a computer
- **localhost** means "this computer"
- A chat app needs a server so users can send messages to each other
- **Go** is the language we'll use to build our server
