# Use an official Golang image as the base
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Use debian as the base image
FROM debian:latest
RUN apt update && apt install -y curl

WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]