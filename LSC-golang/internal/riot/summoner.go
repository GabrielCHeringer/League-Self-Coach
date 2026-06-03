package riot

import (
    "context"
    "errors"
)

func (c *Client) SummonerByName(ctx context.Context, name string) (SummonerDTO, error) {
    _ = ctx
    _ = name
    return SummonerDTO{}, errors.New("not implemented")
}

func (c *Client) SummonerByPUUID(ctx context.Context, puuid string) (SummonerDTO, error) {
    _ = ctx
    _ = puuid
    return SummonerDTO{}, errors.New("not implemented")
}
