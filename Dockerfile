# The Go base image
FROM golang:1.18-alpine

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download dependencies
RUN go mod download

# Run the application using multiple Go files
CMD ["go", "run", "game.go", "database_operations.go", "sender.go", "receiver.go"]
