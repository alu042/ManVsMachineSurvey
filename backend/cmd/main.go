package main

import (
	"fmt"
	"net/http"

	"helseveileder/db"

	"github.com/gin-gonic/gin"
)

type FormData struct {
    Age                string `json:"age"`
    Education          string `json:"education"`
    HealthcarePersonnel bool   `json:"healthcare_personnel"`
    Gender             string `json:"gender"`
}

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

        var requestBody FormData

        if err := c.BindJSON(&requestBody); err != nil {
            fmt.Print(err)
        }

        fmt.Print(requestBody)

        db.InsertData(requestBody.Age, requestBody.Education, requestBody.HealthcarePersonnel, requestBody.Gender)
    })

    // Run the server on port 8080
    router.Run(":8080")
    // db.SetupDb()
}
