# Module 07 — Exercises

## Exercise 1: Token Refresh

Add `POST /api/refresh` that accepts a valid (non-expired) token and returns a new one.

## Exercise 2: Rate Limiter

Add middleware that allows max 5 login attempts per IP per minute.

## Exercise 3: Password Rules

Enhance signup validation: password must have 1 uppercase, 1 number, min 8 chars.

## Exercise 4: Tests

Write tests for:
- Signup with weak password → 400
- Login with wrong password → 401
- Profile without token → 401
- Admin route as regular user → 403
