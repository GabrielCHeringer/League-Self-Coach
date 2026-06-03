CREATE TABLE matches (
    id            SERIAL PRIMARY KEY,
    match_id      TEXT UNIQUE NOT NULL,
    game_mode     TEXT,
    game_version  TEXT,
    game_duration INT,
    played_at     TIMESTAMP
);
