package service

import (
    "context"

    "lsc-golang/internal/repository"
    "lsc-golang/internal/riot"
)

type CollectorService struct {
    Riot      *riot.Client
    Summoners repository.SummonerRepository
    Matches   repository.MatchRepository
}

func (s *CollectorService) CollectSummoner(ctx context.Context, summonerName string) error {
    _ = ctx
    _ = summonerName
    return nil
}
