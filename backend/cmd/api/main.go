package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Define a basic GET request handler
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    // Run the server on port 8080
    router.Run(":8080")
}
