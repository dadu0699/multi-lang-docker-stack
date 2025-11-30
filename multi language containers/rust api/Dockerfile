# Stage 1: Build a static Rust binary with musl
FROM rust:1.91.1-alpine AS builder

# Install system dependencies required to build a static binary with musl.
RUN apk add --no-cache \
    musl-dev \
    pkgconfig \
    openssl-dev \
    ca-certificates

# Set the working directory where the code will be compiled.
WORKDIR /app

# Copy manifest files first to leverage Docker layer caching for dependencies.
COPY Cargo.toml .
COPY Cargo.lock .

# Copy the actual project source code.
COPY src ./src

# Add the musl target for static builds.
RUN rustup target add x86_64-unknown-linux-musl

# Build the application in release mode for the musl target (static binary).
RUN cargo build --release --target x86_64-unknown-linux-musl

# Stage 2: Minimal runtime image using distroless
FROM gcr.io/distroless/static-debian13:nonroot

# Set the working directory inside the container.
WORKDIR /app

# Copy the statically linked binary from the builder stage.
COPY --from=builder /app/target/x86_64-unknown-linux-musl/release/rust-api /app/rust-api

# Set environment variables used by the application.
ENV RUST_LOG=info

# Document the port the service listens on.
EXPOSE 8080

# Run as non-root user for better security
USER nonroot:nonroot

# Use ENTRYPOINT so additional flags can be appended during `docker run`.
ENTRYPOINT ["/app/rust-api"]
