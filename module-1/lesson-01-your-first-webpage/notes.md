# Lesson 1: Your First Webpage

## What is a Webpage?

Every website you've ever visited — Google, YouTube, Amazon — is made up of files that your browser reads and displays. The most basic type of file is an **HTML** file.

**HTML** stands for **HyperText Markup Language**. It's not really a "programming language" — it's a way of describing the **structure** and **content** of a page. Think of it like the blueprint for a house: it tells the browser *what* to put on the page.

## How HTML Works

HTML uses **tags** to label content. Tags look like this:

```html
<tagname>content goes here</tagname>
```

- The opening tag `<tagname>` says "this thing is starting"
- The closing tag `</tagname>` (with the `/`) says "this thing is ending"
- Everything between them is the content

For example:
```html
<p>This is a paragraph of text.</p>
```

The `<p>` tag means "paragraph." The browser reads this and knows to display it as a paragraph.

## The Basic Structure of a Webpage

Every HTML page has the same skeleton:

```html
<!DOCTYPE html>
<html>
  <head>
    <title>My Page Title</title>
  </head>
  <body>
    <!-- Your visible content goes here -->
  </body>
</html>
```

Let's break this down:

- `<!DOCTYPE html>` — Tells the browser "this is an HTML file" (you just always include this)
- `<html>` — The container for everything
- `<head>` — Information *about* the page (like its title). This stuff doesn't show up on the page itself.
- `<title>` — The text that appears in the browser tab
- `<body>` — Everything the user actually **sees** goes here
- `<!-- comment -->` — A comment. The browser ignores these. They're notes for yourself.

## Common HTML Tags

Here are the most useful tags to start with:

| Tag | What It Does | Example |
|-----|-------------|---------|
| `<h1>` to `<h6>` | Headings (h1 is biggest, h6 is smallest) | `<h1>Welcome!</h1>` |
| `<p>` | A paragraph of text | `<p>Hello there.</p>` |
| `<strong>` | **Bold** text | `<strong>important</strong>` |
| `<em>` | *Italic* text | `<em>emphasis</em>` |
| `<br>` | A line break (no closing tag needed) | `Line one<br>Line two` |
| `<ul>` | An unordered (bullet) list | See below |
| `<ol>` | An ordered (numbered) list | See below |
| `<li>` | A list item (goes inside ul or ol) | `<li>Item one</li>` |
| `<a>` | A link | `<a href="https://google.com">Click me</a>` |
| `<img>` | An image (no closing tag) | `<img src="photo.jpg">` |
| `<div>` | A generic container/box | `<div>stuff inside</div>` |

### Lists Example

```html
<ul>
  <li>Apples</li>
  <li>Bananas</li>
  <li>Oranges</li>
</ul>
```

This shows as:
- Apples
- Bananas
- Oranges

## Nesting

Tags can go inside other tags. This is called **nesting**:

```html
<p>This is <strong>very</strong> important.</p>
```

The `<strong>` tag is *nested inside* the `<p>` tag. Just make sure you close them in the right order — the last one opened is the first one closed.

## What is Indentation?

You'll notice the code examples use spaces to push things to the right:

```html
<body>
  <h1>Hello</h1>
  <p>World</p>
</body>
```

This is called **indentation**. It makes your code easier to read by showing which tags are inside which. The browser doesn't care about indentation — it's purely for humans. Use 2 spaces for each level.

## Try It!

Open the `exercise.html` file in this folder:
1. Open it in your text editor to see the code
2. Open it in your browser to see the result (double-click the file)
3. Follow the instructions in the comments to add your own content
4. Save the file, then refresh your browser to see your changes

## Key Takeaways

- HTML describes the **structure** of a webpage using **tags**
- Tags have an opening `<tag>` and closing `</tag>` (usually)
- The `<head>` has info about the page; the `<body>` has the visible stuff
- Tags can be **nested** inside each other
- **Indentation** makes code readable
