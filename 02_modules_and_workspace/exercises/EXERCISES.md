# Module 02 — Exercises

## Exercise 1: Create a Local Package

1. Create folder `mathutil/` with its own `go.mod`
2. Add `func Add(a, b int) int` in `mathutil/math.go`
3. Import in `main.go` and print `mathutil.Add(3, 4)`

## Exercise 2: go mod tidy

1. Run `go get github.com/google/uuid`
2. Add a line in main that prints `uuid.New()`
3. Run `go mod tidy` — verify `go.sum` has uuid
4. Remove uuid usage, run `go mod tidy` again — uuid should be removed

## Exercise 3: Workspace

From repo root:
```bash
go work init
go work use ./02_modules_and_workspace ./02_modules_and_workspace/greet
```
Edit `greet/greet.go` message and re-run `main.go` without publishing anything.
