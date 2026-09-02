# Go Skills Matrix

Map of **every skill** taught in this repo → **which module proves it**.
Use this for self-assessment, interviews, and recruiter evaluation.

---

## How to Use

| Who | How |
| :--- | :--- |
| **Learner** | Rate yourself 1–5 per skill after completing the module |
| **Recruiter** | Ask candidate to demo the module code + explain concepts |
| **Interviewer** | Pick 3–5 skills from the role's required level and drill down |

### Rating Scale

| Level | Meaning |
| :---: | :--- |
| 1 | Heard of it, can't implement |
| 2 | Can implement with docs/examples |
| 3 | Can implement independently |
| 4 | Can design systems using this skill |
| 5 | Can teach others, knows trade-offs |

---

## Core Go Skills

| Skill | Module | Proof |
| :--- | :---: | :--- |
| Variables, types, zero values | 01 | Run `01_go_fundamentals/main.go` |
| Control flow (if/switch/for) | 01 | Explain why Go has no `while` |
| Functions, closures, defer | 01 | Write a closure example |
| Pointers | 01 | Explain when to use `*` vs `&` |
| Structs & composition | 01 | Model a `User` struct |
| Interfaces (implicit) | 01 | Implement `io.Reader` |
| Error handling | 01 | Handle errors without panic |
| `go mod` dependency management | 02 | Create module, add dependency |
| Multi-module workspaces | 02 | Set up `go.work` |
| Goroutines | 03 | Launch 10 goroutines safely |
| Channels | 03 | Producer-consumer pattern |
| `sync.WaitGroup` / `Mutex` | 03 | Protect shared state |
| `context` cancellation | 03 | Timeout a slow operation |
| Worker pools | 03 | Process N jobs with M workers |

---

## Backend Engineering Skills

| Skill | Module | Proof |
| :--- | :---: | :--- |
| Table-driven tests | 04 | Write tests for 3+ cases |
| Benchmarks | 04 | `go test -bench=.` |
| Fuzzing | 04 | Run `go test -fuzz` |
| `database/sql` | 05 | CRUD with prepared statements |
| GORM / ORM | 05 | Model + migrate a table |
| Redis caching | 05 | Cache-aside pattern |
| REST API design | 06 | CRUD with proper status codes |
| Gin framework | 06 | Router + middleware chain |
| Request validation | 06 | Validate JSON input |
| JWT authentication | 07 | Generate + verify tokens |
| Password hashing (bcrypt) | 07 | Hash and compare passwords |
| RBAC | 07 | Role-based route protection |
| OAuth2 flow | 07 | Explain authorization code flow |

---

## Architecture & Advanced Skills

| Skill | Module | Proof |
| :--- | :---: | :--- |
| Generics | 08 | Write a generic `Stack[T]` |
| Functional options | 08 | Configure a service with options |
| Middleware pattern | 08 | Chain 3 middleware functions |
| Clean Architecture layers | 09 | Draw dependency direction |
| Hexagonal (Ports & Adapters) | 09 | Run `09_backend_architecture` |
| Dependency Injection | 09 | Wire interfaces to implementations |
| DDD entities & aggregates | 09 | Model an Order aggregate |
| gRPC server | 10 | Run `10_microservices/server` |
| gRPC client | 10 | Run `10_microservices/client` |
| Protocol Buffers | 10 | Write a `.proto` file |
| Circuit breaker | 10 | Explain open/closed/half-open |
| Service discovery | 10 | Explain K8s DNS vs Consul |

---

## Cloud & DevOps Skills

| Skill | Module | Proof |
| :--- | :---: | :--- |
| AWS/GCP SDK | 11 | Upload file to object storage |
| Cloud Functions | 11 | Deploy a Go Lambda |
| MQTT protocol | 12 | Subscribe/publish messages |
| Sensor data ingestion | 12 | Worker pool for MQTT messages |
| Multi-stage Dockerfile | 13 | Build image < 20MB |
| Docker Compose | 13 | Multi-service local setup |
| K8s Deployments | 13 | Apply `deployment.yaml` |
| K8s Services & Ingress | 13 | Expose service externally |
| GitHub Actions CI | 14 | Run tests on push |
| Linting (golangci-lint) | 14 | Zero lint errors |
| Release automation | 14 | Tag + release workflow |

---

## Capstone Skills (Portfolio Level)

| Skill | Project | Proof |
| :--- | :--- | :--- |
| Multi-service architecture | E-Commerce | 3+ services communicating via gRPC |
| Event-driven design | E-Commerce | OrderPlaced → email + inventory |
| High-throughput ingestion | IoT Platform | 10K+ messages/sec with worker pools |
| Real-time dashboard | IoT Platform | WebSocket data streaming |
| Full CI/CD pipeline | Either | Green CI badge on GitHub |
| Production deployment | Either | Running on K8s or cloud |

---

## Role-Based Skill Requirements

### Junior Go Backend Developer

**Must have (Level 3+):**
- Core Go (modules 01–03)
- Testing (04)
- REST API (06)
- Basic SQL (05)

**Nice to have:**
- Docker basics (13)
- JWT auth (07)

---

### Mid-Level Go Backend Developer

**Must have (Level 3+):**
- All Junior skills
- Concurrency patterns (03)
- Database design (05)
- Auth & security (07)
- Clean Architecture (09)

**Nice to have:**
- gRPC (10)
- CI/CD (14)
- K8s basics (13)

---

### Senior Go Backend / Architect

**Must have (Level 4+):**
- All Mid-level skills
- Microservices design (10)
- System design (09, ARCHITECTURE.md)
- Observability (logging, tracing, metrics)
- Cloud native patterns (11, 13)

**Nice to have:**
- IoT / event-driven (12)
- Capstone project completed (15)
- Can present architecture with trade-off analysis

---

## Interview Question Bank (Per Module)

| Module | Sample Question |
| :---: | :--- |
| 01 | What is the zero value of a slice? How is it different from `nil`? |
| 03 | Explain the difference between buffered and unbuffered channels. |
| 05 | How do you prevent SQL injection in Go? |
| 06 | What HTTP status code for validation failure vs server error? |
| 07 | How does JWT work? Where should you store the secret? |
| 09 | Explain the dependency rule in Clean Architecture. |
| 10 | When would you choose gRPC over REST? |
| 13 | Why use multi-stage Docker builds for Go? |
| 14 | What should a CI pipeline for Go include? |

---

## Self-Assessment Template

Copy this into your fork's README or a personal `MY_PROGRESS.md`:

```markdown
## My Go Skills (Date: YYYY-MM-DD)

| Skill Area | Level (1-5) | Module Completed | Notes |
| :--- | :---: | :---: | :--- |
| Core Go | /5 | ☐ 01 ☐ 02 ☐ 03 | |
| Backend Engineering | /5 | ☐ 04-07 | |
| Architecture | /5 | ☐ 08-10 | |
| Cloud & DevOps | /5 | ☐ 11-14 | |
| Capstone | /5 | ☐ 15 | |

**Portfolio:** [link to your capstone repo]
**LinkedIn:** [your profile]
```

---

## Related Docs

- [PROGRESS.md](./PROGRESS.md) — module completion checklist
- [RECRUITER_GUIDE.md](./RECRUITER_GUIDE.md) — how to evaluate candidates
- [LEARNING_PATH.md](./LEARNING_PATH.md) — step-by-step roadmap
