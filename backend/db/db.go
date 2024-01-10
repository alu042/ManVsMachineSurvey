package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432 // This is the default port for PostgreSQL
	user     = "admin"
	password = "helse123"
	dbname   = "helseveileder"
)

func SetupDb() {
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

	// SQL statements to create tables
	createTableStatements := []string{
		`CREATE TABLE IF NOT EXISTS Spørsmål (
			spørsmålID SERIAL PRIMARY KEY,
			tekst TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS SpørsmålSvar (
			svarID SERIAL PRIMARY KEY,
			spørsmålID INT REFERENCES Spørsmål(spørsmålID),
			chatgpt BOOL NOT NULL,
			svartekst TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS Respondent (
			respondentID SERIAL PRIMARY KEY,
			alder TEXT NOT NULL,
			utdanningsgrad TEXT NOT NULL,
			helsepersonell BOOL NOT NULL,
			kjønn TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS SvarVurdering (
			vurderingID SERIAL PRIMARY KEY,
			respondentID INT REFERENCES Respondent(respondentID),
			svarID INT REFERENCES SpørsmålSvar(svarID),
			kunnskap INT,
			empati INT,
			hjelpsomhet INT
		);`,
		`CREATE TABLE IF NOT EXISTS Evaluering (
			evalueringtekst TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS FeilRapport (
			feiltekst TEXT
		);`,
	}

	// Execute SQL statements
	for _, stmt := range createTableStatements {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatalf("Error creating table: %v\n", err)
		}
	}

	fmt.Println("Tables created successfully.")
}
