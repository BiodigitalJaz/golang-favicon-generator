# Golang Favicon Generator

A simple Go web service that generates random favicons on-the-fly.

## Description

This service creates a Gin-based web server that dynamically generates random colorful 16x16 pixel favicons whenever a browser or client requests a `/favicon.ico`. Each request produces a unique, randomly-colored favicon.

## Features

- HTTP server using the Gin framework
- Dynamic favicon generation with random colors
- Simple JSON response on the root endpoint
- Configurable port via server initialization

## Requirements

- Go 1.22+ (uses the new `math/rand/v2` package)
- [Gin](https://github.com/gin-gonic/gin) web framework

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/biodigitalJaz/golang-favicon-generator.git
   cd golang-favicon-generator
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

## Usage

### Run the server

```
go run main.go
```

The server will start on port 8080 by default.

### Endpoints

- `GET /`: Returns a simple JSON response confirming the service is running
- `GET /favicon.ico`: Generates and serves a random 16x16 pixel favicon

### Changing the port

To use a different port, modify the `NewServer` call in `main.go`:

```go
s := NewServer("3000") // Change to your desired port
```

## How it works

1. The server is initialized with a specified port
2. When a request to `/favicon.ico` is received, the `faviconHandler` function is called
3. The handler generates a 16x16 pixel image with random colors for each pixel
4. The image is encoded as PNG and sent in the response with the appropriate content type
