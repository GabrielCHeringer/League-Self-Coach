package ratelimit

import (
    "context"
    "net/http"
)

type Limiter struct {}

func New() *Limiter {
    return &Limiter{}
}

func (l *Limiter) Wait(ctx context.Context) error {
    _ = ctx
    return nil
}

func (l *Limiter) UpdateFromHeaders(headers http.Header) {
    _ = headers
}
