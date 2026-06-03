package configs

import (
    "os"
    "strconv"
)

type Config struct {
    RiotAPIKey               string
    RiotRegion               string
    RiotCluster              string
    DatabaseURL              string
    ServerPort               string
    CollectorIntervalSeconds int
}

func Load() Config {
    return Config{
        RiotAPIKey:               os.Getenv("RIOT_API_KEY"),
        RiotRegion:               envOrDefault("RIOT_REGION", "br1"),
        RiotCluster:              envOrDefault("RIOT_CLUSTER", "americas"),
        DatabaseURL:              os.Getenv("DATABASE_URL"),
        ServerPort:               envOrDefault("SERVER_PORT", "8080"),
        CollectorIntervalSeconds: envOrDefaultInt("COLLECTOR_INTERVAL", 60),
    }
}

func envOrDefault(key, fallback string) string {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }
    return value
}

func envOrDefaultInt(key string, fallback int) int {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }
    parsed, err := strconv.Atoi(value)
    if err != nil {
        return fallback
    }
    return parsed
}
