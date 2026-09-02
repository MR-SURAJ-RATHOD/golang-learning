# 07. Authentication & Security

> **Time:** 4–5 hours | **Prerequisites:** [Module 06](../06_rest_api_and_gin/README.md)

Secure your APIs with JWT, bcrypt password hashing, RBAC, and security headers.

---

## What You Will Learn

- [ ] Password hashing with bcrypt
- [ ] JWT token generation and validation
- [ ] Auth middleware in Gin
- [ ] Role-Based Access Control (RBAC)
- [ ] Security headers
- [ ] Environment-based secrets (12-factor)

---

## Step-by-Step Lessons

### Step 1 — Start server

```bash
cd 07_authentication_and_security
go mod tidy
go run main.go
```

### Step 2 — Sign up a user

```bash
curl -X POST http://localhost:8081/signup \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"alice\",\"password\":\"secret123\"}"
```

### Step 3 — Login and get JWT

```bash
curl -X POST http://localhost:8081/login \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"alice\",\"password\":\"secret123\"}"
```

Copy the `token` from the response.

### Step 4 — Access protected route

```bash
curl http://localhost:8081/api/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### Step 5 — Test RBAC

```bash
# Sign up admin
curl -X POST http://localhost:8081/signup \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"admin\",\"password\":\"admin123\"}"

# Login as admin, get token, then:
curl http://localhost:8081/api/admin \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

Regular users get `403 Forbidden` on `/api/admin`.

### Step 6 — Use environment secret (production)

```bash
set JWT_SECRET=my-super-secret-key   # Windows
export JWT_SECRET=my-super-secret-key  # Linux/Mac
go run main.go
```

---

## API Reference

| Method | Path | Auth | Description |
| :--- | :--- | :---: | :--- |
| POST | `/signup` | No | Register user |
| POST | `/login` | No | Get JWT token |
| GET | `/api/profile` | Yes | User profile |
| GET | `/api/admin` | Yes (admin) | Admin only |

---

## Security Checklist

- [x] Passwords hashed with bcrypt (never store plain text)
- [x] JWT secret from environment variable
- [x] Token expiration (15 minutes)
- [x] Security headers (X-Frame-Options, etc.)
- [x] RBAC for admin routes
- [ ] HTTPS in production (use reverse proxy)
- [ ] Rate limiting on login endpoint

---

## Exercises

See [exercises/EXERCISES.md](./exercises/EXERCISES.md)

---

## Checkpoint

Before Module 08:

- [ ] Explain JWT structure (header.payload.signature)
- [ ] Implement auth middleware from scratch
- [ ] Hash and verify passwords with bcrypt
- [ ] Return 401 vs 403 correctly

---

## Next Module

→ [08 Advanced Go Patterns](../08_advanced_go_patterns/README.md)
