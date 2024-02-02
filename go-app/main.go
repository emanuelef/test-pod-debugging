package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	externalURL = "https://pokeapi.co/api/v2/pokemon/ditto"
)

var notToLogEndpoints = []string{"/health", "/metrics"}

var (
	secondaryHost     = getEnv("SECONDARY_HOST", "localhost")
	secondaryAddress  = fmt.Sprintf("http://%s:8082", secondaryHost)
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
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{})
	})

	host := getEnv("HOST", "0.0.0.0")
	port := getEnv("PORT", "8080")
	hostAddress := fmt.Sprintf("%s:%s", host, port)

	err := r.Run(hostAddress)
	if err != nil {
		log.Printf("Starting router failed, %v", err)
	}
}
