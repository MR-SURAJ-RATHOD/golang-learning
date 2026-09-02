# Project Structure & Best Practices

## Recommended layout
- Use modules (`go mod init`) at the repo root.
- Organize code into packages, e.g.:
  - `cmd/` for commands
  - `pkg/` for reusable packages
  - `internal/` for non-public packages

## Example: this workspace
- Each folder like `1_hello_world` can be a small `package main` example.
- To run all: `go run ./...` from the module root.

## GOPATH vs Modules
- Use Go modules for dependency/version management (since Go 1.13+).
- GOPATH still exists for historical code, but new projects should use modules.