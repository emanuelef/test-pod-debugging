package main

import (
	"fmt"
	"log"
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
	apiHost         = getEnv("API_HOST", "localhost")
	apiHostAddress  = fmt.Sprintf("http://%s:8080", apiHost)
	helloServiceUrl = fmt.Sprintf("%s/hello/john", apiHostAddress)
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
		log.Println("Called /health")
		c.JSON(http.StatusNoContent, gin.H{})
	})

	// Define a hello endpoint
	r.GET("/hello", func(c *gin.Context) {
		log.Println("Called /hello")
		c.JSON(http.StatusNoContent, gin.H{})
	})

	// Run the router in a goroutine
	go func() {
		host := getEnv("HOST", "0.0.0.0")
		port := getEnv("PORT", "8080")
		hostAddress := fmt.Sprintf("%s:%s", host, port)

		log.Println("Starting server")
		err := r.Run(hostAddress)
		if err != nil {
			log.Println("Starting router failed, %v", err)
		}
	}()

	// Start a ticker to perform periodic calls
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Perform the GET request to helloServiceUrl here
			resp, err := http.Get(helloServiceUrl)
			if err != nil {
				log.Printf("Error making GET request to helloServiceUrl %v\n", err)
			} else {
				defer resp.Body.Close()
				log.Printf("GET request to helloServiceUrl %s", resp.Status)
			}
		}
	}
}
