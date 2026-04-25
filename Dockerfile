# Stage 1: Build (Optimization: Use Alpine to reduce build-time disk usage)
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency manifests
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o social-forge ./cmd/server/main.go

# Stage 2: Runtime (Reliability: Use Debian + Tini for robust process management)
FROM debian:bookworm-slim

# Install Chromium, fonts, curl (healthcheck), and tini (zombie protection)
RUN apt-get update && apt-get install -y \
    chromium \
    fonts-liberation \
    libnss3 \
    libatk-bridge2.0-0 \
    libxcomposite1 \
    libxrandr2 \
    libgbm1 \
    libasound2 \
    curl \
    tini \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/social-forge .

# Set environment variables
ENV PORT=7860
ENV API_KEY=drashtika-default-key
ENV ENV=production

# Expose the service port
EXPOSE 7860

# Reliability: Use tini as init process
ENTRYPOINT ["/usr/bin/tini", "--"]

# Run the service
CMD ["./social-forge"]
