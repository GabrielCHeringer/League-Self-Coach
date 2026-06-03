package service

import (
    "context"

    "lsc-golang/internal/domain"
    "lsc-golang/internal/repository"
)

type AnalysisService struct {
    Matches repository.MatchRepository
}

func (s *AnalysisService) ChampionStatsByPUUID(ctx context.Context, puuid string) ([]domain.ChampionStats, error) {
    _ = ctx
    _ = puuid
    return nil, nil
}
