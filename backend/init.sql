CREATE TABLE IF NOT EXISTS Spørsmål (spørsmålID SERIAL PRIMARY KEY,tekst TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS SpørsmålSvar (svarID SERIAL PRIMARY KEY,spørsmålID INT REFERENCES Spørsmål(spørsmålID),chatgpt BOOL NOT NULL,svartekst TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS Respondent (respondentID SERIAL PRIMARY KEY,alder TEXT NOT NULL,utdanningsgrad TEXT NOT NULL,helsepersonell BOOL NOT NULL,harlisens BOOL NOT NULL,kjønn TEXT NOT NULL,svartfør BOOL NOT NULL,fylke TEXT NOT NULL,dato TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS SvarVurdering (vurderingID SERIAL PRIMARY KEY,respondentID INT REFERENCES Respondent(respondentID),svarID INT REFERENCES SpørsmålSvar(svarID),kunnskap INT,empati INT,hjelpsomhet INT);
CREATE TABLE IF NOT EXISTS Evaluering (evalueringtekst TEXT);
CREATE TABLE IF NOT EXISTS FeilRapport (feiltekst TEXT);