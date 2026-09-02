# Microservices with Go — Complete Guide

Everything you need to design, build, and deploy microservices in Go.
Module **10** has runnable gRPC code; this guide covers the full A→Z picture.

---

## What Are Microservices?

```mermaid
flowchart LR
    subgraph Monolith["Monolith (Before)"]
        M1[User Logic]
        M2[Product Logic]
        M3[Order Logic]
        M4[(Single DB)]
        M1 --- M2 --- M3 --- M4
    end

    subgraph Micro["Microservices (After)"]
        S1[User Service]
        S2[Product Service]
        S3[Order Service]
        D1[(Users DB)]
        D2[(Products DB)]
        D3[(Orders DB)]
        S1 --> D1
        S2 --> D2
        S3 --> D3
    end

    Monolith -->|Decompose| Micro
```

| Aspect | Monolith | Microservices |
| :--- | :--- | :--- |
| Deployment | All-or-nothing | Independent per service |
| Scaling | Scale entire app | Scale hot services only |
| Technology | One stack | Polyglot possible |
| Complexity | Lower initially | Higher — needs DevOps maturity |
| Go fit | Excellent for start | Excellent — small binaries, fast startup |

---

## When to Use Microservices

**Use when:**
- Team is large enough to own separate services (2-pizza rule)
- Different parts scale differently (orders spike, catalog doesn't)
- You need independent deployment cycles

**Don't use when:**
- Early startup / MVP stage
- Team < 5 engineers
- You haven't mastered a modular monolith first

---

## Communication Patterns

### 1. Synchronous — gRPC (Recommended for internal)

```mermaid
sequenceDiagram
    participant Client
    participant OrderSvc as Order Service
    participant PaySvc as Payment Service

    Client->>OrderSvc: gRPC CreateOrder()
    OrderSvc->>PaySvc: gRPC ProcessPayment()
    PaySvc-->>OrderSvc: PaymentResponse
    OrderSvc-->>Client: OrderResponse
```

**Why gRPC over REST internally:**
- Binary Protocol Buffers — smaller, faster
- Strong typing via `.proto` contracts
- Built-in streaming (unary, server, client, bidirectional)
- Native Go support

### 2. Synchronous — REST (Recommended for external/public APIs)

```mermaid
sequenceDiagram
    participant Mobile as Mobile App
    participant GW as API Gateway
    participant Svc as User Service

    Mobile->>GW: GET /api/v1/users/me
    GW->>Svc: Forward + JWT
    Svc-->>GW: JSON response
    GW-->>Mobile: JSON response
```

### 3. Asynchronous — Event-Driven

```mermaid
sequenceDiagram
    participant OrderSvc as Order Service
    participant Bus as Event Bus
    participant EmailSvc as Email Service
    participant InvSvc as Inventory Service

    OrderSvc->>Bus: Publish OrderPlaced
    Bus->>EmailSvc: OrderPlaced event
    Bus->>InvSvc: OrderPlaced event
    EmailSvc->>EmailSvc: Send confirmation
    InvSvc->>InvSvc: Reserve stock
```

---

## Protocol Buffers & gRPC Setup

### Step 1: Define `.proto` Contract

```protobuf
syntax = "proto3";

package order;

option go_package = "github.com/yourname/orderservice/pb";

service OrderService {
  rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse);
  rpc GetOrder(GetOrderRequest) returns (GetOrderResponse);
}

message CreateOrderRequest {
  string user_id = 1;
  repeated OrderItem items = 2;
}

message OrderItem {
  string product_id = 1;
  int32 quantity = 2;
}

message CreateOrderResponse {
  string order_id = 1;
  string status = 2;
}
```

### Step 2: Generate Go Code

```bash
# Install tools (once)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate
protoc --go_out=. --go-grpc_out=. proto/order.proto
```

### Step 3: Implement Server

```go
type orderServer struct {
    pb.UnimplementedOrderServiceServer
    repo OrderRepository
}

func (s *orderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    order, err := s.repo.Save(req)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "save failed: %v", err)
    }
    return &pb.CreateOrderResponse{OrderId: order.ID, Status: "created"}, nil
}
```

### Step 4: Implement Client

```go
conn, _ := grpc.Dial("order-service:50051", grpc.WithInsecure())
client := pb.NewOrderServiceClient(conn)
resp, err := client.CreateOrder(ctx, &pb.CreateOrderRequest{UserId: "u-123"})
```

**Runnable demo:** `10_microservices/server/main.go` + `client/main.go`

---

## Service Discovery

```mermaid
flowchart TB
    S1[Order Service] -->|register| REG[Service Registry<br/>Consul / etcd]
    S2[Payment Service] -->|register| REG
    S3[User Service] -->|register| REG

    CLIENT[API Gateway] -->|lookup| REG
    REG -->|resolve| S1
    REG -->|resolve| S2
```

| Tool | Use Case |
| :--- | :--- |
| **Kubernetes DNS** | Built-in when running on K8s (`user-svc.default.svc.cluster.local`) |
| **Consul** | Multi-cloud, health checks, KV store |
| **etcd** | K8s uses this internally |

On Kubernetes, you usually don't need external service discovery — K8s Services handle it.

---

## Resilience Patterns

### Circuit Breaker

```mermaid
stateDiagram-v2
    [*] --> Closed
    Closed --> Open: Failure threshold reached
    Open --> HalfOpen: Timeout expires
    HalfOpen --> Closed: Probe succeeds
    HalfOpen --> Open: Probe fails
```

```go
// Using github.com/sony/gobreaker
cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
    Name:        "payment-service",
    MaxRequests: 3,
    Interval:    10 * time.Second,
    Timeout:     30 * time.Second,
})

result, err := cb.Execute(func() (interface{}, error) {
    return paymentClient.ProcessPayment(ctx, req)
})
```

### Retry with Exponential Backoff

```go
for attempt := 0; attempt < maxRetries; attempt++ {
    resp, err := client.Call(ctx, req)
    if err == nil {
        return resp, nil
    }
    time.Sleep(time.Duration(math.Pow(2, float64(attempt))) * 100 * time.Millisecond)
}
```

### Timeout via Context

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
resp, err := client.GetUser(ctx, &pb.GetUserRequest{Id: userID})
```

---

## API Gateway

Single entry point for all clients.

```mermaid
flowchart LR
    MOB[Mobile] --> GW
    WEB[Web] --> GW
    GW[API Gateway<br/>Kong / Traefik / custom Go]
    GW --> USR[User Svc]
    GW --> ORD[Order Svc]
    GW --> CAT[Catalog Svc]
```

**Gateway responsibilities:**
- Authentication / JWT validation
- Rate limiting
- Request routing
- Response aggregation (BFF pattern)
- SSL termination

---

## Data Patterns

### Database per Service

```
user-service     → users_db (PostgreSQL)
catalog-service  → catalog_db (PostgreSQL)
order-service    → orders_db (PostgreSQL)
order-service    → cache (Redis)
```

### Saga Pattern (Distributed Transactions)

```mermaid
sequenceDiagram
    participant OS as Order Service
    participant PS as Payment Service
    participant IS as Inventory Service

    OS->>PS: Charge payment
    PS-->>OS: OK
    OS->>IS: Reserve stock
    IS-->>OS: FAIL (out of stock)
    OS->>PS: Refund payment (compensate)
```

Two approaches:
- **Choreography:** Services react to events (loose coupling)
- **Orchestration:** Central saga coordinator directs steps

### CQRS (Command Query Responsibility Segregation)

```
Write side:  POST /orders → Order Service → orders_db
Read side:   GET  /orders → Order Read Model → orders_read_db (denormalized)
Sync:        OrderPlaced event → update read model
```

---

## Observability in Microservices

```mermaid
flowchart TB
    REQ[Request ID: abc-123] --> GW[Gateway]
    GW -->|trace: abc-123| S1[Service A]
    S1 -->|trace: abc-123| S2[Service B]
    S2 -->|trace: abc-123| S3[Service C]

    S1 --> LOG[Centralized Logs]
    S2 --> LOG
    S3 --> LOG
    S1 --> TRACE[Jaeger Trace Viewer]
    S2 --> TRACE
    S3 --> TRACE
```

**Every request gets a trace ID** propagated via gRPC metadata or HTTP headers.

```go
// Propagate trace context
md := metadata.Pairs("x-trace-id", traceID)
ctx = metadata.NewOutgoingContext(ctx, md)
```

---

## Full Microservices Stack (Capstone Reference)

| Layer | Technology |
| :--- | :--- |
| Language | Go 1.22+ |
| API (external) | REST + Gin |
| API (internal) | gRPC + Protobuf |
| Database | PostgreSQL per service |
| Cache | Redis |
| Message Bus | NATS or Kafka |
| Service Mesh | Istio (optional) |
| Container | Docker multi-stage |
| Orchestration | Kubernetes |
| CI/CD | GitHub Actions |
| Logging | Structured JSON → ELK/Loki |
| Tracing | OpenTelemetry + Jaeger |
| Metrics | Prometheus + Grafana |

---

## Microservices Checklist

Before going to production:

- [ ] Each service has its own `go.mod` and Dockerfile
- [ ] Health check endpoints (`/health`, `/ready`)
- [ ] gRPC + REST contracts defined in `.proto` / OpenAPI
- [ ] Context timeouts on all external calls
- [ ] Circuit breakers on critical dependencies
- [ ] Structured logging with correlation IDs
- [ ] Distributed tracing enabled
- [ ] Database migrations automated
- [ ] CI runs tests + lint on every PR
- [ ] K8s manifests with resource limits

---

## Module Reference

| Topic | Location |
| :--- | :--- |
| gRPC server/client demo | `10_microservices/` |
| Clean architecture | `09_backend_architecture/` |
| REST API | `06_rest_api_and_gin/` |
| Auth/JWT | `07_authentication_and_security/` |
| Docker/K8s deploy | `13_docker_and_kubernetes/` |
| Capstone e-commerce | `15_capstone_projects/` |

---

## Further Reading

- [ARCHITECTURE.md](./ARCHITECTURE.md) — system design patterns
- [LEARNING_PATH.md](./LEARNING_PATH.md) — module order
- [gRPC Go Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [12-Factor App](https://12factor.net/)
