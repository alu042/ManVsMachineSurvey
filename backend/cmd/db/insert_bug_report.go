package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InsertBugReport(bugText string) (error) {
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
	INSERT INTO FeilRapport (feiltekst)
	VALUES ($1)
	`
	stmt, err := db.Prepare(insertStatement)
	if err != nil {
		log.Fatalf("Error preparing statement: %v\n", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(bugText)
	if err != nil {
		log.Fatalf("Error executing statement: %v\n", err)
	}

	fmt.Print("Inserted bug successfully")

	return nil
}
