# Lesson 6: Speaking JSON

## What is JSON?

When the browser and server talk to each other, they need a shared **format** for the data. You can't just send a JavaScript object over the wire — the server doesn't speak JavaScript. And you can't send a Go struct — the browser doesn't speak Go.

**JSON** (JavaScript Object Notation) is the universal format. It's a plain text string that both sides can read and write.

## What Does JSON Look Like?

You've actually already seen it — it looks just like a JavaScript object:

```json
{
    "sender": "Bret",
    "text": "Hello!",
    "time": "2:30 PM"
}
```

Rules:
- Keys are always **strings in double quotes**
- Values can be strings, numbers, booleans, arrays, or other objects
- No trailing commas
- No comments

An array of messages in JSON:

```json
[
    { "sender": "Bret", "text": "Hey!" },
    { "sender": "Dave", "text": "Hi there!" }
]
```

## JSON in JavaScript

JavaScript makes JSON easy since the syntax is almost identical.

### Object → JSON string

```javascript
let message = { sender: "Bret", text: "Hello!", time: "2:30 PM" };
let jsonString = JSON.stringify(message);
console.log(jsonString);
// '{"sender":"Bret","text":"Hello!","time":"2:30 PM"}'
```

### JSON string → Object

```javascript
let jsonString = '{"sender":"Bret","text":"Hello!","time":"2:30 PM"}';
let message = JSON.parse(jsonString);
console.log(message.sender);  // "Bret"
```

- `JSON.stringify()` — turns an object into a JSON string (for sending)
- `JSON.parse()` — turns a JSON string back into an object (for receiving)

## JSON in Go

Go uses the `encoding/json` package and **struct tags** to map struct fields to JSON keys.

### Defining a Struct with JSON Tags

```go
type Message struct {
    Sender string `json:"sender"`
    Text   string `json:"text"`
    Time   string `json:"time"`
}
```

The backtick part (`` `json:"sender"` ``) is a **struct tag**. It tells Go: "When converting to/from JSON, this field is called `sender`" (lowercase, to match what JavaScript expects).

### Go struct → JSON string

```go
import "encoding/json"

msg := Message{
    Sender: "Bret",
    Text:   "Hello!",
    Time:   "2:30 PM",
}

jsonBytes, err := json.Marshal(msg)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println(string(jsonBytes))
// {"sender":"Bret","text":"Hello!","time":"2:30 PM"}
```

`json.Marshal()` converts a struct to JSON. It returns **bytes** (not a string), so we wrap it with `string()` to print it.

### JSON string → Go struct

```go
jsonString := `{"sender":"Bret","text":"Hello!","time":"2:30 PM"}`

var msg Message
err := json.Unmarshal([]byte(jsonString), &msg)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println(msg.Sender)  // "Bret"
fmt.Println(msg.Text)    // "Hello!"
```

`json.Unmarshal()` converts JSON back to a struct. The `&msg` means "put the result into this variable" (the `&` is a Go thing — don't worry about the details yet).

## The Full Picture

Here's what happens when Bret sends a chat message:

```
1. Bret types "Hello!" and clicks Send

2. JavaScript creates an object:
   { sender: "Bret", text: "Hello!", time: "2:30 PM" }

3. JavaScript converts to JSON string:
   '{"sender":"Bret","text":"Hello!","time":"2:30 PM"}'

4. Browser SENDS the JSON string to the server

5. Go server RECEIVES the JSON string

6. Go converts JSON to a struct:
   Message{Sender: "Bret", Text: "Hello!", Time: "2:30 PM"}

7. Go processes it (broadcasts to other users)

8. Go converts struct back to JSON string

9. Server SENDS the JSON to Dave's browser

10. Dave's JavaScript parses the JSON back to an object

11. Dave's browser displays the message
```

JSON is the **common language** in the middle that both sides understand.

## Side by Side Comparison

| Operation | JavaScript | Go |
|-----------|-----------|-----|
| Object/struct → JSON | `JSON.stringify(obj)` | `json.Marshal(struct)` |
| JSON → Object/struct | `JSON.parse(str)` | `json.Unmarshal(bytes, &struct)` |
| Define structure | `{ key: value }` (ad-hoc) | `type Name struct { ... }` (defined) |

## Try It!

Open `exercise/main.go` — you'll practice converting between Go structs and JSON. The exercise also includes a small web server that sends JSON responses.

## Key Takeaways

- **JSON** is a text format for exchanging data between client and server
- Looks like `{"key": "value"}` — similar to JavaScript objects
- JavaScript: `JSON.stringify()` and `JSON.parse()`
- Go: `json.Marshal()` and `json.Unmarshal()` from `encoding/json`
- Go uses **struct tags** (`` `json:"fieldname"` ``) to map fields to JSON keys
- JSON is the bridge between JavaScript objects and Go structs
