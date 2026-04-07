# Lesson 2: Styling with CSS

## What is CSS?

In Lesson 1, we learned that HTML defines the **structure** of a page — what's on it. But it looks pretty plain, right? That's where **CSS** comes in.

**CSS** stands for **Cascading Style Sheets**. It controls how things **look** — colors, fonts, sizes, spacing, layout, and more. If HTML is the blueprint of a house, CSS is the paint, wallpaper, and furniture.

## How to Add CSS to Your Page

There are a few ways, but we'll use the simplest: a `<style>` tag inside the `<head>` of your HTML file.

```html
<!DOCTYPE html>
<html>
  <head>
    <title>My Styled Page</title>
    <style>
      /* CSS rules go here */
    </style>
  </head>
  <body>
    <h1>Hello!</h1>
  </body>
</html>
```

## CSS Rules: The Basics

A CSS rule looks like this:

```css
selector {
  property: value;
}
```

- **Selector** — *what* you want to style (like `h1`, `p`, `body`)
- **Property** — *what aspect* you want to change (like `color`, `font-size`)
- **Value** — *what to change it to* (like `red`, `20px`)

### Example:

```css
h1 {
  color: blue;
  font-size: 36px;
}
```

This makes **all `<h1>` elements** blue and 36 pixels tall.

## Common CSS Properties

| Property | What It Does | Example Values |
|----------|-------------|----------------|
| `color` | Text color | `red`, `blue`, `#FF5733`, `rgb(255,87,51)` |
| `background-color` | Background color | `yellow`, `#333333`, `lightgray` |
| `font-size` | Text size | `16px`, `24px`, `2em` |
| `font-family` | Font type | `Arial`, `"Times New Roman"`, `monospace` |
| `font-weight` | Bold or not | `bold`, `normal`, `700` |
| `text-align` | Text alignment | `left`, `center`, `right` |
| `padding` | Space INSIDE an element | `10px`, `20px 10px` |
| `margin` | Space OUTSIDE an element | `10px`, `0 auto` |
| `border` | A border around an element | `1px solid black`, `2px dashed red` |
| `width` | How wide something is | `200px`, `50%`, `100%` |
| `height` | How tall something is | `100px`, `50%` |
| `border-radius` | Rounded corners | `5px`, `50%` (makes a circle) |

## Colors

You can specify colors in several ways:

```css
color: red;                  /* Named color */
color: #FF0000;              /* Hex code (same red) */
color: rgb(255, 0, 0);       /* RGB values (same red) */
```

There are 140+ named colors like `tomato`, `steelblue`, `coral`, `darkslategray`, etc.

## Padding vs. Margin

This is a concept that confuses even experienced developers, so don't worry if it takes a while to click:

- **Padding** = space between the content and the border (inside the box)
- **Margin** = space between the border and other elements (outside the box)

Think of it like a picture frame:
- The **picture** is the content
- The **mat** (space between picture and frame) is the padding
- The **gap between frames** on the wall is the margin

```css
p {
  padding: 10px;    /* 10px of space inside */
  margin: 20px;     /* 20px of space outside */
  border: 1px solid black;  /* makes it visible */
}
```

## Classes: Styling Specific Elements

What if you have two paragraphs but only want ONE to be red? You use a **class**.

In your HTML, add a `class` attribute:
```html
<p class="highlight">This one is special.</p>
<p>This one is normal.</p>
```

In your CSS, target the class with a dot (`.`):
```css
.highlight {
  color: red;
  background-color: yellow;
}
```

Only the paragraph with `class="highlight"` gets the styling!

You can put any name you want after `class=` — just make it descriptive. And you can reuse the same class on multiple elements.

## IDs: Styling ONE Specific Element

Similar to classes, but for a **single unique element**. Use `id` in HTML and `#` in CSS:

```html
<h1 id="main-title">Welcome</h1>
```

```css
#main-title {
  color: darkgreen;
}
```

**Rule of thumb**: Use classes for things that might repeat. Use IDs for one-of-a-kind elements.

## Combining Multiple Styles

You can put as many properties as you want inside a rule:

```css
.message-box {
  background-color: #f0f0f0;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 15px;
  margin: 10px 0;
  font-size: 14px;
  color: #333;
}
```

## Try It!

Open `exercise.html` in this folder. Follow the instructions in the comments to style the page.

## Key Takeaways

- CSS controls how things **look** (colors, sizes, spacing, layout)
- CSS rules: `selector { property: value; }`
- Use **classes** (`.classname`) to style groups of elements
- Use **IDs** (`#idname`) to style a single unique element
- **Padding** = inside space, **Margin** = outside space
- You can combine many properties in one rule
