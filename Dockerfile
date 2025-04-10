# Use the official Golang image as a builder image
# Using a specific version is recommended for reproducibility
FROM golang:1.24-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY *.go ./

# Build the Go app
# -ldflags="-w -s" strips debug information, reducing binary size
# CGO_ENABLED=0 builds a statically linked binary (important for minimal base images like scratch or alpine)
# GOOS=linux ensures the binary is built for the Linux environment used by Cloud Run
RUN CGO_ENABLED=0 GOOS=linux go build -v -ldflags="-w -s" -o /go-app .

# --- Start a new, smaller stage from alpine for the final image ---
# Alpine Linux is a small, simple, and secure base image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go-app /go-app

# Expose port 8080 to the outside world (documentary, Cloud Run uses the PORT env var)
EXPOSE 8080

# Command to run the executable
# The application listens on the PORT environment variable provided by Cloud Run.
ENTRYPOINT ["/go-app"]