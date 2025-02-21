# syntax=docker/dockerfile:1

# Build Stage
ARG GO_VERSION=1.24.0
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download -x

ARG TARGETARCH
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -trimpath -ldflags="-s -w" -o /bin/server ./cmd/server/main.go

# Production Stage
FROM alpine:latest AS prod

RUN apk --no-cache add ca-certificates tzdata && update-ca-certificates

ARG UID=10001
RUN adduser --disabled-password --gecos "" --home "/nonexistent" \
    --shell "/sbin/nologin" --no-create-home --uid "${UID}" appuser
USER appuser

COPY --from=build /bin/server /bin/

EXPOSE 50051

ENTRYPOINT ["/bin/server"]

# Development Stage
FROM golang:1.24.0 AS dev

RUN \
    --mount=type=cache,target=/var/lib/apt/lists \
    --mount=type=cache,target=/var/cache/apt/archives \
    apt-get update \
    && apt-get install --no-install-recommends -y \
    apt-transport-https \
    build-essential \
    curl \
    git \
    vim \
    wget \
    zip

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server ./cmd/server/main.go

COPY . .

RUN go install github.com/air-verse/air@latest

EXPOSE 50051
CMD ["air", "-c", ".air.toml"]
