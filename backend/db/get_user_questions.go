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
	Answers []QuestionAnswer
}

func GetUserQuestions(respondentID int) ([]UserQuestions, error) {
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
	
	getQuestionsStatement := `SELECT s.SpørsmålID, s.tekst
        FROM Spørsmål s
        LEFT JOIN Spørsmålsvar ss ON s.spørsmålID = ss.spørsmålID
        LEFT JOIN SvarVurdering sv ON ss.svarID = sv.svarID AND sv.respondentID = $1
        WHERE sv.vurderingID IS NULL
        LIMIT 5`

	stmt, err := db.Prepare(getQuestionsStatement)
	if err != nil {
		log.Fatalf("Error preparing statement: %v\n", err)
	}
	defer stmt.Close()

	var questionAnswers []UserQuestions

	rows, err := stmt.Query(respondentID)
	if err != nil {
		return questionAnswers, err
	}
	defer rows.Close()

	for rows.Next() {
		var question FormQuestion
		if err := rows.Scan(&question.QuestionID, &question.QuestionText); err != nil {
			return questionAnswers, err
		}

		getAnswersStatement := 
		`
		WITH TrueAnswer AS (
			SELECT ss.svarID
			FROM SpørsmålSvar ss
			WHERE ss.spørsmålID = $1
			  AND ss.chatgpt = TRUE
			  AND NOT EXISTS (
				SELECT 1
				FROM SvarVurdering sv
				WHERE sv.svarID = ss.svarID
				  AND sv.respondentID = $2
			  )
			LIMIT 1
		),
		FalseAnswer AS (
			SELECT ss.svarID
			FROM SpørsmålSvar ss
			WHERE ss.spørsmålID = $1
			  AND ss.chatgpt = FALSE
			  AND NOT EXISTS (
				SELECT 1
				FROM SvarVurdering sv
				WHERE sv.svarID = ss.svarID
				  AND sv.respondentID = $2
			  )
			LIMIT 1
		)
		SELECT ss.svarID, ss.spørsmålID, ss.chatgpt, ss.svartekst
		FROM SpørsmålSvar ss
		WHERE ss.svarID IN (SELECT svarID FROM TrueAnswer UNION ALL SELECT svarID FROM FalseAnswer);
		`
		
		stmt, err := db.Prepare(getAnswersStatement)
		if err != nil {
			log.Fatalf("Error preparing statement: %v\n", err)
		}
		defer stmt.Close()

		var answers []QuestionAnswer
		// var questionAnswers[5] UserQuestions

		rows, err := stmt.Query(question.QuestionID, respondentID)
		if err != nil {
			return questionAnswers, err
		}
		defer rows.Close()

		for rows.Next() {
			var answer QuestionAnswer
			if err := rows.Scan(&answer.AnswerID, &answer.QuestionID, &answer.IsChatGPT, &answer.AnswerText); err != nil {
				return questionAnswers, err
			}
			answers = append(answers, answer)
		}

		questionAnswers = append(questionAnswers, UserQuestions{Question: question, Answers: answers})
	}

	if err = rows.Err(); err != nil {
		return questionAnswers, err
	}

	return questionAnswers, nil
}