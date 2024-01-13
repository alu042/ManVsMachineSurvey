package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
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
			kjønn TEXT NOT NULL,
			svartfør BOOL NOT NULL,
			fylke TEXT NOT NULL,
			dato TEXT NOT NULL
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
		`INSERT INTO Spørsmål (spørsmålID, tekst) VALUES
		(1, 'Jeg har noen røde prikker på penishodet, som kan klø. Hva kommer dette av?'),
		(2, 'Jeg føler meg svimmel når jeg står opp, er det normalt?'),
		(3, 'Hvor ofte bør jeg ta en blodtrykkstest?'),
		(4, 'Jeg har vedvarende hoste, bør jeg være bekymret?'),
		(5, 'Hva er de beste måtene å forebygge hjertesykdom på?'),
		(6, 'Hvor lenge varer en vanlig influensasymptomer?'),
		(7, 'Er det normalt å føle seg trøtt etter å ha startet en ny medisin?'),
		(8, 'Hvordan kan jeg vite om jeg har diabetes?'),
		(9, 'Hvilke vaksiner bør jeg vurdere å få som voksen?'),
		(10, 'Er det mulig å ha for høyt nivå av vitaminer?'
		);`,
		`INSERT INTO SpørsmålSvar (svarID, spørsmålID, chatgpt, svarTekst) VALUES
		(1, 1, false, 'Røde prikker kan skyldes irritasjon, talgkjertler, kjønnsvorter, spicula glandis, andre utslett eller soppinfeksjon. Det er lurt å oppsøke fastlege for å finne ut av det.'),
		(2, 2, true, 'Svimmelhet ved oppreisning kan skyldes blodtrykksfall. Det kalles ortostatisk hypotensjon. Det er viktig å stå opp sakte, og hvis det vedvarer, kontakt lege.'),
		(3, 3, true, 'Voksne bør ha blodtrykkstest minst en gang hvert annet år. Ved høyt blodtrykk eller andre risikofaktorer, oftere.'),
		(4, 4, false, 'Vedvarende hoste kan være tegn på en rekke tilstander, noen mer alvorlige enn andre. Det er viktig å konsultere lege.'),
		(5, 5, true, 'Forebygging av hjertesykdom inkluderer et sunt kosthold, regelmessig fysisk aktivitet, ikke røyke og moderere alkoholinntaket.'),
		(6, 6, false, 'Influensasymptomer varer vanligvis fra noen dager til to uker. Hvis symptomene forverres eller ikke forbedres, kontakt lege.'),
		(7, 7, true, 'Trøtthet kan være en bivirkning av mange medisiner. Snakk med legen din om symptomene dine, spesielt hvis de er alvorlige.'),
		(8, 8, false, 'Diabetesdiagnose stilles gjennom blodsukkermålinger. Symptomer inkluderer tørste, hyppig vannlating, tretthet og tåkesyn.'),
		(9, 9, true, 'Voksne bør vurdere å få stivkrampe, difteri, kikhoste, influensa og pneumokokkvaksiner, blant andre.'),
		(10, 10, false, 'Ja, det er mulig å ha for høyt nivå av vitaminer, spesielt fra kosttilskudd. Det er viktig å følge anbefalte doser.'),
		(11, 1, true, 'Dette kan være en mild betennelse eller infeksjon. Unngå å irritere området, og hvis det ikke forbedres, ta kontakt med lege for en sjekk.'),
		(12, 2, false, 'Enkelte kan oppleve svimmelhet hvis de reiser seg for fort på grunn av blodtrykksendringer. Hvis dette skjer ofte, bør du diskutere det med fastlegen.'),
		(13, 3, false, 'Det anbefales at voksne får målt blodtrykket sitt minst hvert femte år. Hvis du har risikofaktorer for hjertesykdom, kanskje årlig.'),
		(14, 4, true, 'Langvarig hoste kan være tegn på kroniske tilstander som astma eller GERD. En lege kan hjelpe deg med å diagnostisere årsaken.'),
		(15, 5, false, 'Unngå tobakksrøyk, spis et kosthold rikt på frukt og grønnsaker, vær fysisk aktiv, og kontroller din kolesterol og blodtrykk.'),
		(16, 6, true, 'Influensasymptomer varer typisk i 1-2 uker. Hvis du opplever alvorlige symptomer eller de varer lenger enn to uker, søk medisinsk hjelp.'),
		(17, 7, false, 'Tretthet kan være en vanlig bivirkning av nye medisiner. Hvis det vedvarer, snakk med legen din for å se om doseringen trenger justering.'),
		(18, 8, true, 'Tegn på diabetes kan inkludere økt tørste og hyppig vannlating. Blodprøver kan bekrefte diagnosen.'),
		(19, 9, false, 'Det er viktig å holde seg oppdatert på vaksinasjoner. Ta en samtale med legen din om influensa, pneumokokk, og andre relevante vaksiner.'),
		(20, 10, true, 'For mye av visse vitaminer, som A, D, E og K, kan være skadelig. Disse er fettløselige og lagres i kroppen. Følg anbefalte doseringer.'
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
