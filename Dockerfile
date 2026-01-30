# ========================
# Build stage
# ========================
FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-server

# ========================
# Runtime stage
# ========================
FROM debian:12-slim

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /build/go-server /app/go-server

EXPOSE 8080
CMD ["/app/go-server"]
