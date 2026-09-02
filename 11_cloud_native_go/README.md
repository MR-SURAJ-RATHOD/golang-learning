# 11. Cloud Native Go

> **Time:** 3–4 hours | **Prerequisites:** [Module 10](../10_microservices/README.md)

Build cloud-ready Go applications with 12-factor principles, environment config, and graceful shutdown.

---

## What You Will Learn

- [ ] 12-Factor App principles for Go
- [ ] Configuration from environment variables
- [ ] Graceful shutdown on SIGTERM/SIGINT
- [ ] AWS SDK basics (optional, needs credentials)
- [ ] Cloud deployment readiness

---

## Step-by-Step Lessons

### Step 1 — Run with defaults

```bash
cd 11_cloud_native_go
go run main.go
```

### Step 2 — Run with custom config

```bash
# Windows PowerShell
$env:APP_NAME="my-service"; $env:PORT="9090"; $env:AWS_REGION="ap-south-1"
go run main.go

# Linux/Mac
APP_NAME=my-service PORT=9090 AWS_REGION=ap-south-1 go run main.go
```

### Step 3 — Test graceful shutdown

Run the program and press **Ctrl+C**. Watch the shutdown message.

### Step 4 — AWS S3 (optional)

If you have AWS credentials configured:

```bash
set BUCKET_NAME=your-bucket-name
go run main.go
```

---

## 12-Factor Checklist for Go

| Factor | Go Implementation |
| :--- | :--- |
| Config | `os.Getenv()` — never hardcode secrets |
| Processes | Stateless — state in DB/Redis |
| Port binding | `http.ListenAndServe(":"+port, nil)` |
| Disposability | Graceful shutdown with `signal.NotifyContext` |
| Logs | Structured logging to stdout (`log/slog`) |
| Admin | Health endpoint `/health` |

---

## Exercises

See [exercises/EXERCISES.md](./exercises/EXERCISES.md)

---

## Checkpoint

Before Module 12:

- [ ] Load all config from environment
- [ ] Implement graceful shutdown
- [ ] Explain 12-factor app in your own words

---

## Next Module

→ [12 IoT Backends](../12_iot_backend_projects/README.md)
