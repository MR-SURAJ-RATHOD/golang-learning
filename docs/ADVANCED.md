# Advanced Topics

## Concurrency (goroutines & channels)
```go
func worker(ch chan int) {
    for v := range ch {
        fmt.Println(v)
    }
}
// Start: go worker(ch)
```

### select
Use `select` to wait on multiple channels.

## Context
Use `context.Context` to manage timeouts and cancellation.

## Interfaces
Interfaces define behavior; small interfaces are idiomatic (e.g., `io.Reader`).

## Error handling
Wrap errors: `fmt.Errorf("read failed: %w", err)`

## Generics (since Go 1.18)
`func Map[T any](s []T, f func(T) T) []T { /*...*/ }` — use generics sparingly and where it improves clarity.