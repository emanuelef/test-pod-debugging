package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	externalURL = "https://pokeapi.co/api/v2/pokemon/ditto"
)

var notToLogEndpoints = []string{"/health", "/metrics"}

var (
	secondaryHost     = getEnv("API_HOST", "localhost")
	secondaryAddress  = fmt.Sprintf("http://%s:8080", secondaryHost)
	secondaryHelloUrl = fmt.Sprintf("%s/hello", secondaryAddress)
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	// Create a new Gin router
	r := gin.New()
	r.Use(gin.Recovery())

	// Define a health endpoint
	r.GET("/health", func(c *gin.Context) {
		slog.Info("Called /health")
		c.JSON(http.StatusNoContent, gin.H{})
	})

	// Define a hello endpoint
	r.GET("/hello", func(c *gin.Context) {
		slog.Info("Called /hello")
		c.JSON(http.StatusNoContent, gin.H{})
	})

	// Run the router in a goroutine
	go func() {
		host := getEnv("HOST", "0.0.0.0")
		port := getEnv("PORT", "8080")
		hostAddress := fmt.Sprintf("%s:%s", host, port)

		slog.Info("Starting server")
		err := r.Run(hostAddress)
		if err != nil {
			slog.Info("Starting router failed, %v", err)
		}
	}()

	// Start a ticker to perform periodic calls
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Perform the GET request to secondaryHelloUrl here
			resp, err := http.Get(secondaryHelloUrl)
			if err != nil {
				slog.Info("Error making GET request to secondaryHelloUrl: %v", err)
			} else {
				defer resp.Body.Close()
				slog.Info("GET request to secondaryHelloUrl successful")
			}
		}
	}
}
