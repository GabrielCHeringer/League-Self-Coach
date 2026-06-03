package riot

import (
    "net/http"
    "time"

    "lsc-golang/pkg/ratelimit"
)

type Client struct {
    HTTPClient  *http.Client
    APIKey      string
    Region      string
    Cluster     string
    RateLimiter *ratelimit.Limiter
}

func NewClient(apiKey, region, cluster string, limiter *ratelimit.Limiter) *Client {
    httpClient := &http.Client{Timeout: 10 * time.Second}
    if limiter == nil {
        limiter = ratelimit.New()
    }

    return &Client{
        HTTPClient:  httpClient,
        APIKey:      apiKey,
        Region:      region,
        Cluster:     cluster,
        RateLimiter: limiter,
    }
}
