# Stage 1: Build stage
FROM golang:1.22.1 AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the application
COPY . .

# Optionally copy the .env file if it's needed
COPY .env .  

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Stage 2: Final stage
FROM alpine:latest

WORKDIR /app

# Install certificates to handle secure connections
RUN apk add --no-cache ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .

# Optionally copy the .env file if it's needed
COPY --from=builder /app/.env .

# Expose the application port
EXPOSE 8087

# Command to run the executable
CMD ["./myapp"]
