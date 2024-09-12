CREATE TABLE IF NOT EXISTS Spørsmål (spørsmålID SERIAL PRIMARY KEY,tekst TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS SpørsmålSvar (svarID SERIAL PRIMARY KEY,spørsmålID INT REFERENCES Spørsmål(spørsmålID),chatgpt BOOL NOT NULL,svartekst TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS Respondent (respondentID SERIAL PRIMARY KEY,alder TEXT NOT NULL,utdanningsgrad TEXT NOT NULL,helsepersonell BOOL NOT NULL,harlisens BOOL NOT NULL,kjønn TEXT NOT NULL,svartfør BOOL NOT NULL,fylke TEXT NOT NULL,dato TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS SvarVurdering (vurderingID SERIAL PRIMARY KEY,respondentID INT REFERENCES Respondent(respondentID),svarID INT REFERENCES SpørsmålSvar(svarID),kunnskap INT,empati INT,hjelpsomhet INT);
CREATE TABLE IF NOT EXISTS Evaluering (evalueringtekst TEXT);
CREATE TABLE IF NOT EXISTS FeilRapport (feiltekst TEXT);

-- Insert sample questions
INSERT INTO Spørsmål (tekst) VALUES
('What are the common symptoms of the flu?'),
('How can I prevent catching a cold?'),
('What is the recommended daily water intake?'),
('How often should I exercise?'),
('What are the benefits of a balanced diet?');

-- Insert sample answers (both ChatGPT and non-ChatGPT)
INSERT INTO SpørsmålSvar (spørsmålID, chatgpt, svartekst) VALUES
(1, true, 'Common flu symptoms include fever, cough, sore throat, runny or stuffy nose, body aches, headache, fatigue, and sometimes vomiting and diarrhea.'),
(1, false, 'Flu symptoms typically include high fever, severe body aches, fatigue, and respiratory symptoms like coughing and congestion.'),
(2, true, 'To prevent catching a cold: wash your hands frequently, avoid touching your face, stay hydrated, get enough sleep, eat a balanced diet, and avoid close contact with sick people.'),
(2, false, 'Prevent colds by maintaining good hygiene, boosting your immune system with vitamins, and avoiding crowded places during cold season.'),
(3, true, 'The general recommendation is to drink about 8 cups (64 ounces) of water per day, but this can vary based on individual needs, climate, and activity level.'),
(3, false, 'Aim for 2-3 liters of water daily, adjusting based on your activity level and climate.'),
(4, true, 'Adults should aim for at least 150 minutes of moderate-intensity aerobic activity or 75 minutes of vigorous-intensity aerobic activity per week, along with muscle-strengthening activities at least 2 days a week.'),
(4, false, 'Try to exercise for 30 minutes a day, 5 days a week, combining cardio and strength training for best results.'),
(5, true, 'A balanced diet provides essential nutrients, supports immune function, helps maintain a healthy weight, reduces the risk of chronic diseases, improves heart health, and promotes overall well-being.'),
(5, false, 'Eating a variety of foods ensures you get all necessary nutrients, helps maintain a healthy weight, and reduces risk of various health issues.');