# Use official Go image for build stage
FROM golang:1.24-alpine AS builder

# Install git to fetch dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first (to cache dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source files
COPY . .

# Build the server binary
RUN go build -o server main.go

# Use a minimal base image for running
FROM alpine:latest

# Copy the binary from builder stage
COPY --from=builder /app/server /server

# Expose the port Shadowsocks will listen on
EXPOSE 8388

# Run the Shadowsocks server
CMD ["/server"]
