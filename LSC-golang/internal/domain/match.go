package domain

import "time"

type Match struct {
    MatchID      string
    GameMode     string
    GameVersion  string
    GameDuration int
    PlayedAt     time.Time
    Participants []Participant
}

type Participant struct {
    PUUID        string
    ChampionID   int
    ChampionName string
    Kills        int
    Deaths       int
    Assists      int
    Win          bool
    TotalDamage  int
    Role         string
}
