# Module 11 — Exercises

## Exercise 1: Structured Logging

Replace `fmt.Println` with `log/slog` JSON logger.

## Exercise 2: Health Endpoint

Add HTTP server on `PORT` with `GET /health` returning `{"status":"ok"}`.

## Exercise 3: Lambda Handler

Create `lambda/main.go`:
```go
func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
```

## Exercise 4: Config Validation

Fail fast at startup if `JWT_SECRET` or `DATABASE_URL` is missing in production mode (`ENV=production`).
