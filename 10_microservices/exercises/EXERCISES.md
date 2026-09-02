# Module 10 — Exercises

## Exercise 1: Custom Proto

Create `proto/greeter.proto` with `SayHello` RPC.
Generate Go code:
```bash
protoc --go_out=. --go-grpc_out=. proto/greeter.proto
```

## Exercise 2: gRPC Status Codes

Return `codes.InvalidArgument` when name is empty in `SayHello`.

## Exercise 3: Logging Interceptor

Add a unary server interceptor that logs each RPC method and duration.

## Exercise 4: Client Timeout

Use `context.WithTimeout` on the client — verify server receives cancellation.

## Exercise 5: Order Service gRPC

Define `OrderService` proto with `CreateOrder` and `GetOrder` RPCs.
Implement server with in-memory storage.
