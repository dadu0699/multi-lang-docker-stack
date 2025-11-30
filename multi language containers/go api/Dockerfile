# Stage 1: Build the app  
FROM golang:1.25-alpine AS builder

# Install minimal build dependencies
RUN apk add --no-cache ca-certificates git

# Set working directory inside the container
WORKDIR /app

# Copy Go modules definition to leverage Docker layer cache
COPY go.mod go.sum ./

# Download Go module dependencies (cached between builds)
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy the rest of the source code
COPY . .

# Path to the main package (can be overridden at build time)
# Example:
#   docker build --build-arg APP_PATH=./cmd/signsvc -t myapp .
ARG APP_PATH=./cmd/server

# Build a small, statically linked binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" \
    -o /out/app ${APP_PATH}

# Stage 2: Runtime image
FROM gcr.io/distroless/static-debian13:nonroot

# Application working directory
WORKDIR /srv

# Copy the compiled binary from the builder stage
COPY --from=builder /out/app /srv/app

# Environment configuration (read by your Go app)
ENV GIN_MODE=release
ENV PORT=8080

# Document the port the service listens on
EXPOSE 8080

# Run as non-root user for better security
USER nonroot:nonroot

# Use ENTRYPOINT so additional args can be passed at `docker run`
ENTRYPOINT ["/srv/app"]
