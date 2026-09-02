# Module 06 — Exercises

## Exercise 1: Search Endpoint

Add `GET /albums/search?artist=Miles` that filters albums by artist name.

## Exercise 2: Auth Middleware Stub

Create middleware that checks for header `X-API-Key: secret123`.
Return 401 if missing. Apply only to POST/PUT/DELETE routes.

## Exercise 3: Handler Tests

Create `main_test.go` using `httptest`:
```go
w := httptest.NewRecorder()
req, _ := http.NewRequest("GET", "/ping", nil)
router.ServeHTTP(w, req)
```

## Exercise 4: Connect Database

Replace in-memory `albums` slice with SQLite repository from Module 05.
