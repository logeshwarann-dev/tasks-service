# Use the official Golang image as the base image
FROM golang:1.23.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Copy the rest of the application code into the container
COPY . .

# Download all Go dependencies
RUN go mod download


# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the application
CMD ["./main"]
