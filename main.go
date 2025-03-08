package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand/v2"

	"github.com/gin-gonic/gin"
)

// Struct to represent our Server
type Server struct {
	Router *gin.Engine
	Port   string
}

// New Server function
func NewServer(port string) *Server {
	return &Server{
		Router: gin.Default(),
		Port:   port,
	}
}

func main() {

	// create a new server with a provided port
	s := NewServer("8080")

	// run server unless we have an error
	err := s.Run()
	if err != nil {
		return
	}
}

// Run server function with router and routes
func (s *Server) Run() error {

	// Favicon endpoint
	s.Router.GET("/favicon.ico", s.faviconHandler)

	// Root server ack
	s.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"demo": "golang-favicon-generator"})
	})

	// Start server
	if err := s.Router.Run(":" + s.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	return nil
}

// faviconHandler generates and returns a random favicon
func (s *Server) faviconHandler(c *gin.Context) {
	// Generate a 16x16 random image
	img := generateFavicon(16, 16)

	// Set response content type
	c.Writer.Header().Set("Content-Type", "image/png")

	// Encode image as PNG and write to response
	png.Encode(c.Writer, img)
}

// generateFavicon creates a simple random-colored favicon
func generateFavicon(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill with random colors using math/rand/v2
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r := uint8(rand.IntN(256)) // Random Red value
			g := uint8(rand.IntN(256)) // Random Green value
			b := uint8(rand.IntN(256)) // Random Blue value
			img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}
	return img
}
