package db

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

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

	// Write bugText to backup-file
	file, err := os.OpenFile("cmd/db/backups/bugreport_backup.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln("error opening file", err)
	}
	defer file.Close()

	// Prepare the row to be written
	rowSlice := []string{bugText}

	// Create a CSV writer and write the row
	csvwriter := csv.NewWriter(file)
	if err := csvwriter.Write(rowSlice); err != nil {
		log.Fatalln("error writing record to file", err)
	}

	// Flush the writer and check for errors
	csvwriter.Flush()
	if err := csvwriter.Error(); err != nil {
		log.Fatal(err)
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