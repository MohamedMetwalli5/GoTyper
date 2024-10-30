# The official Go image
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o go_typer .

# Use a lightweight image to run the application
FROM alpine:latest

# Set the working directory
WORKDIR /app/

# Copy the binary from the builder stage
COPY --from=builder /app/go_typer .

# Give permission
RUN chmod +x go_typer

# Command to run the application
CMD ["./go_typer"]