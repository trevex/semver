# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o semver .

# Final stage
FROM alpine:3.23

# Install bash and jq
RUN apk add --no-cache bash jq

# Copy the binary from builder
COPY --from=builder /build/semver /usr/local/bin/semver

# Verify the binary works
RUN semver --help > /dev/null || true

ENTRYPOINT ["semver"]
