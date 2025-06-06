# Base Image
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /code

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main cmd/*

# Start a new stage from scratch
FROM scratch

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /code/main .

# Expose port 7543 to the outside world
EXPOSE 7543

# Entry point to run the binary
ENTRYPOINT ["./main"]