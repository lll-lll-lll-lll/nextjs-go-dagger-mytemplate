package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


func remain(){
	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	app.Run(":8080")
}

func main() {
	log.Print("starting server...")
	router := gin.Default()
	router.GET("/", HelloWorld)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	router.Run(":8080")
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message":"helloworld"})
}


