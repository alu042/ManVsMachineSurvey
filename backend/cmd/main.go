package main

import (
	"fmt"
	"net/http"
	"strconv"

	"helseveileder/cmd/db"

	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

type UserformData struct {
    Age                string `json:"age"`
    Education          string `json:"education"`
    HealthcarePersonnel bool  `json:"healthcare_personnel"`
    IsLicensed         bool  `json:"is_licensed"`
    Gender             string `json:"gender"`
    AnsweredBefore     bool `json:"answered_before"`
    County             string `json:"county"`
    SubmitDate         string `json:"submit_date"`
}

type FormData struct {
    FormAnswers        string `json:"allFormAnswers"`
    RespondentId       int    `json:"respondentID"`
}

type BugReport struct {
    BugText string `json:"bugText"`
}

type Evaluation struct {
    EvaluationText string `json:"evaluationText"`
}

func main() {
    router := gin.Default()
    // router.Use(cors.Default())

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:5173"} // Add your frontend URL here
    config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
    router.Use(cors.New(config))

    router.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"hello":"world"})
    }) 
    
    // Info about user
    router.POST("/submituserform", func(c *gin.Context) {

        var requestBody UserformData

        if err := c.BindJSON(&requestBody); err != nil {
            fmt.Print(err)
        }

        // Capture both the ID and error returned from InsertData
		respondentId, err := db.InsertUserData(requestBody.Age, requestBody.Education, requestBody.HealthcarePersonnel, requestBody.IsLicensed, requestBody.Gender, requestBody.AnsweredBefore, requestBody.County, requestBody.SubmitDate)
		
        if err != nil {
			fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to insert data"})
			return
		}

        // Respond with the ID of the newly inserted respondent
		c.JSON(http.StatusOK, gin.H{"respondentID": respondentId})
    })

    // Get questions & answers from database
    router.GET("/userquestions", func(c *gin.Context) {
        respondentID, err := strconv.Atoi(c.Query("respondentID"))

        if err != nil {
            fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong respondentID-format (should be int)."})
			return
        }

        questions, err := db.GetUserQuestions(respondentID)

        if err != nil {
            fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting questions for given user."})
			return
        }

        c.JSON(http.StatusOK, gin.H{"questions": questions})
    })

    router.POST("/submitanswers", func(c *gin.Context) {
        var requestBody FormData

        if err := c.BindJSON(&requestBody); err != nil {
            fmt.Print(err)
        }

        err := db.InsertUserAnswers(requestBody.RespondentId, requestBody.FormAnswers)
        
        if err != nil {
			fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to insert data"})
			return
		}

        // Respond with the ID of the newly inserted respondent
		c.JSON(http.StatusOK, "Successfully inserted formdata!")
    })

    router.POST("/submitbug", func(c *gin.Context) {
        var requestBody BugReport

        // Bind JSON to the requestBody struct
        if err := c.BindJSON(&requestBody); err != nil {
            fmt.Print(err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }

        // Here, you'd insert the bug text into your database
        err := db.InsertBugReport(requestBody.BugText)
        if err != nil {
            fmt.Print(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to insert bug report"})
            return
        }

        // Respond with a success message
        c.JSON(http.StatusOK, "Successfully submitted bug report!")
    })

    router.POST("/submiteval", func(c *gin.Context) {
        var requestBody Evaluation

        // Bind JSON to the requestBody struct
        if err := c.BindJSON(&requestBody); err != nil {
            fmt.Print(err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }

        // Here, you'd insert the bug text into your database
        err := db.InsertEvaluation(requestBody.EvaluationText)
        if err != nil {
            fmt.Print(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to insert evaluation"})
            return
        }

        // Respond with a success message
        c.JSON(http.StatusOK, "Successfully submitted evaluation!")
    })

    // Run the server on port 8080
    db.SetupDb()
    router.Run(":8080")
}
