package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type FormQuestion struct {
	QuestionID int
	QuestionText string
}

type QuestionAnswer struct {
	AnswerID int
	QuestionID int
	IsChatGPT bool
	AnswerText string
}

type UserQuestions struct {
	Question FormQuestion
	Answers [2]QuestionAnswer
}

func GetUserQuestions(respondendID int) ([5]UserQuestions, error) {
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
	
	getQuestionsStatement := fmt.Sprintf("SELECT s.SpørsmålID, s.tekst"+
		"FROM Spørsmål s"+
		"LEFT JOIN Spørsmålsvar ss ON s.spørsmålID == ss.spørsmålID"+
		"LEFT JOIN SvarVurdering sv ON ss.svarID == sv.svarID AND sv.respondentID = :$respondendID"+
		"WHERE sv.vurderingID IS NULL"+
		"LIMIT 5", 
		respondendID)

	stmt, err := db.Prepare(getQuestionsStatement)
	if err != nil {
		log.Fatalf("Error preparing statement: %v\n", err)
	}
	defer stmt.Close()

	var questions [5]FormQuestion

	rows, err := stmt.Query(getQuestionsStatement)

	

}