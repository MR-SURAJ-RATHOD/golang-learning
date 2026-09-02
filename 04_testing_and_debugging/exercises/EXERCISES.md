# Module 04 — Exercises

## Exercise 1: Test Multiply

`Multiply` already has tests. Add a case for `Multiply(-1, -1)` expecting `1`.

## Exercise 2: New Function + Tests

Add `func (c Calculator) Subtract(a, b int) int` to `main.go`.
Write table-driven tests with at least 4 cases.

## Exercise 3: Benchmark Compare

Add `BenchmarkMultiply` and compare ns/op with `BenchmarkAdd`.

## Exercise 4: Table-Driven Error Test

Add `func (c Calculator) Sqrt(n int) (int, error)` that errors on negative `n`.
Write tests for valid and invalid inputs.

## Exercise 5: Coverage

```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```
Aim for >80% coverage on Calculator methods.
