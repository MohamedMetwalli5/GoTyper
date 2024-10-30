# The Go base image
FROM golang:1.18-alpine

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application with all files
RUN go build -o main game.go database_operations.go sender.go receiver.go

# Run the compiled binary
CMD ["./main"]
