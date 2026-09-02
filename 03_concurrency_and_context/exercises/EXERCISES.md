# Module 03 — Exercises

## Exercise 1: Worker Pool

Modify `lesson03WorkerPool` to use **5 workers** and **50 jobs**.
Print total successful jobs at the end.

## Exercise 2: Fan-Out / Fan-In

Create `exercises/fanout.go`:
- Fan-out: distribute 20 numbers to 4 goroutines
- Each goroutine squares its numbers
- Fan-in: merge results into one channel
- Print all squared values

## Exercise 3: Context Cancel

Use `context.WithCancel`:
```go
ctx, cancel := context.WithCancel(context.Background())
// call cancel() after 2 seconds from another goroutine
```
Workers should stop gracefully when context is cancelled.

## Exercise 4: Race Detector

Write code that increments a counter in 1000 goroutines **without** mutex.
Run `go run -race .` and fix it.
