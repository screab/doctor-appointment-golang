# Dockerfile

# Stage 1: Build the Go application
FROM golang:1.18 AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o doctor-appointment main.go


# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /root/

# Copy the executable from the builder stage
COPY --from=builder /app/doctor-appointment .

# Command to run the executable
CMD ["./doctor-appointment"]
