# Go Modules (Dependency Management)

- Initialize:
```powershell
go mod init github.com/yourname/yourrepo
```
- Add a dependency:
```powershell
go get github.com/some/module@v1.2.3
```
- Tidy (remove unused, update `go.sum`):
```powershell
go mod tidy
```
- `go.sum` records exact module checksums for reproducible builds.

## Version rules
- Major versions v2+ typically require `/vN` in module paths.

## Replace (local testing)
In `go.mod`:
```
replace github.com/foo/bar => ../local/bar
```

## Troubleshooting
- `GOMODCACHE` is the module cache location (see `go env GOMODCACHE`).