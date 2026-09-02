# Progress Tracker

Mark each item as you complete it. Fork this repo and keep your checklist updated — recruiters can see your journey through commit history.

**How to use:** Change `[ ]` to `[x]` as you finish each item.

---

## Overall Progress

| Phase | Modules | Status |
| :--- | :---: | :---: |
| Phase 1 — Core Go | 01–03 | ⬜ Not started |
| Phase 2 — Backend | 04–07 | ⬜ Not started |
| Phase 3 — Architecture | 08–10 | ⬜ Not started |
| Phase 4 — Cloud & IoT | 11–13 | ⬜ Not started |
| Phase 5 — DevOps & Capstone | 14–15 | ⬜ Not started |

> Update status: ⬜ Not started → 🟡 In progress → ✅ Complete

---

## Phase 1: Core Go

### Module 01 — Go Fundamentals
- [ ] Read README
- [ ] Run `go run main.go`
- [ ] Understand zero values
- [ ] Understand structs and interfaces
- [ ] Understand error handling pattern
- [ ] **Checkpoint:** Explain Go's approach to OOP (composition)

### Module 02 — Modules & Workspace
- [ ] Read README
- [ ] Run `go run main.go`
- [ ] Create a new module with `go mod init`
- [ ] Add an external dependency with `go get`
- [ ] **Checkpoint:** Explain `go mod tidy` purpose

### Module 03 — Concurrency & Context
- [ ] Read README
- [ ] Run `go run main.go`
- [ ] Write a goroutine example
- [ ] Implement a worker pool
- [ ] Use `context.WithTimeout`
- [ ] **Checkpoint:** Build worker pool with 5 workers, 100 jobs

---

## Phase 2: Backend Engineering

### Module 04 — Testing & Debugging
- [ ] Read README
- [ ] Run `go test ./...`
- [ ] Write a table-driven test
- [ ] Run a benchmark (`go test -bench=.`)
- [ ] **Checkpoint:** Achieve tests for all public functions

### Module 05 — Databases
- [ ] Read README
- [ ] Run `go run main.go`
- [ ] Perform CRUD with `database/sql`
- [ ] Understand connection pooling
- [ ] **Checkpoint:** Explain SQL injection prevention in Go

### Module 06 — REST API & Gin
- [ ] Read README
- [ ] Run API server (`go run main.go`)
- [ ] Test endpoints with curl/Postman
- [ ] Add custom middleware
- [ ] **Checkpoint:** Build CRUD API for a resource

### Module 07 — Authentication & Security
- [ ] Read README
- [ ] Run `go run main.go`
- [ ] Generate and validate JWT
- [ ] Hash password with bcrypt
- [ ] **Checkpoint:** Protect an endpoint with auth middleware

---

## Phase 3: Architecture & Microservices

### Module 08 — Advanced Go Patterns
- [ ] Read README
- [ ] Run `go run main.go`
- [ ] Implement functional options pattern
- [ ] Write a generic function
- [ ] **Checkpoint:** Explain when NOT to use reflection

### Module 09 — Backend Architecture
- [ ] Read README
- [ ] Run `go run main.go`
- [ ] Read [ARCHITECTURE.md](./ARCHITECTURE.md)
- [ ] Draw Clean Architecture layers
- [ ] **Checkpoint:** Explain ports vs adapters

### Module 10 — Microservices
- [ ] Read README
- [ ] Run gRPC server + client
- [ ] Read [MICROSERVICES.md](./MICROSERVICES.md)
- [ ] Write a `.proto` file
- [ ] **Checkpoint:** Explain gRPC vs REST trade-offs

---

## Phase 4: Cloud Native & IoT

### Module 11 — Cloud Native Go
- [ ] Read README
- [ ] Run `go run main.go`
- [ ] Understand cloud SDK basics
- [ ] **Checkpoint:** Explain 12-factor app principles

### Module 12 — IoT Backends
- [ ] Read README
- [ ] Understand MQTT pub/sub
- [ ] **Checkpoint:** Design sensor ingestion pipeline

### Module 13 — Docker & Kubernetes
- [ ] Read README
- [ ] Build Docker image (`docker build`)
- [ ] Apply K8s manifests
- [ ] **Checkpoint:** Explain multi-stage Docker build benefits

---

## Phase 5: DevOps & Capstone

### Module 14 — CI/CD Pipeline
- [ ] Read README
- [ ] Run `go run main.go`
- [ ] Review `ci.yml` workflow
- [ ] **Checkpoint:** List what a Go CI pipeline should include

### Module 15 — Capstone Projects
- [ ] Read README
- [ ] Choose a capstone project
- [ ] Set up project structure
- [ ] Implement core services
- [ ] Add tests
- [ ] Add Docker + K8s manifests
- [ ] Add CI/CD pipeline
- [ ] Deploy (local K8s or cloud)
- [ ] **Checkpoint:** Demo working system end-to-end

---

## Capstone Project Checklist

### Option A: E-Commerce Microservices

- [ ] User Service (auth, profiles)
- [ ] Product/Catalog Service
- [ ] Order Service
- [ ] Payment Service (mock)
- [ ] gRPC communication between services
- [ ] PostgreSQL per service
- [ ] Redis caching
- [ ] API Gateway
- [ ] Docker images for all services
- [ ] Kubernetes deployment
- [ ] CI/CD pipeline (GitHub Actions)
- [ ] Structured logging
- [ ] Health check endpoints
- [ ] README with architecture diagram

### Option B: IoT Data Platform

- [ ] MQTT broker connection
- [ ] High-throughput ingestion (worker pools)
- [ ] TimescaleDB or PostgreSQL storage
- [ ] WebSocket real-time dashboard
- [ ] Alert system (threshold-based)
- [ ] Docker deployment
- [ ] Grafana dashboards
- [ ] CI/CD pipeline
- [ ] README with architecture diagram

---

## Skills Self-Assessment

After completing all modules, rate yourself using [SKILLS_MATRIX.md](./SKILLS_MATRIX.md):

| Skill Area | My Level (1-5) | Date Completed |
| :--- | :---: | :--- |
| Core Go | | |
| Backend Engineering | | |
| Architecture & Patterns | | |
| Microservices | | |
| Cloud & DevOps | | |
| Capstone Project | | |

---

## Share Your Progress

When ready to show recruiters:

1. **Fork this repo** and keep your checklist updated
2. **Add a `MY_PROGRESS.md`** in your fork with skills ratings
3. **Link your capstone** project repository
4. **Add to resume/LinkedIn:**
   ```
   Go Backend Learning — 15-module curriculum covering
   microservices, Clean Architecture, gRPC, Kubernetes
   GitHub: [your-fork-url]
   ```

---

## Related Docs

- [LEARNING_PATH.md](./LEARNING_PATH.md) — what to study in each module
- [SKILLS_MATRIX.md](./SKILLS_MATRIX.md) — skills per module
- [RECRUITER_GUIDE.md](./RECRUITER_GUIDE.md) — how recruiters evaluate this
