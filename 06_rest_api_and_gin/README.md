# 06. REST API with Gin

> **Time:** 4–5 hours | **Prerequisites:** [Module 05](../05_databases_sql_nosql/README.md)

Build production-style REST APIs with the Gin framework — routing, middleware, validation, full CRUD.

---

## What You Will Learn

- [ ] REST principles (resources, HTTP verbs, status codes)
- [ ] Gin router and handler functions
- [ ] JSON request binding and validation
- [ ] Custom middleware (logging)
- [ ] Full CRUD: GET, POST, PUT, DELETE

---

## Step-by-Step Lessons

### Step 1 — Start the server

```bash
cd 06_rest_api_and_gin
go mod tidy
go run main.go
```

Server runs at `http://localhost:8080`

### Step 2 — Test each endpoint

Open a **second terminal** and run:

```bash
# Health check
curl http://localhost:8080/ping

# List all albums
curl http://localhost:8080/albums

# Get one album
curl http://localhost:8080/albums/1

# Create album
curl -X POST http://localhost:8080/albums \
  -H "Content-Type: application/json" \
  -d "{\"id\":\"4\",\"title\":\"Kind of Blue\",\"artist\":\"Miles Davis\",\"price\":29.99}"

# Update album
curl -X PUT http://localhost:8080/albums/4 \
  -H "Content-Type: application/json" \
  -d "{\"title\":\"Kind of Blue (Remaster)\",\"artist\":\"Miles Davis\",\"price\":34.99}"

# Delete album
curl -X DELETE http://localhost:8080/albums/4
```

### Step 3 — Test validation

```bash
# Missing required field — should return 400
curl -X POST http://localhost:8080/albums \
  -H "Content-Type: application/json" \
  -d "{\"id\":\"5\",\"price\":10}"
```

### Step 4 — Read middleware

Open `requestLogger()` — it wraps every request and logs method, path, status, duration.

---

## API Reference

| Method | Path | Description | Status |
| :--- | :--- | :--- | :---: |
| GET | `/ping` | Health check | 200 |
| GET | `/albums` | List all | 200 |
| GET | `/albums/:id` | Get by ID | 200 / 404 |
| POST | `/albums` | Create | 201 / 400 |
| PUT | `/albums/:id` | Update | 200 / 404 |
| DELETE | `/albums/:id` | Delete | 204 / 404 |

---

## Exercises

See [exercises/EXERCISES.md](./exercises/EXERCISES.md)

---

## Checkpoint

Before Module 07:

- [ ] Implement all CRUD endpoints from memory
- [ ] Use correct HTTP status codes
- [ ] Write custom middleware
- [ ] Validate JSON with binding tags

---

## Next Module

→ [07 Authentication & Security](../07_authentication_and_security/README.md)
