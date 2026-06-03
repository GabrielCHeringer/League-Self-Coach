package main

import (
    "log/slog"
    "net/http"
    "os"
    "time"

    "lsc-golang/configs"
)

func main() {
    cfg := configs.Load()

    // TODO: wire handlers and routes.
    srv := &http.Server{
        Addr:         ":" + cfg.ServerPort,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    slog.Info("api starting", "port", cfg.ServerPort)

    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        slog.Error("api stopped", "error", err)
        os.Exit(1)
    }
}
