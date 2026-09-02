# 05. Databases (SQL & NoSQL)

> **Time:** 4–6 hours | **Prerequisites:** [Module 04](../04_testing_and_debugging/README.md)

Learn data persistence with `database/sql`, repository pattern, and SQL best practices.

---

## What You Will Learn

- [ ] `database/sql` — connections, queries, scanning
- [ ] Parameterized queries (SQL injection prevention)
- [ ] Repository pattern (interface + implementation)
- [ ] Connection pooling (`SetMaxOpenConns`)
- [ ] Handling `sql.ErrNoRows`
- [ ] SQLite for local learning (no Docker needed)

---

## Step-by-Step Lessons

| Step | Code Section | Topic |
| :---: | :--- | :--- |
| 1 | `sql.Open("sqlite", ":memory:")` | Open DB connection |
| 2 | `SetMaxOpenConns` | Connection pool |
| 3 | `CREATE TABLE` migration | Schema setup |
| 4 | `INSERT` with `?` placeholders | Parameterized writes |
| 5 | `QueryRow` + `Scan` | Read single row |
| 6 | `Query` + `rows.Next()` | Read multiple rows |

### Step 1 — Install dependency and run

```bash
cd 05_databases_sql_nosql
go mod tidy
go run main.go
```

### Step 2 — Expected output

```
=== 05 Databases (SQL) ===
✅ Created: Alice
✅ Created: Bob
🔍 Found: &{ID:1 Name:Alice Email:alice@example.com}
📋 Total users: 2
✅ Correctly handled: user 99 not found
```

### Step 3 — Read the repository pattern

Notice `UserRepository` interface — business logic depends on the interface, not SQLite.
This lets you swap to PostgreSQL later without changing callers.

---

## Repository Pattern

```mermaid
flowchart LR
    APP[Application] --> IFACE[UserRepository interface]
    IFACE --> SQLITE[SQLiteRepository]
    IFACE --> PG[PostgresRepository]
```

---

## Production Notes

| Topic | Recommendation |
| :--- | :--- |
| Migrations | Use `golang-migrate` or `goose` |
| ORM | GORM for rapid dev; `sqlc` for type-safe SQL |
| PostgreSQL | Replace DSN: `postgres://user:pass@localhost/db` |
| Redis cache | Cache-aside: check Redis → miss → query DB → set Redis |

---

## Exercises

See [exercises/EXERCISES.md](./exercises/EXERCISES.md)

---

## Checkpoint

Before Module 06:

- [ ] Write a parameterized `INSERT` query
- [ ] Handle `sql.ErrNoRows` correctly
- [ ] Explain repository pattern benefits
- [ ] Set connection pool limits

---

## Next Module

→ [06 REST API & Gin](../06_rest_api_and_gin/README.md)
