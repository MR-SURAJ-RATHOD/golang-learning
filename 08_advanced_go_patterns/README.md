# 08. Advanced Go Patterns

> **Time:** 3–4 hours | **Prerequisites:** [Module 07](../07_authentication_and_security/README.md)

Expert-level Go patterns used in production codebases.

---

## What You Will Learn

- [ ] Generics (`[T any]`) — type-safe reusable code
- [ ] Functional options pattern — clean configuration
- [ ] Middleware chaining — composable handlers
- [ ] Reflection — when to use (and when NOT to)

---

## Step-by-Step Lessons

| Step | Code | Pattern |
| :---: | :--- | :--- |
| 1 | `Stack[T]` | Generics |
| 2 | `NewServer(WithPort(...))` | Functional options |
| 3 | `chain(handler, middlewares...)` | Middleware |
| 4 | `reflect.TypeOf` | Reflection |

### Run

```bash
cd 08_advanced_go_patterns
go run main.go
```

---

## When to Use Each Pattern

| Pattern | Use When |
| :--- | :--- |
| Generics | Reusable data structures (Stack, Queue, Result[T]) |
| Functional Options | Configuring structs with many optional fields |
| Middleware | Cross-cutting concerns (logging, auth, metrics) |
| Reflection | Serialization frameworks, ORMs — avoid in app code |

---

## Exercises

See [exercises/EXERCISES.md](./exercises/EXERCISES.md)

---

## Checkpoint

Before Module 09:

- [ ] Write a generic `Queue[T]`
- [ ] Configure a struct with 3+ functional options
- [ ] Chain 2 middleware functions
- [ ] Explain why reflection is slow and fragile

---

## Next Module

→ [09 Backend Architecture](../09_backend_architecture/README.md)
