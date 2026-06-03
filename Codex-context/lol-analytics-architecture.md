# lol-analytics — Arquitetura do Projeto

> Documento de referência para scaffolding e implementação.  
> Stack: **Go · PostgreSQL · Riot Games API (League of Legends)**

---

## Objetivo

Aplicação backend em Go que consome a Riot Games API para coletar dados de partidas e summoners do League of Legends, persiste no PostgreSQL e expõe endpoints HTTP para alimentar um dashboard público de análise (winrate, KDA, performance por campeão, etc.).

---

## Estrutura de Diretórios

```
lol-analytics/
├── cmd/
│   ├── api/
│   │   └── main.go          # Entrypoint do servidor HTTP
│   └── collector/
│       └── main.go          # Entrypoint do job de coleta de dados
│
├── internal/
│   ├── riot/                # Client da Riot Games API
│   │   ├── client.go        # HTTP client base com rate limiting
│   │   ├── summoner.go      # Endpoints: /summoner/v4
│   │   ├── match.go         # Endpoints: /match/v5
│   │   └── models.go        # Structs dos responses da Riot (raw)
│   │
│   ├── domain/              # Entidades internas do projeto
│   │   ├── summoner.go      # Entidade Summoner (limpa, sem ruído da API)
│   │   ├── match.go         # Entidade Match e Participant
│   │   └── stats.go         # Entidade de estatísticas agregadas
│   │
│   ├── repository/          # Camada de acesso ao banco de dados
│   │   ├── postgres/
│   │   │   ├── summoner.go  # CRUD de summoners
│   │   │   └── match.go     # CRUD de partidas
│   │   └── interfaces.go    # Interfaces dos repositórios (para injeção de dependência)
│   │
│   ├── service/             # Regras de negócio e análise
│   │   ├── analysis.go      # Cálculo de winrate, KDA, stats por campeão/patch
│   │   └── collector.go     # Orquestra: chama Riot API → transforma → persiste
│   │
│   └── handler/             # HTTP handlers (REST)
│       ├── summoner.go      # GET /summoners/:name
│       ├── match.go         # GET /matches/:puuid
│       └── stats.go         # GET /stats/:puuid
│
├── pkg/
│   └── ratelimit/
│       └── ratelimit.go     # Rate limiter que respeita headers da Riot API
│
├── migrations/              # Arquivos SQL versionados
│   ├── 001_create_summoners.sql
│   ├── 002_create_matches.sql
│   └── 003_create_participants.sql
│
├── configs/
│   └── config.go            # Carrega variáveis de ambiente
│
├── .env.example
├── go.mod
└── go.sum
```

---

## Responsabilidade de Cada Camada

### `cmd/`
Ponto de entrada da aplicação. Cada subdiretório gera um binário independente.

- **`cmd/api`** — sobe o servidor HTTP, registra rotas, injeta dependências.
- **`cmd/collector`** — job de coleta periódica (pode ser executado como cron). Não expõe HTTP.

### `internal/riot/`
Encapsula toda comunicação com a Riot Games API.

- `client.go` — instância do `http.Client` com timeout, headers de autenticação (`X-Riot-Token`) e integração com o rate limiter.
- `summoner.go` — busca summoner por nome ou PUUID via `/lol/summoner/v4`.
- `match.go` — lista e detalha partidas via `/lol/match/v5`.
- `models.go` — structs que espelham o JSON da Riot (verbosas, com todos os campos).

> ⚠️ **Nunca use os modelos da Riot diretamente fora deste pacote.** Eles são voláteis e poluídos.

### `internal/domain/`
Define as entidades que **o seu sistema** entende, independentes da Riot.

```go
// Exemplo: domain/stats.go
type ChampionStats struct {
    ChampionID int
    Winrate    float64
    KDA        float64
    GamesPlayed int
    AvgDamage  float64
}
```

### `internal/repository/`
Acesso ao PostgreSQL via `sqlx` ou `pgx`. Cada arquivo implementa as interfaces definidas em `interfaces.go`.

```go
// interfaces.go
type MatchRepository interface {
    Save(ctx context.Context, match domain.Match) error
    FindByPUUID(ctx context.Context, puuid string) ([]domain.Match, error)
}
```

### `internal/service/`
Lógica de negócio pura, sem dependência de HTTP ou banco diretamente.

