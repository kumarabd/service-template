check:
	golangci-lint run

test:
	go test ./...

run:
	DEBUG="true" go run cmd/main.go

build:
	go build -ldflags "-w -s -X github.com/kumarabd/service-template/internal/config.AdminToken=test -X github.com/kumarabd/service-template/internal/config.ApplicationName=test -X github.com/kumarabd/service-template/internal/config.ApplicationVersion=test" -a -o service cmd/main.go