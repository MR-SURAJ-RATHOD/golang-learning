# Go Basics

## Packages & imports
- Every Go file starts with a `package` declaration.
- `package main` with `func main()` produces an executable.

## Variables & types
```go
var x int = 5
y := 10 // short declaration
const Pi = 3.14
```

## Control flow
- `if`, `for`, `switch` (no `while` — `for` is used instead)

Example:
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## Functions
```go
func add(a, b int) int { return a + b }
```

## Common built-ins
- `make`, `new`, slices, maps, structs

For more, read samples in `2_simple_values` and `3_variables`.