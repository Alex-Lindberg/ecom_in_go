# Use the official Go image as a base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download all dependencies. They will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

### Start a new stage from scratch ###
FROM golang:latest  

WORKDIR /app

# Install air for hot reloading
RUN go install github.com/cosmtrek/air@latest

# Copy the binary file from the previous stage
COPY --from=builder /app/main .

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. They will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["air", "-c", ".air.toml"]
