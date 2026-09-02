# Backend Architecture Guide

Industry-standard patterns used in this learning platform.
Module **09** has runnable code; this document explains the **full picture**.

---

## Architecture Evolution

```mermaid
flowchart LR
    A[Monolith<br/>Single binary] --> B[Modular Monolith<br/>Clean layers]
    B --> C[Microservices<br/>Independent deploy]
    C --> D[Cloud Native<br/>K8s + Observability]

    style A fill:#f9f,stroke:#333
    style D fill:#9f9,stroke:#333
```

| Stage | When to Use | Go Advantage |
| :--- | :--- | :--- |
| Monolith | MVP, small team | Fast compile, single binary |
| Modular Monolith | Growing codebase | Package boundaries, easy refactor |
| Microservices | Scale teams independently | Small images, gRPC performance |
| Cloud Native | Production at scale | Excellent K8s support |

---

## Clean Architecture

Separates business logic from frameworks, databases, and UI.

```mermaid
flowchart TB
    subgraph Outer["Frameworks & Drivers"]
        WEB[Web / gRPC]
        DB[(Database)]
        MQ[Message Queue]
    end

    subgraph Interface["Interface Adapters"]
        CTRL[Controllers / Handlers]
        REPO[Repository Impl]
        GW[External Gateway]
    end

    subgraph App["Application Business Rules"]
        UC[Use Cases]
    end

    subgraph Enterprise["Enterprise Business Rules"]
        ENT[Entities / Domain Models]
    end

    WEB --> CTRL
    CTRL --> UC
    UC --> ENT
    UC --> REPO
    REPO --> DB
    GW --> UC
    MQ --> GW
```

### Dependency Rule

**Dependencies point inward only.** Inner layers never import outer layers.

```
Entities  ←  Use Cases  ←  Interface Adapters  ←  Frameworks
(core)        (app logic)    (handlers, repos)      (gin, postgres)
```

### Go Project Layout

```
my-service/
├── cmd/
│   └── api/
│       └── main.go          # Entry point, wiring
├── internal/
│   ├── domain/              # Entities, business rules
│   │   └── order.go
│   ├── usecase/             # Application logic
│   │   └── order_service.go
│   ├── adapter/
│   │   ├── handler/         # HTTP/gRPC handlers (primary adapter)
│   │   │   └── order_handler.go
│   │   └── repository/      # DB implementations (secondary adapter)
│   │       └── postgres_order_repo.go
│   └── port/                # Interfaces (contracts)
│       └── order_repository.go
├── pkg/                     # Shared, importable packages
└── go.mod
```

---

## Hexagonal Architecture (Ports & Adapters)

Same idea as Clean Architecture, different naming.

```mermaid
flowchart TB
    subgraph External["External World"]
        CLI[CLI]
        HTTP[HTTP API]
        GRPC[gRPC]
        PG[(PostgreSQL)]
        REDIS[(Redis)]
    end

    subgraph Hexagon["Application Core"]
        subgraph Ports["Ports — Interfaces"]
            IN_PORT[Input Port<br/>OrderService]
            OUT_PORT[Output Port<br/>OrderRepository]
        end
        CORE[Domain Logic<br/>CreateOrder, Validate]
    end

    CLI --> IN_PORT
    HTTP --> IN_PORT
    GRPC --> IN_PORT
    IN_PORT --> CORE
    CORE --> OUT_PORT
    OUT_PORT --> PG
    OUT_PORT --> REDIS
```

**Runnable example:** `09_backend_architecture/main.go`

---

## Domain-Driven Design (DDD) — Essentials

```mermaid
classDiagram
    class Order {
        +ID string
        +Items []OrderItem
        +Total() float64
        +Place() error
    }
    class OrderItem {
        +ProductID string
        +Quantity int
        +Price float64
    }
    class OrderService {
        +CreateOrder(cmd) Order
        +CancelOrder(id) error
    }
    class OrderRepository {
        <<interface>>
        +Save(order) error
        +FindByID(id) Order
    }

    Order "1" --> "*" OrderItem
    OrderService --> Order
    OrderService --> OrderRepository
```

### DDD Building Blocks

| Concept | Description | Go Example |
| :--- | :--- | :--- |
| **Entity** | Object with identity | `type Order struct { ID string }` |
| **Value Object** | Immutable, no identity | `type Money struct { Amount int; Currency string }` |
| **Aggregate** | Cluster of entities | Order + OrderItems |
| **Repository** | Persistence abstraction | `OrderRepository` interface |
| **Service** | Domain logic across entities | `OrderService` |

---

## Microservices Architecture (Capstone)

Full e-commerce backend from Module 15.

