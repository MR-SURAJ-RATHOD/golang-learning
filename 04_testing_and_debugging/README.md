# 04. Testing & Debugging

> **Time:** 3–4 hours | **Prerequisites:** [Module 03](../03_concurrency_and_context/README.md)

Write reliable Go code with table-driven tests, benchmarks, fuzzing, and the race detector.

---

## What You Will Learn

- [ ] Unit tests with `testing` package
- [ ] Table-driven tests with `t.Run`
- [ ] Benchmarks (`go test -bench`)
- [ ] Fuzz testing (`go test -fuzz`)
- [ ] Race detector (`go test -race`)
- [ ] Debugging with Delve

---

## Step-by-Step Lessons

### Step 1 — Run tests

```bash
cd 04_testing_and_debugging
go test -v
```

Expected: all tests PASS.

### Step 2 — Table-driven tests

Open `main_test.go` → `TestDivide`. Notice the pattern:

```go
tests := []struct {
    name        string
    a, b        int
    want        int
    expectError bool
}{ ... }

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) { ... })
}
```

This pattern is **idiomatic Go** — use it everywhere.

### Step 3 — Run benchmarks

```bash
go test -bench=. -benchmem
```

Output shows ns/op (nanoseconds per operation).

### Step 4 — Run fuzz tests

```bash
go test -fuzz=FuzzDivide -fuzztime=10s
```

Fuzzing finds edge cases automatically.

### Step 5 — Race detector

```bash
go test -race .
```

Use this whenever you write concurrent code.

### Step 6 — Debugging with Delve (optional)

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
dlv debug main.go
# In dlv: break main.main, continue, next, print calc
```

---

## Test File Naming

| File | Purpose |
| :--- | :--- |
| `main.go` | Source code |
| `main_test.go` | Tests (same package) |
| `*_test.go` | Must end with `_test.go` |

Test functions must start with `Test`, benchmarks with `Benchmark`, fuzz with `Fuzz`.

---

## Exercises

See [exercises/EXERCISES.md](./exercises/EXERCISES.md)

---

## Checkpoint

Before Module 05:

- [ ] Write a table-driven test with 5+ cases
- [ ] Run a benchmark and read the output
- [ ] Explain when to use `t.Fatal` vs `t.Error`
- [ ] Run tests with `-race` flag

---

## Next Module

→ [05 Databases](../05_databases_sql_nosql/README.md)
