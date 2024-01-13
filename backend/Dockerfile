# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD inside the container
COPY go.* .

# Download all the dependencies
RUN go mod download && go mod verify

COPY ./cmd ./cmd

RUN go build -v -o main ./cmd

# Command to run the executable
CMD ["./main"]
