ARG app_name
ARG app_version
ARG admin_token

FROM golang:1.15 as build-env
WORKDIR /github.com/kumarabd/service-template
COPY go.mod go.sum ./
RUN go mod download

COPY cmd/main.go cmd/main.go
COPY internal/ internal/
COPY pkg/ pkg/

RUN GO111MODULE=on go build -ldflags "-w -s -X github.com/kumarabd/service-template/internal/config.AdminToken=$admin_token -X github.com/kumarabd/service-template/internal/config.ApplicationName=$app_name -X github.com/kumarabd/service-template/internal/config.ApplicationVersion=$app_version" -a -o service cmd/main.go

FROM gcr.io/distroless/base:nonroot-amd64
COPY internal/config/config.yaml internal/config/config.yaml
ENV DISTRO="debian"
ENV GOARCH="amd64"
WORKDIR /$HOME
COPY --from=build-env /github.com/kumarabd/service-template .
ENTRYPOINT ["./service"]
