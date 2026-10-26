# syntax=docker/dockerfile:1

FROM golang:1.27-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
ARG VERSION=local
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Version=${VERSION}" -o build/metio-operator ./cmd/metio-operator

FROM alpine:3.24
RUN apk add --no-cache ca-certificates && \
    adduser -D -u 65532 nonroot
EXPOSE 8080 8081
WORKDIR /app
COPY --from=builder /app/build/metio-operator /app/metio-operator
USER nonroot:nonroot
ENTRYPOINT ["/app/metio-operator"]