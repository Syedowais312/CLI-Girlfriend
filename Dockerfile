# Stage 1: Build
FROM golang:1.25.4-alpine AS builder

WORKDIR /build

# Copy all project files
COPY . .

# Download dependencies
RUN go mod download

# Build the binary (using local module path, Go should resolve it automatically)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags="-w -s" -o my-girlfriend

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS calls to Gemini API
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /build/my-girlfriend /app/my-girlfriend

# Make binary executable
RUN chmod +x /app/my-girlfriend

# Set entrypoint
ENTRYPOINT ["/app/my-girlfriend"]

# Default command (can be overridden)
CMD ["chat", "--help"]
