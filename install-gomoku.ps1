if (-not (Test-Path Env:\GOROOT)) {
	Write-Error @"

Go is not installed.
To install go on this machine,
run the install-go.ps1 script provided in this repo.
"@
	exit 1
}

$directory = Get-Location

${Env:\GOPATH} = ""

if (-not (Test-Path Env:\GOPATH)) {
	mkdir "$directory\go" -Force
	$env:GOPATH = "$directory\go"
}
mkdir "$env:GOPATH\src\gomoku\gomoku" -Force
mkdir "$env:GOPATH\src\gomoku\pepito"  -Force
echo "$env:GOPATH"

Copy-Item .\gomoku\*.go "$env:GOPATH\src\gomoku\gomoku"
Copy-Item .\pepito\*.go "$env:GOPATH\src\gomoku\pepito"

go build gomoku/pepito
