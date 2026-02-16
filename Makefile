# Makefile for Doctor Appointment Golang Project

.PHONY: all build run clean docker-build docker-run force

# default target
all: build

# Build the Go application
build:
	go build -o doctor-appointment ./main.go

# Run the application
run:
	./doctor-appointment

# Clean the build artifacts
clean:
	go clean
	del doctor-appointment

# Build the Docker image
docker-build:
	docker build -t doctor-appointment:latest .

# Run the Docker container
docker-run:
	docker run --rm -p 8080:8080 doctor-appointment:latest

# Force clean build artifacts
force: clean build
