# Use the official Go image as a parent image
FROM golang:alpine

# Set the Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Compile the application to run on Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o loadbalancer .

# Command to run the executable
CMD ["./loadbalancer"]
