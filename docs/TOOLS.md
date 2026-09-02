# Tools & Editor Setup

## VS Code
- Install the **Go** extension (`ms-vscode.Go`).
- Install `gopls` (language server):
```powershell
go install golang.org/x/tools/gopls@latest
```

## Formatting & linting
- `go fmt ./...` or `gofmt -w .`
- `go vet ./...`
- Install `golangci-lint`:
```powershell
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Other useful tools
- `delve` for debugging (install `github.com/go-delve/delve/cmd/dlv`)
- `go test`, `go build`, `go doc` for documentation

Configure your editor settings to run `gofmt`/`goimports` on save for consistent style.