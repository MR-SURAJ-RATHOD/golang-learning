# 13. Docker & Kubernetes

> **Time:** 4–6 hours | **Prerequisites:** [Module 12](../12_iot_backend_projects/README.md)

Containerize Go apps with Docker and deploy to Kubernetes.

---

## What You Will Learn

- [ ] Multi-stage Dockerfile for small Go images
- [ ] Docker Compose for local development
- [ ] Health check endpoints (`/health`, `/ready`)
- [ ] Kubernetes Deployment and Service manifests
- [ ] Container best practices

---

## Project Files

```
13_docker_and_kubernetes/
├── main.go              # HTTP server with health endpoints
├── Dockerfile           # Multi-stage build
├── docker-compose.yml   # Local orchestration
├── deployment.yaml      # Kubernetes manifests
└── .dockerignore
```

---

## Step-by-Step Lessons

### Step 1 — Run locally

```bash
cd 13_docker_and_kubernetes
go run main.go
```

Test:
```bash
curl http://localhost:8080/health
curl http://localhost:8080/ready
```

### Step 2 — Build Docker image

```bash
docker build -t go-learning-api .
docker run -p 8080:8080 go-learning-api
```

### Step 3 — Run with Docker Compose

```bash
docker compose up --build
```

Visit `http://localhost:8080/health`

### Step 4 — Deploy to Kubernetes (optional)

Requires a K8s cluster (minikube, kind, or cloud):

```bash
kubectl apply -f deployment.yaml
kubectl get pods
kubectl port-forward svc/go-learning-api 8080:80
```

---

## Dockerfile Explained

| Stage | Purpose |
| :--- | :--- |
| `builder` | Compile Go binary with full SDK |
| `alpine` | Run binary in minimal image (~15MB) |

Key flags: `CGO_ENABLED=0` for static binary, `GOOS=linux` for containers.

---

## Exercises

See [exercises/EXERCISES.md](./exercises/EXERCISES.md)

---

## Checkpoint

Before Module 14:

- [ ] Build a multi-stage Docker image
- [ ] Run app with Docker Compose
- [ ] Explain liveness vs readiness probes
- [ ] Apply K8s deployment manifest

---

## Next Module

→ [14 CI/CD Pipeline](../14_ci_cd_pipeline/README.md)