- `collector.go` — dado um summoner, busca as últimas N partidas na Riot API, transforma para `domain.Match` e persiste via repositório.
- `analysis.go` — agrega dados do banco para calcular métricas (winrate por campeão, evolução de KDA por patch, etc.).

### `internal/handler/`
Handlers HTTP que recebem requisições, delegam para os services e retornam JSON.

```go
// Exemplo de rota
GET /summoners/:name          → busca summoner + stats gerais
GET /summoners/:name/matches  → histórico de partidas
GET /summoners/:name/stats    → análise agregada por campeão
```

### `pkg/ratelimit/`
Rate limiter reutilizável. A Riot retorna os limites nos headers de resposta:

```
X-App-Rate-Limit: 20:1,100:120
X-App-Rate-Limit-Count: 1:1,1:120
```

O limiter deve ler esses headers e se auto-ajustar, além de implementar retry com backoff exponencial em respostas `429 Too Many Requests`.

---

## Banco de Dados — Estrutura Principal

```sql
-- migrations/001_create_summoners.sql
CREATE TABLE summoners (
    id          SERIAL PRIMARY KEY,
    puuid       TEXT UNIQUE NOT NULL,
    name        TEXT NOT NULL,
    region      TEXT NOT NULL,
    summoner_level INT,
    updated_at  TIMESTAMP DEFAULT NOW()
);

-- migrations/002_create_matches.sql
CREATE TABLE matches (
    id          SERIAL PRIMARY KEY,
    match_id    TEXT UNIQUE NOT NULL,
    game_mode   TEXT,
    game_version TEXT,
    game_duration INT,
    played_at   TIMESTAMP
);

-- migrations/003_create_participants.sql
CREATE TABLE participants (
    id           SERIAL PRIMARY KEY,
    match_id     TEXT REFERENCES matches(match_id),
    puuid        TEXT REFERENCES summoners(puuid),
    champion_id  INT,
    champion_name TEXT,
    kills        INT,
    deaths       INT,
    assists      INT,
    win          BOOLEAN,
    total_damage INT,
    role         TEXT
);
```

---

## Stack de Bibliotecas

| Necessidade        | Biblioteca recomendada              |
|--------------------|-------------------------------------|
| HTTP server        | `github.com/go-chi/chi/v5`          |
| PostgreSQL driver  | `github.com/jackc/pgx/v5`           |
| Query builder      | `github.com/jmoiron/sqlx`           |
| Migrações SQL      | `github.com/golang-migrate/migrate` |
| Config / env       | `github.com/joho/godotenv`          |
| Rate limiting      | `golang.org/x/time/rate`            |
| Logs estruturados  | `log/slog` (stdlib Go 1.21+)        |

---

## Variáveis de Ambiente

```env
# .env.example
RIOT_API_KEY=RGAPI-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
RIOT_REGION=br1
RIOT_CLUSTER=americas

DATABASE_URL=postgres://user:password@localhost:5432/lol_analytics?sslmode=disable

SERVER_PORT=8080
COLLECTOR_INTERVAL=60   # segundos entre cada ciclo de coleta
```

---

## Fluxo de Dados

```
[Riot API]
    │
    ▼
internal/riot/          ← HTTP client + rate limiter
    │  (raw models)
    ▼
internal/service/collector.go   ← transforma para domain.*
    │  (domain models)
    ▼
internal/repository/postgres/   ← persiste no PostgreSQL
    │
    ▼
[PostgreSQL]
    │
    ▼
internal/service/analysis.go    ← agrega e calcula métricas
    │
    ▼
internal/handler/               ← expõe via REST JSON
    │
    ▼
[Dashboard / Cliente HTTP]
```

---

## Pontos de Atenção

1. **Rate Limiting da Riot** — o plano gratuito permite 20 req/s e 100 req/2min. Implemente fila + backoff antes de qualquer outra coisa.
2. **PUUID vs Summoner ID** — o Match v5 usa PUUID. A cadeia de chamadas é: `name → summonerId → puuid → matchIds → matchDetail`.
3. **Regiões vs Clusters** — endpoints de summoner usam região (`br1.api.riotgames.com`); endpoints de match usam cluster (`americas.api.riotgames.com`). Modele isso no client.
4. **Cache de partidas** — antes de buscar uma partida na API, verifique se o `match_id` já existe no banco. Partidas não mudam após encerradas.
5. **Versionamento de patch** — salve `game_version` em cada partida para permitir análise de performance por patch.
