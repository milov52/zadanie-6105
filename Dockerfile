# Start from a small, secure base image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Install goose for database migrations
RUN GOBIN=/go/bin go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

# Create a minimal production image
FROM alpine:latest

# It's essential to regularly update the packages within the image to include security patches
# Also install bash to run wait-for-it.sh
RUN apk update && apk upgrade && apk add bash postgresql-client

# Reduce image size
RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

# Avoid running code as a root user
RUN adduser -D appuser
USER appuser

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/cmd/wait-for-it.sh .
COPY --from=builder /app/internal/app/migrations ./migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Expose the port that the application listens on
EXPOSE 8080

# Set environment variables for PostgreSQL connection (these will be passed in at runtime)
ENV POSTGRES_USERNAME=postgres
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5432
ENV POSTGRES_DATABASE=tenders

# Define the DSN for goose migrations
ENV LOCAL_MIGRATION_DSN="postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable"
ENV LOCAL_MIGRATION_DIR="./migrations"

# Run the binary when the container starts
CMD ./wait-for-it.sh ${POSTGRES_HOST}:${POSTGRES_PORT} -- goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v && ./app