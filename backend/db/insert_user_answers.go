package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InsertUserAnswers(respondentId int, allAnswers string) (int, error) {
	// Parse allAnswers from JSON string to map
	var allAnswersArray []string
	err := json.Unmarshal([]byte(allAnswers), &allAnswersArray)
	if err != nil {
		return 0, fmt.Errorf("invalid JSON format for allAnswers: %v", err)
	}

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

	insertStatement := `
	INSERT INTO SvarVurdering (svarID, respondentID, kunnskap, empati, hjelpsomhet)
	VALUES ($1, $2, $3, $4, $5)
	`

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error creating transaction: %v\n", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("Transaction commit failed: %v\n", err)
	}

	fmt.Printf("Data inserted successfully for respondentID: %d\n", respondentId)
	return respondentId, nil
}