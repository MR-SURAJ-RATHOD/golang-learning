# Module 13 — Exercises

## Exercise 1: Smaller Image

Add `-ldflags="-s -w"` to Dockerfile build step. Compare image size before/after.

## Exercise 2: Ingress

Create `ingress.yaml` to route `api.local` to your service.

## Exercise 3: Environment Config

Pass `APP_ENV=production` via Docker Compose and print it in `/health` response.

## Exercise 4: Multi-Service Compose

Add a Redis service to `docker-compose.yml` and connect from Go app.

## Exercise 5: K8s Probes

Add `livenessProbe` and `readinessProbe` to `deployment.yaml` using `/health` and `/ready`.
