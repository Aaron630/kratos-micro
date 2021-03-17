set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-X main.Version=1.0.0 -X main.Name=gateway" -o ./bin/ ./...
pause