# Module 09 — Exercises

Complete these to deepen your architecture skills.

## Exercise 1: Postgres Adapter

Create `PostgresOrderRepo` implementing `OrderRepository`.
Use SQLite from Module 05 as a simpler alternative if Postgres isn't available.

## Exercise 2: HTTP Adapter

Add a Gin handler `POST /orders` that calls `OrderService.CreateOrder`.
Return JSON `{"id": "...", "status": "created"}`.

## Exercise 3: Domain Validation

Move `amount > 0` validation into the `Order` entity:
```go
func (o *Order) Validate() error
```

## Exercise 4: Cancel Order

Add `CancelOrder(id string) error` to the service.
Order cannot be cancelled if amount is 0.

## Exercise 5: Package Layout

Split `main.go` into:
```
internal/domain/order.go
internal/port/repository.go
internal/usecase/order_service.go
internal/adapter/memory_repo.go
cmd/api/main.go
```
