# Stage 1: Build the Go binary
FROM golang:alpine AS builder

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o myapp .


# Stage 2: Create a minimal container to run the Go binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Command to run the executable
CMD ["./myapp"]
