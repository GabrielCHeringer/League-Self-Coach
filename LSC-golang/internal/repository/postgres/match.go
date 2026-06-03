package postgres

import (
    "context"
    "database/sql"

    "lsc-golang/internal/domain"
)

type MatchRepository struct {
    DB *sql.DB
}

func (r *MatchRepository) Save(ctx context.Context, match domain.Match) error {
    _ = ctx
    _ = match
    return nil
}

func (r *MatchRepository) FindByPUUID(ctx context.Context, puuid string) ([]domain.Match, error) {
    _ = ctx
    _ = puuid
    return nil, nil
}

func (r *MatchRepository) Exists(ctx context.Context, matchID string) (bool, error) {
    _ = ctx
    _ = matchID
    return false, nil
}
