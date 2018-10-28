$directory = Get-Location

if (-not (Test-Path Env:\GOROOT)) {
	Write-Warning @"
Go is not installed on this machine.
I will proceed to a local installation in the current directory.
"@
    powershell.exe -ExecutionPolicy Unrestricted .\install-go.ps1 -version 1.11.1
    $env:GOROOT = "$directory\go1.11.1"
}

if ("$directory" -ne "$env:GOPATH\src\gomoku") {
    if (-not (Test-Path  Env:\GOPATH)) {
        mkdir "$directory\go" -Force
        $env:GOPATH = "$directory\go"
    }
    mkdir "$env:GOPATH\src\gomoku\gomoku" -Force
    mkdir "$env:GOPATH\src\gomoku\pepito" -Force
    echo "$env:GOPATH"

    Copy-Item .\gomoku\*.go "$env:GOPATH\src\gomoku\gomoku"
    Copy-Item .\pepito\*.go "$env:GOPATH\src\gomoku\pepito"
}

& "${env:GOROOT}\bin\go.exe" "build" "gomoku/pepito"
Move-Item .\pepito.exe "pbrain-PARIS-Witrand.Gaspard.exe" -Force
