# Stage 1 — builder: compiles the binary.
# Use this stage to extract bin/ais without
# installing Go locally (see mk/container.mk).
FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath \
      -ldflags "-s -w" \
      -o /ais .

# Stage 2 — runtime: minimal image, binary only.
# Compatible with both podman and docker.
FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /ais /usr/local/bin/ais
ENTRYPOINT ["ais"]
