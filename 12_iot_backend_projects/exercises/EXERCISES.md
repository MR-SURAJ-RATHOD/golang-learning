# Module 12 — Exercises

## Exercise 1: Separate Binaries

Split into `cmd/publisher/main.go` and `cmd/subscriber/main.go`.
Run them in two terminals.

## Exercise 2: Store Readings

Save each reading to SQLite (from Module 05) with timestamp.

## Exercise 3: WebSocket Dashboard

Add WebSocket endpoint that streams latest temperature to a browser.

## Exercise 4: Email Alert

When temp > 28°C, log alert with timestamp and send to a mock notification function.

## Exercise 5: Scale Workers

Make worker count configurable via `WORKERS=5` environment variable.
