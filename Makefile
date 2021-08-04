check:
	golangci-lint run

test:
	go test ./...

run:
	DEBUG="true" go run cmd/main.go

build:
	go build -o jarvis cmd/main.go