# Master Golang — Zero to Cloud-Native Backend Architect

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)
[![Modules](https://img.shields.io/badge/Modules-15-blue?style=for-the-badge)]()
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](http://makeapullrequest.com)

> **Professional Go learning platform — A to Z.**
> Learn backend development, microservices, cloud-native systems, and production architecture using idiomatic Go.
> Built so **learners can prove their skills** and **recruiters can verify them** through code, checklists, and capstone projects.

**👉 Start learning: [docs/START_HERE.md](docs/START_HERE.md) — follow modules 01 → 15 in order.**

---

## Who Is This For?

| Audience | What You Get |
| :--- | :--- |
| **Learners** | Step-by-step modules, runnable code, exercises, progress tracker |
| **Recruiters / Hiring Managers** | [Skills Matrix](docs/SKILLS_MATRIX.md), [Recruiter Guide](docs/RECRUITER_GUIDE.md), capstone proof-of-work |
| **Developers switching stacks** | Clean Architecture, gRPC, K8s — real patterns, not toy examples |

---

## Learning Journey (Basic → Advanced → Production)

```mermaid
flowchart TB
    subgraph Phase1["Phase 1 — Core Go"]
        M01[01 Fundamentals]
        M02[02 Modules & Workspace]
        M03[03 Concurrency & Context]
    end

    subgraph Phase2["Phase 2 — Backend Engineering"]
        M04[04 Testing & Debugging]
        M05[05 Databases]
        M06[06 REST API & Gin]
        M07[07 Auth & Security]
    end

    subgraph Phase3["Phase 3 — Architecture & Microservices"]
        M08[08 Advanced Patterns]
        M09[09 Backend Architecture]
        M10[10 Microservices & gRPC]
    end

    subgraph Phase4["Phase 4 — Cloud & IoT"]
        M11[11 Cloud Native Go]
        M12[12 IoT Backends]
        M13[13 Docker & Kubernetes]
    end

    subgraph Phase5["Phase 5 — DevOps & Capstone"]
        M14[14 CI/CD Pipeline]
        M15[15 Capstone Projects]
    end

    M01 --> M02 --> M03
    M03 --> M04 --> M05 --> M06 --> M07
    M07 --> M08 --> M09 --> M10
    M10 --> M11 --> M12 --> M13
    M13 --> M14 --> M15

    M15 --> CERT["Portfolio Ready — Show Recruiters"]
```

---

## Full System Architecture (What You Will Build)

```mermaid
flowchart LR
    subgraph Client["Clients"]
        WEB[Web / Mobile App]
        IOT[IoT Devices / MQTT]
    end

    subgraph Gateway["API Gateway"]
        GW[Ingress / API Gateway]
    end

    subgraph Services["Microservices — Go"]
        USR[User Service]
        PRD[Product Service]
        ORD[Order Service]
        PAY[Payment Service]
    end

    subgraph Data["Data Layer"]
        PG[(PostgreSQL)]
        RD[(Redis Cache)]
        MQ[Message Queue]
    end

    subgraph Observability["Observability"]
        LOG[Centralized Logging]
        TRACE[Distributed Tracing]
        MET[Metrics / Prometheus]
    end

    WEB --> GW
    IOT --> MQ
    GW --> USR
    GW --> PRD
    GW --> ORD
    GW --> PAY

    USR --> PG
    PRD --> PG
    ORD --> PG
    ORD --> RD
    ORD --> MQ
    PAY --> PG

    USR -.-> LOG
    PRD -.-> LOG
    ORD -.-> TRACE
    PAY -.-> MET
```

---

## Clean Architecture Layers (Module 09)

```mermaid
flowchart TB
    subgraph Adapters["Adapters — Infrastructure"]
        HTTP[HTTP / gRPC Handlers]
        DB[Database Repositories]
        EXT[External APIs]
    end

    subgraph Ports["Ports — Interfaces"]
        IN[Input Ports]
        OUT[Output Ports]
    end

    subgraph Core["Core — Domain"]
        ENT[Entities]
        UC[Use Cases / Services]
        DOM[Business Rules]
    end

    HTTP --> IN
    IN --> UC
    UC --> ENT
    UC --> DOM
    UC --> OUT
    OUT --> DB
    OUT --> EXT
```

---

## Course Curriculum (15 Modules)

Each module has **README + runnable `main.go` + learning objectives**. Track progress in [docs/PROGRESS.md](docs/PROGRESS.md).

### Phase 1: Core Systems & Concurrency

| # | Module | Topic | Key Skills |
| :---: | :--- | :--- | :--- |
| 01 | [Go Fundamentals](./01_go_fundamentals) | Syntax, Types, Errors | Variables, structs, interfaces, `if err != nil` |
| 02 | [Modules & Workspace](./02_modules_and_workspace) | Dependency Management | `go.mod`, `go.work`, versioning |
| 03 | [Concurrency & Context](./03_concurrency_and_context) | Goroutines & Channels | WaitGroups, mutexes, worker pools, `context` |

### Phase 2: Modern Backend Engineering

| # | Module | Topic | Key Skills |
| :---: | :--- | :--- | :--- |
| 04 | [Testing & Debugging](./04_testing_and_debugging) | Quality Assurance | Table-driven tests, benchmarks, fuzzing |
| 05 | [Databases](./05_databases_sql_nosql) | Data Persistence | `database/sql`, GORM, Redis, migrations |
| 06 | [REST API & Gin](./06_rest_api_and_gin) | API Development | Gin, middleware, validation, OpenAPI |
| 07 | [Auth & Security](./07_authentication_and_security) | Security | JWT, OAuth2, RBAC, password hashing |

### Phase 3: Advanced Architecture & Microservices

| # | Module | Topic | Key Skills |
| :---: | :--- | :--- | :--- |
| 08 | [Advanced Patterns](./08_advanced_go_patterns) | Expert Go | Generics, functional options, reflection |
| 09 | [Backend Architecture](./09_backend_architecture) | System Design | Clean Architecture, Hexagonal, DDD |
| 10 | [Microservices](./10_microservices) | Distributed Systems | gRPC, Protobuf, service discovery |

### Phase 4: Cloud Native & IoT

| # | Module | Topic | Key Skills |
| :---: | :--- | :--- | :--- |
| 11 | [Cloud Native Go](./11_cloud_native_go) | Cloud SDKs | AWS/GCP, Lambda, object storage |
| 12 | [IoT Backends](./12_iot_backend_projects) | IoT Integration | MQTT, sensor ingestion, pipelines |
| 13 | [Docker & Kubernetes](./13_docker_and_kubernetes) | Deployment | Multi-stage builds, K8s manifests, Helm |

### Phase 5: DevOps & Capstones

| # | Module | Topic | Key Skills |
| :---: | :--- | :--- | :--- |
| 14 | [CI/CD Pipeline](./14_ci_cd_pipeline) | Automation | Local scripts, linting, pre-push testing |
| 15 | [Capstone Projects](./15_capstone_projects) | Production Projects | Full microservices backend, IoT platform |

---

## Quick Start (5 Minutes)

### Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [Docker](https://www.docker.com/) (for modules 13+)
- VS Code + [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go)

### Install & Run

```bash
git clone https://github.com/your-username/golang-learning.git
cd golang-learning

# Run first module
cd 01_go_fundamentals
go run main.go

# Verify all modules build (Windows)
cd ..
.\build_and_test.ps1
```

### Recommended Learning Order

1. Read [docs/INSTALL.md](docs/INSTALL.md) → [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md)
2. Follow [docs/LEARNING_PATH.md](docs/LEARNING_PATH.md) module by module
3. Check off items in [docs/PROGRESS.md](docs/PROGRESS.md)
4. Complete [15_capstone_projects](./15_capstone_projects) for portfolio proof

---

## Prove Your Skills (For Learners & Recruiters)

```mermaid
flowchart LR
    LEARN[Complete Modules 01–14] --> CODE[Push Your Fork / PRs]
    CODE --> CAP[Build Capstone Project]
    CAP --> PROOF[Share Proof Package]

    subgraph Proof["What Recruiters Can Verify"]
        GH[GitHub Repo with Commits]
        TEST[Passing Tests & CI]
        ARCH[Architecture Docs]
        DEMO[Runnable Demo / Deploy]
    end

    PROOF --> GH
    PROOF --> TEST
    PROOF --> ARCH
    PROOF --> DEMO
```

| Document | Purpose |
| :--- | :--- |
| [docs/PROGRESS.md](docs/PROGRESS.md) | Self-assessment checklist — mark modules done |
| [docs/SKILLS_MATRIX.md](docs/SKILLS_MATRIX.md) | Skill-to-module mapping for interviews |
| [docs/RECRUITER_GUIDE.md](docs/RECRUITER_GUIDE.md) | How to evaluate a candidate using this repo |
| [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) | Full system design with diagrams |
| [docs/MICROSERVICES.md](docs/MICROSERVICES.md) | Microservices patterns A to Z |

---

## Repository Structure

```
golang-learning/
├── 01_go_fundamentals/          # Phase 1 — Core Go
├── 02_modules_and_workspace/
├── 03_concurrency_and_context/
├── 04_testing_and_debugging/    # Phase 2 — Backend
├── 05_databases_sql_nosql/
├── 06_rest_api_and_gin/
├── 07_authentication_and_security/
├── 08_advanced_go_patterns/     # Phase 3 — Architecture
├── 09_backend_architecture/
├── 10_microservices/
├── 11_cloud_native_go/          # Phase 4 — Cloud & IoT
├── 12_iot_backend_projects/
├── 13_docker_and_kubernetes/
├── 14_ci_cd_pipeline/           # Phase 5 — DevOps
├── 15_capstone_projects/
├── docs/                        # Guides, cheat sheets, progress
├── legacy/                      # Original beginner examples
└── build_and_test.ps1           # Verify all modules
```

See [docs/PROJECT_STRUCTURE.md](docs/PROJECT_STRUCTURE.md) for package layout best practices.

---

## Documentation Index

| Guide | Description |
| :--- | :--- |
| [docs/LEARNING_PATH.md](docs/LEARNING_PATH.md) | Complete A→Z learning roadmap |
| [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) | System design & architecture patterns |
| [docs/MICROSERVICES.md](docs/MICROSERVICES.md) | Microservices communication & patterns |
| [docs/PROGRESS.md](docs/PROGRESS.md) | Module completion tracker |
| [docs/SKILLS_MATRIX.md](docs/SKILLS_MATRIX.md) | Skills map for interviews |
| [docs/RECRUITER_GUIDE.md](docs/RECRUITER_GUIDE.md) | Hiring manager evaluation guide |
| [docs/CHEAT_SHEET.md](docs/CHEAT_SHEET.md) | Quick command reference |
| [docs/RESOURCES.md](docs/RESOURCES.md) | External learning links |

---

## Capstone Projects (Portfolio)

After completing modules 01–14, build one of these to demonstrate production readiness:

| Project | Stack | Proves |
| :--- | :--- | :--- |
| **E-Commerce Microservices** | Go, gRPC, PostgreSQL, Redis, K8s | Distributed systems, transactions, deployment |
| **IoT Data Platform** | Go, MQTT, TimescaleDB, Grafana | High-throughput ingestion, real-time dashboards |

Details: [15_capstone_projects/README.md](./15_capstone_projects/README.md)

---

## Contributing

Contributions welcome! Fork → branch → PR.

1. Fork the project
2. Create feature branch (`git checkout -b feature/new-module`)
3. Commit changes (`git commit -m 'Add new module'`)
4. Push and open a Pull Request

---

## License

MIT License — see [LICENSE](LICENSE).

---

<p align="center">
  <strong>Learn Go. Build Systems. Prove It.</strong><br>
  Start here → <a href="docs/LEARNING_PATH.md">Learning Path</a> · <a href="docs/PROGRESS.md">Track Progress</a>
</p>
