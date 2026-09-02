# 03. Concurrency & Context

> **Time:** 5–7 hours | **Prerequisites:** [Module 02](../02_modules_and_workspace/README.md)

Master Go's concurrency model — goroutines, channels, mutexes, and context for cancellation.

---

## What You Will Learn

- [ ] Goroutines — lightweight threads
- [ ] `sync.WaitGroup` — wait for goroutines to finish
- [ ] `sync.Mutex` — protect shared data from race conditions
- [ ] Channels — communicate between goroutines
- [ ] `select` — multiplex channel operations
- [ ] `context.Context` — cancellation and timeouts
- [ ] Worker pool pattern

---

## Step-by-Step Lessons

| Step | Function | Topic |
| :---: | :--- | :--- |
| 1 | `lesson01Goroutines()` | Launch goroutines, WaitGroup |
| 2 | `lesson02Mutex()` | Race condition fix with Mutex |
| 3 | `lesson03WorkerPool()` | Channels + worker pool + context timeout |

### Step 1 — Run the code

```bash
cd 03_concurrency_and_context
go run main.go
```

### Step 2 — Detect race conditions

Remove `mu.Lock()` / `mu.Unlock()` from `lesson02Mutex` and run:

```bash
go run -race main.go
```

You will see a **DATA RACE** warning. Always use `-race` during development.

### Step 3 — Understand worker pool

Read `lesson03WorkerPool()`:
1. Context with 3-second timeout
2. Buffered job channel (10 slots)
3. 3 workers process jobs concurrently
4. Results collected via results channel

---

## Key Rules

| Rule | Why |
| :--- | :--- |
| Don't communicate by sharing memory | Share memory by communicating (channels) |
| Always `defer cancel()` after `context.WithTimeout` | Prevents context leak |
| Close channels from sender side only | Receiver uses `range` |
| Use `WaitGroup` for goroutine lifecycle | Know when all work is done |

---

## Exercises

See [exercises/EXERCISES.md](./exercises/EXERCISES.md)

1. Build a worker pool: 5 workers, 50 jobs
2. Implement fan-out/fan-in pattern
3. Use `context.WithCancel` to stop workers on Ctrl+C

---

## Checkpoint

Before Module 04:

- [ ] Launch 10 goroutines and wait with WaitGroup
- [ ] Explain buffered vs unbuffered channels
- [ ] Fix a race condition with Mutex
- [ ] Cancel a long operation with `context.WithTimeout`

**Self-test:** What happens if you send to a closed channel?
> *Panic. Only the sender should close.*

---

## Next Module

→ [04 Testing & Debugging](../04_testing_and_debugging/README.md)
