package postgres

import (
    "context"
    "database/sql"

    "lsc-golang/internal/domain"
)

type SummonerRepository struct {
    DB *sql.DB
}

func (r *SummonerRepository) Save(ctx context.Context, summoner domain.Summoner) error {
    _ = ctx
    _ = summoner
    return nil
}

func (r *SummonerRepository) FindByName(ctx context.Context, name string) (domain.Summoner, error) {
    _ = ctx
    _ = name
    return domain.Summoner{}, nil
}

func (r *SummonerRepository) FindByPUUID(ctx context.Context, puuid string) (domain.Summoner, error) {
    _ = ctx
    _ = puuid
    return domain.Summoner{}, nil
}
