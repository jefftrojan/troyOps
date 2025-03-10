FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /troyops cmd/troyops.go

# Create a minimal image
FROM alpine:3.19

WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /troyops /troyops

# Set the entrypoint
ENTRYPOINT ["/troyops"] 