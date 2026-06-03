package repository

import (
    "context"

    "lsc-golang/internal/domain"
)

type SummonerRepository interface {
    Save(ctx context.Context, summoner domain.Summoner) error
    FindByName(ctx context.Context, name string) (domain.Summoner, error)
    FindByPUUID(ctx context.Context, puuid string) (domain.Summoner, error)
}

type MatchRepository interface {
    Save(ctx context.Context, match domain.Match) error
    FindByPUUID(ctx context.Context, puuid string) ([]domain.Match, error)
    Exists(ctx context.Context, matchID string) (bool, error)
}
