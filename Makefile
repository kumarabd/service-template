SRC_DIR=pkg
check:
	golangci-lint run

generate_tests:
	# install and setup https://github.com/cweill/gotests
	@find $(SRC_DIR) -type f -name '*.go' -exec bash -c ' \
        for file do \
			gotests -all -w $$file; \
        done' _ {} +
	
generate_proto:
	protoc -I ./api \
	--go_out ./api --go_opt paths=source_relative \
	--go-grpc_out ./api --go-grpc_opt paths=source_relative \
	--openapiv2_out ./api \
	--openapiv2_opt logtostderr=true \
	--openapiv2_opt generate_unbound_methods=true \
	./api/sample/sample.proto

generate_gql:
	cd api/graphql && \
	go run github.com/99designs/gqlgen generate && \
	cd ../..

test:
	go test ./... -cover

run:
	DEBUG="true" go run cmd/main.go

build:
	go build -ldflags "-w -s -X github.com/kumarabd/service-template/internal/config.ApplicationName=test -X github.com/kumarabd/service-template/internal/config.ApplicationVersion=test" -a -o service cmd/main.go

.PHONY: generate_tests