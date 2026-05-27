# API Check 🚀

A fast and lightweight **CLI API endpoint checker** written in Go.

`apicheck` reads a list of URLs from a file and checks:

- ✅ API availability
- ⚡ Response time
- 📡 HTTP status codes
- ❌ Connection failures
- ⏱ Request timeout handling

Built with **Go concurrency (goroutines + WaitGroup)** for fast parallel checks.

---

## Features

- Concurrent API checking
- Response time measurement
- HTTP status code reporting
- Timeout support
- Invalid URL detection
- Error handling
- Lightweight and fast
- Simple CLI interface

---

## Installation

### Clone Repository

```bash
git clone https://github.com/yourusername/apicheck.git
cd apicheck
```

### Initialize Go Modules

```bash
go mod tidy
```

### Build

Linux/macOS:

```bash
go build -o apicheck
```

Windows:

```bash
go build -o apicheck.exe
```

---

## Usage

Command:

```bash
apicheck urls.txt
```

Linux/macOS:

```bash
./apicheck urls.txt
```

Windows:

```bash
apicheck.exe urls.txt
```

---

## Example Input File

Create a file named `urls.txt`:

```txt
https://jsonplaceholder.typicode.com/posts
https://httpbin.org/status/200
https://httpbin.org/status/500
https://api.github.com
https://google.com
invalid-url
https://thisdoesnotexist123456.com
```

You can also add comments:

```txt
# Production APIs
https://api.github.com
https://google.com

# Testing
https://httpbin.org/status/500
```

Empty lines and comments are ignored.

---

## Example Output

Success:

```txt
✓ https://google.com 200 (143ms)
✓ https://api.github.com 200 (221ms)
```

Warning/Error Status:

```txt
⚠ https://httpbin.org/status/500 500 (98ms)
```

Invalid URL:

```txt
✗ invalid-url ERROR: invalid URL
```

Connection Failure:

```txt
✗ https://thisdoesnotexist123456.com ERROR: dial tcp: lookup failed
```

---

## Project Structure

```txt
apicheck/
│── main.go
│── urls.txt
│── go.mod
│── README.md
```

---

## How It Works

1. Reads URLs from a text file
2. Launches concurrent goroutines
3. Sends HTTP GET requests
4. Measures response time
5. Prints status code and latency
6. Handles failures gracefully

---

## Tech Stack

- **Go**
- `net/http`
- `goroutines`
- `channels`
- `sync.WaitGroup`

---

## Error Handling

The program handles:

- Missing input file
- Invalid URLs
- DNS failures
- Connection errors
- HTTP failures
- Request timeouts
- Empty files

---

## Why This Project?

This project demonstrates practical Go fundamentals:

- CLI development
- File handling
- HTTP networking
- Concurrency
- Synchronization
- Error handling
- Clean terminal output

Ideal as a **beginner-to-intermediate Go portfolio project**.

---

## Future Improvements

Planned ideas:

- [ ] JSON output mode
- [ ] CSV report export
- [ ] Colored terminal output
- [ ] Retry mechanism
- [ ] Configurable timeout
- [ ] Support for custom headers
- [ ] API authentication tokens
- [ ] Worker pool concurrency
- [ ] HTTP method selection (`GET`, `POST`, `HEAD`)

---

## Contributing

Pull requests are welcome.

If you find bugs or want to improve the project, feel free to open an issue or submit a PR.

---

## License

MIT License

---

## Example Showcase

```bash
$ apicheck urls.txt

✓ https://google.com 200 (143ms)
✓ https://api.github.com 200 (221ms)
⚠ https://httpbin.org/status/500 500 (98ms)
✗ invalid-url ERROR: invalid URL
```

Fast. Concurrent. Simple.
