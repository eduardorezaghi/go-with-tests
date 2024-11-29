# Start from the official Golang base image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
# COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o server ./server.go && chmod +x server

# Expose port 6767 to the outside world
EXPOSE 6767

# Command to run the executable
CMD ["./server"]

# Run the container based on this iamge
# docker build -t go-server .
# docker run -p 6767:6767 go-server