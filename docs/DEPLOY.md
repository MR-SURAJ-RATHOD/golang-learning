# Building & Deployment

## Build for current OS
```powershell
go build -o myapp.exe ./...
```

## Cross-compilation (example for Linux AMD64 on Windows PowerShell)
```powershell
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o myapp-linux
```

## Strip symbols (smaller binary)
```powershell
go build -ldflags "-s -w" -o myapp
```

## Containerization
- Create a small Dockerfile to build and copy the static binary into a lightweight image (alpine/debian).

## Distribution
- Release artifacts for relevant OS/arch combos and include checksums.