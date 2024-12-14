# Build phase (compiling the Go application)
FROM golang:1.23.4-alpine AS build

# Install build dependencies
RUN apk add --no-cache build-base git

WORKDIR /app

# Copy only dependency files first to leverage Docker cache
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod tidy && go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application for Linux (amd64 architecture)
RUN GOOS=linux GOARCH=amd64 go build -o main cmd/server/main.go

# Final image (smaller, only runtime dependencies)
FROM alpine:3.21

# Install only the necessary runtime dependencies (ca-certificates for SSL)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the built binary from the build phase
COPY --from=build /app/main .

# Set Gin mode to 'release' for production
ENV GIN_MODE=release

# Expose the port the application will run on
EXPOSE 8080

# Run the built application binary
CMD ["./main"]
