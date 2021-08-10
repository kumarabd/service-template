ARG server_version

FROM golang:1.15 as build-env
WORKDIR /github.com/kumarabd/service-template
COPY go.mod go.sum ./
RUN go mod download

COPY cmd/main.go cmd/main.go
COPY internal/ internal/
COPY pkg/ pkg/

RUN GO111MODULE=on go build -ldflags="-w -s" -a -o service cmd/main.go

FROM gcr.io/distroless/base:nonroot-amd64
ENV DISTRO="debian"
ENV GOARCH="amd64"
ENV SERVER_VERSION=$server_version
WORKDIR /$HOME
COPY --from=build-env /github.com/kumarabd/service-template .
ENTRYPOINT ["./service"]
