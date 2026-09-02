# Getting Started with Go

## Verify installation
```powershell
go version
go env
```

## Your first program
Create `main.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
Run it:
```powershell
go run main.go
```
Build:
```powershell
go build -o hello.exe main.go
.\hello.exe
```

## Running the examples in this workspace
```powershell
go run 1_hello_world/main.go
go run 2_simple_values/main.go
```

## Modules (quick start)
From the project root:
```powershell
go mod init github.com/yourname/project
go run ./...
```

## Quick tips
- Use `go fmt` or `gofmt -w .` to format code.
- `go doc` or `godoc` to read docs for packages.
- Use `go env GOPATH` to check paths.