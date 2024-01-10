package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InsertUserData(age string, education string, healthcarepersonell bool, gender string, answeredbefore bool, county string, submitdate string) (int, error) {
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
		INSERT INTO Respondent (alder, utdanningsgrad, helsepersonell, kjønn, svartfør, fylke, dato)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING respondentID
	`
	
	stmt, err := db.Prepare(insertStatement)
	if err != nil {
		log.Fatalf("Error preparing statement: %v\n", err)
	}
	defer stmt.Close()

	var respondentID int
	err = stmt.QueryRow(age, education, healthcarepersonell, gender, answeredbefore, county, submitdate).Scan(&respondentID)
	if err != nil {
		log.Fatalf("Error inserting userdata: %v\n", err)
	}

	fmt.Printf("Data inserted successfully with respondentID: %d\n", respondentID)

	return respondentID, nil
}