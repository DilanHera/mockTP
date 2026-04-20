New-Item -ItemType Directory -Force -Path build | Out-Null

# Windows
$env:GOOS="windows"; $env:GOARCH="amd64"
go build -o build/mocktp-windows.exe

# Linux
$env:GOOS="linux"; $env:GOARCH="amd64"
go build -o build/mocktp-linux

# macOS Intel
$env:GOOS="darwin"; $env:GOARCH="amd64"
go build -o build/mocktp-mac

# macOS ARM
$env:GOOS="darwin"; $env:GOARCH="arm64"
go build -o build/mocktp-mac-arm
