package riot

import (
    "context"
    "errors"
)

func (c *Client) MatchIDsByPUUID(ctx context.Context, puuid string, count int) ([]string, error) {
    _ = ctx
    _ = puuid
    _ = count
    return nil, errors.New("not implemented")
}

func (c *Client) MatchByID(ctx context.Context, matchID string) (MatchDTO, error) {
    _ = ctx
    _ = matchID
    return MatchDTO{}, errors.New("not implemented")
}
