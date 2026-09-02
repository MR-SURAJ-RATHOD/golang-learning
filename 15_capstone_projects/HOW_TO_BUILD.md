# Module 15 — Step-by-Step Capstone Guide

This is your **portfolio project**. Follow these steps in order.

---

## Before You Start

- [ ] Completed Modules 01–14
- [ ] Checked off [docs/PROGRESS.md](../docs/PROGRESS.md)
- [ ] Chosen Project 1 (E-Commerce) or Project 2 (IoT)

---

## Week 1: Foundation

### Day 1–2: Project Setup
```bash
mkdir my-capstone && cd my-capstone
go mod init github.com/YOUR_USERNAME/my-capstone
mkdir -p cmd/api internal/domain internal/usecase internal/adapter
```

### Day 3–4: Domain Layer
- Define entities (User, Order, Product, or SensorReading)
- Write business validation rules
- No framework imports in `internal/domain/`

### Day 5–7: First Service
- One complete service with Clean Architecture
- Repository interface + in-memory implementation
- HTTP or gRPC handler
- Tests with table-driven pattern

---

## Week 2: Expand

### Day 8–10: Second & Third Services
- Add communication (gRPC between services)
- Each service has own `go.mod` or packages

### Day 11–12: Database
- PostgreSQL or SQLite per service
- Migrations

### Day 13–14: Docker
- Dockerfile per service
- `docker-compose.yml` for local dev

---

## Week 3: Production Ready

### Day 15–17: CI/CD
- GitHub Actions: test + lint + build on push
- Green badge on README

### Day 18–19: Observability
- Structured logging
- Health endpoints

### Day 20–21: Documentation
- README with architecture mermaid diagram
- How to run locally
- What you learned

---

## Submission Checklist

Share with recruiters:

- [ ] GitHub repo link (public)
- [ ] README with architecture diagram
- [ ] `docker compose up` works
- [ ] CI badge is green
- [ ] 3+ commits showing progression
- [ ] LinkedIn post linking the project

---

## Evaluation Rubric

| Criteria | Points |
| :--- | :---: |
| Clean Architecture | 25 |
| Multiple services communicating | 25 |
| Tests + CI | 20 |
| Docker deployment | 15 |
| Documentation | 15 |

---

## Starter Code

IoT skeleton: `impl/iot-platform/main.go` — WebSocket echo server.
Extend it or start fresh in your own repo.
