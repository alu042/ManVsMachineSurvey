# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD inside the container
COPY go.* .

COPY .env /app/.env

# Download all the dependencies
RUN go mod download && go mod verify

COPY ./cmd ./cmd

RUN go build -v -o main ./cmd

# Install wait-for-it script
ADD https://github.com/vishnubob/wait-for-it/raw/master/wait-for-it.sh /usr/local/bin/wait-for-it
RUN chmod +x /usr/local/bin/wait-for-it

# Use wait-for-it script to wait for the database before starting the application
CMD ["sh", "-c", "wait-for-it postgres:5432 -- ./main"]

# Command to run the executable
# CMD ["./main"]
