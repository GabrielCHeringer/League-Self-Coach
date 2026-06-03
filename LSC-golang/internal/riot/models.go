package riot

type SummonerDTO struct {
    ID            string `json:"id"`
    AccountID     string `json:"accountId"`
    PUUID         string `json:"puuid"`
    Name          string `json:"name"`
    SummonerLevel int    `json:"summonerLevel"`
}

type MatchlistDTO struct {
    MatchIDs []string `json:"matchIds"`
}

type MatchDTO struct {
    Metadata MatchMetadataDTO `json:"metadata"`
    Info     MatchInfoDTO     `json:"info"`
}

type MatchMetadataDTO struct {
    MatchID      string   `json:"matchId"`
    Participants []string `json:"participants"`
}

type MatchInfoDTO struct {
    GameMode     string               `json:"gameMode"`
    GameVersion  string               `json:"gameVersion"`
    GameDuration int                  `json:"gameDuration"`
    GameCreation int64                `json:"gameCreation"`
    Participants []MatchParticipantDTO `json:"participants"`
}

type MatchParticipantDTO struct {
    PUUID                         string `json:"puuid"`
    ChampionID                    int    `json:"championId"`
    ChampionName                  string `json:"championName"`
    Kills                         int    `json:"kills"`
    Deaths                        int    `json:"deaths"`
    Assists                       int    `json:"assists"`
    Win                           bool   `json:"win"`
    TotalDamageDealtToChampions   int    `json:"totalDamageDealtToChampions"`
    TeamPosition                  string `json:"teamPosition"`
}
