package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type FormAnswer struct {
    AnswerID int
    Ratings  []int
}

type AllFormAnswers []FormAnswer

func InsertUserAnswers(respondentId int, allAnswers string) (error) {

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	var convertedFormAnswers = convertAnswerFormat(allAnswers)

	for _,formanswer := range convertedFormAnswers {
		insertStatement := `
		INSERT INTO SvarVurdering (svarID, respondentID, kunnskap, empati, hjelpsomhet)
		VALUES ($1, $2, $3, $4, $5)
		`
		stmt, err := db.Prepare(insertStatement)
		if err != nil {
			log.Fatalf("Error preparing statement: %v\n", err)
		}
		defer stmt.Close()

		kunnskap := formanswer.Ratings[0]
		empati := formanswer.Ratings[1]
		hjelpsomhet := formanswer.Ratings[2]

		_, err = stmt.Exec(formanswer.AnswerID, respondentId, kunnskap, empati, hjelpsomhet)
		if err != nil {
			log.Fatalf("Error executing statement: %v\n", err)
		}
	}

	fmt.Print("Inserted data successfully")

	return nil
}

func convertAnswerFormat(allAnswers string) (AllFormAnswers) {
	// Define an intermediate structure that matches the input format
    var intermediate [][]interface{}

    // Parse the JSON string into the intermediate structure
    err := json.Unmarshal([]byte(allAnswers), &intermediate)
    if err != nil {
        log.Fatalf("Error parsing JSON: %v", err)
    }

    // Convert the intermediate structure to OuterSlice
    var result AllFormAnswers
    for _, pair := range intermediate {
        if len(pair) != 2 {
            log.Fatalf("Invalid pair length: %v", pair)
        }

        number, ok := pair[0].(float64) // JSON numbers are parsed as float64
        if !ok {
            log.Fatalf("Expected a number, got: %v", pair[0])
        }

        arrayInterface, ok := pair[1].([]interface{})
        if !ok {
            log.Fatalf("Expected an array, got: %v", pair[1])
        }

        var array []int
        for _, item := range arrayInterface {
            num, ok := item.(float64)
            if !ok {
                log.Fatalf("Expected a number in array, got: %v", item)
            }
            array = append(array, int(num))
        }

        result = append(result, FormAnswer{AnswerID: int(number), Ratings: array})
    }

	return result
}