# Use the official Golang image as a parent image
FROM golang:1.23-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download
# Download and install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./dashboard ./cmd/server/main.go

# Run templ generation
RUN templ generate

# Use minimal image 'scratch' for the final stage
FROM scratch

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/dashboard .

# Copy src files
COPY --from=builder /app/src ./src

# Expose port 9100
EXPOSE 9000

# Run the binary
CMD ["./dashboard"]