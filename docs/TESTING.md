# Testing in Go

## Writing tests
Create `foo_test.go` in the same package:
```go
package foo

import "testing"

func TestAdd(t *testing.T) {
    if Add(1, 2) != 3 {
        t.Fatalf("expected 3")
    }
}
```

## Table-driven tests
```go
func TestSum(t *testing.T) {
    cases := []struct{a,b,exp int}{ {1,2,3}, {2,3,5} }
    for _, c := range cases {
        if got := Sum(c.a, c.b); got != c.exp {
            t.Fatalf("%v", got)
        }
    }
}
```

## Benchmarks
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1,2)
    }
}
```

Run tests:
```powershell
go test ./...
go test -v ./pkg -run TestName
go test -bench=.
```

Use coverage:
```powershell
go test ./... -cover
```