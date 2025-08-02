# 🧠 Go Exercises

This folder contains small, focused Go programs that I’ve built as part of my journey learning the Go programming language. Each subfolder represents a single exercise, complete built as its own package.

These exercises help me practice Go fundamentals, explore packages from the standard library and even the standard library extension, and solve real-world problems in small chunks.

---

## 📁 Directory Structure

```bash
exercises/
├── html-parser/
│ ├── html-parser.go
│ └── README.md
├── another-exercise/
│ ├── ...
```


Each folder includes:

- `xxx.go` — the Go source file for the exercise  
- `README.md` — a short description of the exercise, how to run it, and any key notes

---

## ✅ Completed Exercises

### 1. [`html-parser`](./html-parser)

**Description:**  
A Go program that parses a raw HTML page and returns:

- The number of **words**
- The number of **images**

**Key Concepts:**

- Parsing HTML with `golang.org/x/net/html`
- Navigating the HTML node tree
- Counting text content and `<img>` tags

**Usage:**

```bash
cd exercises/html-parser
go run main.go https://example.com

```
If no URL is provided, the program can optionally read from a local HTML file or stdin (based on how you extend it).

### 🚀 How to Run Any Exercise
Each folder is self-contained. To run any exercise:

```bash
cd exercises/<exercise-folder>
go run main.go [arguments...]
```
### Example:

```bash
cd exercises/html-parser
go run main.go 
```


### 🛠️ Upcoming Exercises (Planned)
 
  [ ] JSON encoder/decoder playground

  [ ] Simple CLI file linter

  [ ] Concurrent web scraper

  [ ] File watcher using fsnotify

  [ ] Goroutine-based worker pool demo

### 💡 Why This Structure?
Organizing exercises this way allows me to:

Focus on one concept at a time

Revisit and improve past exercises easily

Share and reuse code cleanly

Track my Go learning milestones

### 📚 Learning Resources
[The Go Programming Language]([http](https://github.com/heavykenny/book/blob/master/Go/The.Go.Programming.Language.pdf))

[Matt Holiday Go Class](https://www.youtube.com/watch?v=iDQAZEJK8lI&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)

Happy learning! 🚀

