# syntax=docker/dockerfile:1

# Start from the latest golang base image
FROM golang:latest AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app - change directory to 'cmd' where main.go resides
RUN go build -o /app/hello

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/hello .

# Command to run the executable
CMD ["./hello"]
