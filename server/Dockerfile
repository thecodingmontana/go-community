# Build stage
FROM golang:1.23.5-alpine AS builder

# Install necessary dependencies
RUN apk add --no-cache git gcc musl-dev bash

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod tidy && go mod vendor

# Install sqlc and goose
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy source code
COPY . .

# Generate code using sqlc.
RUN sqlc generate

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache bash curl

WORKDIR /app

# Copy application binary and necessary files
COPY --from=builder /app/main .

COPY --from=builder /app/internal/database/schema ./internal/database/schema
COPY --from=builder /app/internal/database/queries ./internal/database/queries

# Copy goose binary to the final image for migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Expose application port
EXPOSE 8000

# Run migrations and start the application
CMD ["sh", "-c", "goose -dir ./internal/database/schema postgres \"$DATABASE_URL\" up && ./main"]