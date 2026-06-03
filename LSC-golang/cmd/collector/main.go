package main

import (
    "log/slog"
    "time"

    "lsc-golang/configs"
)

func main() {
    cfg := configs.Load()

    ticker := time.NewTicker(time.Duration(cfg.CollectorIntervalSeconds) * time.Second)
    defer ticker.Stop()

    slog.Info("collector starting", "interval_seconds", cfg.CollectorIntervalSeconds)

    for range ticker.C {
        // TODO: call collector service
        slog.Info("collector tick")
    }
}