```mermaid
flowchart TB
    subgraph Clients
        APP[Mobile / Web]
    end

    subgraph Edge
        ING[Ingress / Load Balancer]
        GW[API Gateway]
    end

    subgraph ServiceMesh["Microservices"]
        USR[User Service<br/>Auth, Profiles]
        CAT[Catalog Service<br/>Products]
        ORD[Order Service<br/>Orders, Cart]
        PAY[Payment Service<br/>Transactions]
        NOT[Notification Service<br/>Email, SMS]
    end

    subgraph Messaging
        KAFKA[Event Bus<br/>Kafka / NATS]
    end

    subgraph DataStores
        PG1[(Users DB)]
        PG2[(Catalog DB)]
        PG3[(Orders DB)]
        CACHE[(Redis)]
    end

    subgraph Observability
        JAEGER[Jaeger Tracing]
        PROM[Prometheus]
        GRAF[Grafana]
        ELK[ELK Logging]
    end

    APP --> ING --> GW
    GW -->|REST| USR
    GW -->|REST| CAT
    GW -->|REST| ORD

    ORD -->|gRPC| PAY
    ORD -->|gRPC| CAT
    ORD --> CACHE
    ORD --> PG3
    ORD -->|event| KAFKA

    KAFKA --> NOT
    USR --> PG1
    CAT --> PG2

    USR -.-> JAEGER
    ORD -.-> PROM
    PAY -.-> ELK
```

### Communication Patterns

| Pattern | Use Case | Protocol |
| :--- | :--- | :--- |
| **Sync Request-Response** | Get user profile, place order | REST / gRPC |
| **Async Event-Driven** | Order placed → send email | Kafka, NATS, RabbitMQ |
| **CQRS** | Separate read/write models | Event sourcing + projections |
| **Saga** | Distributed transactions | Choreography or orchestration |

---

## API Gateway Pattern

```mermaid
sequenceDiagram
    participant Client
    participant Gateway
    participant UserSvc
    participant OrderSvc
    participant PaymentSvc

    Client->>Gateway: POST /orders
    Gateway->>Gateway: Validate JWT
    Gateway->>UserSvc: gRPC GetUser(id)
    UserSvc-->>Gateway: User{verified}
    Gateway->>OrderSvc: gRPC CreateOrder(req)
    OrderSvc->>PaymentSvc: gRPC ProcessPayment(amount)
    PaymentSvc-->>OrderSvc: Payment{success}
    OrderSvc-->>Gateway: Order{id, status}
    Gateway-->>Client: 201 Created
```

---

## Data Management

### Database per Service

Each microservice owns its database. No shared tables.

```mermaid
flowchart LR
    USR[User Service] --> UDB[(users_db)]
    CAT[Catalog Service] --> CDB[(catalog_db)]
    ORD[Order Service] --> ODB[(orders_db)]
```

### Caching Strategy

```
Read:  Client → Redis (hit?) → PostgreSQL (miss) → populate Redis
Write: Client → PostgreSQL → invalidate Redis key
```

---

## Observability Stack

```mermaid
flowchart TB
    subgraph Services
        S1[Service A]
        S2[Service B]
        S3[Service C]
    end

    subgraph Logs
        S1 -->|structured JSON| ELK[ELK / Loki]
    end

    subgraph Traces
        S1 -->|OpenTelemetry| JAEGER[Jaeger]
        S2 --> JAEGER
        S3 --> JAEGER
    end

    subgraph Metrics
        S1 -->|/metrics| PROM[Prometheus]
        S2 --> PROM
        PROM --> GRAF[Grafana Dashboards]
    end
```

### Go Libraries

| Concern | Library |
| :--- | :--- |
| Structured logging | `log/slog`, `zerolog`, `zap` |
| Tracing | `go.opentelemetry.io/otel` |
| Metrics | `prometheus/client_golang` |

---

## Security Architecture

```mermaid
flowchart LR
    REQ[Incoming Request] --> MW[Auth Middleware]
    MW --> JWT{Valid JWT?}
    JWT -->|No| REJ[401 Unauthorized]
    JWT -->|Yes| RBAC{Has Role?}
    RBAC -->|No| FORB[403 Forbidden]
    RBAC -->|Yes| HND[Handler]
```

See Module 07: `07_authentication_and_security/`

---

## Deployment Architecture (Kubernetes)

```mermaid
flowchart TB
    subgraph K8s Cluster
        ING[Ingress Controller]
        
        subgraph NS["Namespace: production"]
            DEP1[Deployment: user-svc<br/>replicas: 3]
            DEP2[Deployment: order-svc<br/>replicas: 3]
            SVC1[Service: user-svc]
            SVC2[Service: order-svc]
            HPA[HPA: Auto-scaling]
        end

        CM[ConfigMap]
        SEC[Secrets]
    end

    ING --> SVC1
    ING --> SVC2
    SVC1 --> DEP1
    SVC2 --> DEP2
    DEP2 --> HPA
    DEP1 --> CM
    DEP1 --> SEC
```

See Module 13: `13_docker_and_kubernetes/`

---

## Architecture Decision Checklist

Before building, answer these:

- [ ] Monolith or microservices? (Start monolith unless you have a clear split)
- [ ] Sync (gRPC/REST) or async (events) for each interaction?
- [ ] Which database per service?
- [ ] How will you handle distributed failures? (retries, circuit breakers)
- [ ] How will you trace requests across services?
- [ ] How will you deploy? (Docker → K8s → CI/CD)

---

## Related Docs

- [MICROSERVICES.md](./MICROSERVICES.md) — gRPC, service mesh, patterns
- [LEARNING_PATH.md](./LEARNING_PATH.md) — module order
- [Module 09 code](../09_backend_architecture/main.go) — hexagonal example
