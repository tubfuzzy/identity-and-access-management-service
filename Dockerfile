# Use arm64v8/golang as base image with Go 1.21.6
FROM arm64v8/golang:1.21.6-alpine3.15

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp .

# Expose the port the application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
