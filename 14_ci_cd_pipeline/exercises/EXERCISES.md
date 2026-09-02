# Module 14 — Exercises

## Exercise 1: Run Full Local CI

```powershell
.\build_and_test.ps1
```
Fix any module that fails build or test.

## Exercise 2: Coverage

```bash
cd 14_ci_cd_pipeline
go test -cover ./...
```

## Exercise 3: Lint

```bash
golangci-lint run
```

## Exercise 4: Pre-Push Script

Create `pre-push.ps1` that runs tests and blocks push if any fail.

## Exercise 5: Docker Build

Build Module 13 Docker image as part of your local check script.
