FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o humanos-api ./cmd/api

# Final stage - minimal runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy built binary from builder
COPY --from=builder /app/humanos-api .

# Copy JSON schemas
COPY shared/schemas/ ./shared/schemas/

EXPOSE 8080

CMD ["./humanos-api"]
