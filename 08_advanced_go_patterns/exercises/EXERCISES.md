# Module 08 — Exercises

## Exercise 1: Generic Queue

Implement `Queue[T]` with `Enqueue(T)` and `Dequeue() (T, bool)`.

## Exercise 2: HTTP Client Options

Create `NewHTTPClient` with functional options:
- `WithTimeout(d time.Duration)`
- `WithBaseURL(url string)`
- `WithRetry(count int)`

## Exercise 3: Middleware for Gin

Port the `chain` pattern to Gin middleware — logging + request ID.

## Exercise 4: Avoid Reflection

Rewrite `inspectStruct` using an interface `Describer` with method `Describe() string` instead of reflection.
