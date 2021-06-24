#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-brief main.go
CGO_ENABLED=0 go build -o go-brief main.go