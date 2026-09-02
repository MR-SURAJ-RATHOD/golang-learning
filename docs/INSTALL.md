# Installing Go (Windows-first)

## Option A — Official installer (recommended)
1. Visit https://go.dev/dl and download the Windows MSI for the latest stable version.
2. Run the MSI and follow the installer.
3. Open a new PowerShell and verify:

```powershell
go version
go env
```

## Option B — winget
```powershell
winget install -e --id GoLang.Go
```

## Option C — Chocolatey
```powershell
choco install golang -y
```

## Path & environment
- The installer typically adds `C:\Go\bin` to `PATH`.
- Optional: set GOPATH (defaults to `%USERPROFILE%\go`):
  ```powershell
  setx GOPATH "%USERPROFILE%\go"
  ```
- Verify:
  ```powershell
  go env GOPATH GOROOT
  ```

## Upgrading
- Re-run the MSI or use your package manager: `winget upgrade` or `choco upgrade`.

## Notes
- Use modules (recommended) instead of relying on GOPATH for new projects.
- If you use VS Code, install the Go extension and `gopls` for IDE features.