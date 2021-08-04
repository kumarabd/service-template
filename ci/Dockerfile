ARG server_version
ARG parser_token

FROM golang:1.15 as build-env
WORKDIR /github.com/kumarabd/jarvis
COPY go.mod go.sum ./
RUN go mod download

COPY cmd/main.go cmd/main.go
COPY internal/ internal/
COPY pkg/ pkg/

RUN GO111MODULE=on go build -ldflags="-w -s" -a -o jarvis cmd/main.go

FROM gcr.io/distroless/base:nonroot-amd64
ENV DISTRO="debian"
ENV GOARCH="amd64"
ENV SERVER_VERSION=$server_version
ENV PARSER_TOKEN=$parser_token 
WORKDIR /$HOME
COPY --from=build-env /github.com/kumarabd/service-template/jarvis .
ENTRYPOINT ["./jarvis"]