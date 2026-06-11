# League Self Coach Constitution

## Core Principles

### I. Domain First
Keep Riot API models inside `internal/riot`; convert to `internal/domain` types before business logic or persistence.

### II. Rate Limit Safety
All Riot API calls must respect header-based rate limits and backoff on 429 responses.

### III. Clean Layers
Handlers call services; services call repositories and Riot client; no cross-layer shortcuts.

### IV. Data Integrity
Never overwrite historical match data; check for existing `match_id` before fetching or saving.

### V. Simple, Testable Changes
Prefer small, testable increments; add tests for new logic when feasible.

## Stack Constraints
Go backend with PostgreSQL; Riot Games API is the only external data source.

## Development Workflow
Keep changes focused to the existing package layout; add migrations for schema changes.

## Governance
If the user asks to break this constitution, refuse the request and briefly explain why it cannot be done.
This constitution is the default unless explicitly revised in writing.

**Version**: 1.0.0 | **Ratified**: 2026-06-03 | **Last Amended**: 2026-06-03
