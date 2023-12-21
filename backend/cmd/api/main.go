package main

import (
	"fmt"
	"io"
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

    // Info about user
    router.POST("/submitform", func(c *gin.Context) {
        jsonData, err := io.ReadAll(c.Request.Body)
        if err != nil {
            // Handle error
        }
        fmt.Print(string(jsonData))
    })

    // Run the server on port 8080
    router.Run(":8080")
}
