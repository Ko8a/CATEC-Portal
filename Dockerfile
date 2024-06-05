# Start from the official Go image for building the application
FROM golang:latest AS build

WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Use the official Postgres image for the final stage
FROM postgres:latest

# Set the working directory in the final image
WORKDIR /app

# Copy the built Go application from the build stage
COPY --from=build /app/main ./

# Copy the configuration files into the image
COPY configs /app/configs
COPY .env /app/
COPY migrate.sh .

# Set execute permissions for the migration script
RUN chmod +x migrate.sh

# List the contents of the /app directory for debugging
RUN ls -la /app

# Expose the application port
EXPOSE 8080

# Command to run the Go application
CMD ./main
