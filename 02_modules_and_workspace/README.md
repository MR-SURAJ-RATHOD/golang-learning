# 02. Modules & Workspace

> **Time:** 2–3 hours | **Prerequisites:** [Module 01](../01_go_fundamentals/README.md)

Learn how Go manages dependencies with modules and multi-module workspaces.

---

## What You Will Learn

- [ ] `go.mod` — module definition and versioning
- [ ] `go get` — adding external packages
- [ ] `go mod tidy` — cleaning dependencies
- [ ] `go.work` — working with multiple modules locally
- [ ] Local packages vs external packages

---

## Project Structure

```
02_modules_and_workspace/
├── main.go          # Uses external + local packages
├── go.mod           # This module's dependencies
├── greet/           # Local module (separate go.mod)
│   ├── greet.go
│   └── go.mod
└── exercises/
```

---

## Step-by-Step Lessons

### Step 1 — Understand go.mod

Open `go.mod` in the editor. Key lines:

```go
module github.com/suraj-iot-engineer/golang-learning/02_modules_and_workspace
go 1.22
require github.com/fatih/color v1.x.x  // external dependency
```

### Step 2 — Run the program

```bash
cd 02_modules_and_workspace
go mod tidy    # download missing dependencies
go run main.go
```

### Step 3 — Explore local module

The `greet/` folder is a **separate Go module** with its own `go.mod`:

```bash
cd greet
go run greet.go   # won't work — it's a library, not main
```

It's imported in `main.go` as:
```go
import "github.com/suraj-iot-engineer/golang-learning/02_modules_and_workspace/greet"
```

### Step 4 — Set up Go Workspace (from repo root)

```bash
cd ..   # back to golang-learning root
go work init
go work use ./02_modules_and_workspace ./02_modules_and_workspace/greet
go work sync
```

This lets you edit `greet/` and `main.go` together without publishing `greet` anywhere.

### Step 5 — Try adding a dependency

```bash
go get github.com/google/uuid
go mod tidy
```

Then remove the import if you don't use it and run `go mod tidy` again.

---

## Essential Commands

| Command | Purpose |
| :--- | :--- |
| `go mod init <path>` | Create new module |
| `go get pkg@version` | Add or update dependency |
| `go mod tidy` | Add missing, remove unused deps |
| `go mod download` | Download deps to cache |
| `go list -m all` | List all dependencies |
| `go work init` | Create workspace file |
| `go work use ./path` | Add module to workspace |

---

## Exercises

See [exercises/EXERCISES.md](./exercises/EXERCISES.md)

1. Create a new local package `mathutil` with `Add(a, b int) int`
2. Import it in `main.go` and print the result
3. Run `go mod tidy` and verify `go.sum` updates

---

## Checkpoint

Before Module 03, you should be able to:

- [ ] Create a module with `go mod init`
- [ ] Add an external package with `go get`
- [ ] Explain what `go mod tidy` does
- [ ] Set up `go.work` for two local modules

---

## Next Module

→ [03 Concurrency & Context](../03_concurrency_and_context/README.md)
