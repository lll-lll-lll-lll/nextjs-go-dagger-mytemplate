package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


func main() {
	log.Print("starting server...")
	router := gin.Default()
	router.GET("/api/hello", HelloWorld)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "7071"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	router.Run(":7071")
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message":"helloworld"})
}


