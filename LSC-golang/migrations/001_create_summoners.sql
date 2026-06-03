CREATE TABLE summoners (
    id             SERIAL PRIMARY KEY,
    puuid          TEXT UNIQUE NOT NULL,
    name           TEXT NOT NULL,
    region         TEXT NOT NULL,
    summoner_level INT,
    updated_at     TIMESTAMP DEFAULT NOW()
);
