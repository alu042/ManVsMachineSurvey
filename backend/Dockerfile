# syntax=docker/dockerfile:1

# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app - change directory to 'cmd' where main.go resides
RUN go build -o main ./cmd

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./cmd/main"]
