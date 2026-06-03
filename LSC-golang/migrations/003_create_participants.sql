CREATE TABLE participants (
    id            SERIAL PRIMARY KEY,
    match_id      TEXT REFERENCES matches(match_id),
    puuid         TEXT REFERENCES summoners(puuid),
    champion_id   INT,
    champion_name TEXT,
    kills         INT,
    deaths        INT,
    assists       INT,
    win           BOOLEAN,
    total_damage  INT,
    role          TEXT
);
