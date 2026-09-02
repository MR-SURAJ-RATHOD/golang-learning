$ErrorActionPreference = "Continue"

# Run from repo root: .\build_and_test.ps1

function Build-Module ($Path) {
    Write-Host "Building module: $Path"
    Push-Location $Path
    try {
        go build ./...
        if ($LASTEXITCODE -eq 0) {
            Write-Host "Build Success"
        } else {
            Write-Host "Build Failed"
        }
    } catch {
        Write-Host "Build Failed with Exception"
    } finally {
        Pop-Location
    }
}

function Run-Example ($Path) {
    Write-Host "Running example: $Path"
    Push-Location $Path
    try {
        go run .
        if ($LASTEXITCODE -ne 0) {
             Write-Host "Run Failed"
        }
    } catch {
        Write-Host "Run Failed with Exception"
    } finally {
        Pop-Location
    }
}

Build-Module "01_go_fundamentals"
Run-Example "01_go_fundamentals"

Build-Module "02_modules_and_workspace"
Run-Example "02_modules_and_workspace"

Build-Module "03_concurrency_and_context"
Run-Example "03_concurrency_and_context"

Build-Module "04_testing_and_debugging"
Run-Example "04_testing_and_debugging"

Build-Module "05_databases_sql_nosql"
Run-Example "05_databases_sql_nosql"

Build-Module "06_rest_api_and_gin"
Build-Module "07_authentication_and_security"

Build-Module "08_advanced_go_patterns"
Run-Example "08_advanced_go_patterns"

Build-Module "09_backend_architecture"
Run-Example "09_backend_architecture"

Build-Module "10_microservices"

Build-Module "11_cloud_native_go"
Run-Example "11_cloud_native_go"

Build-Module "12_iot_backend_projects"
Build-Module "13_docker_and_kubernetes"

Build-Module "14_ci_cd_pipeline"
Run-Example "14_ci_cd_pipeline"

Build-Module "15_capstone_projects/impl/iot-platform"

Write-Host "Verification Complete"
