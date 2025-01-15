# Use an official Golang image to build the Go app
FROM golang:alpine AS builder

# Set the working directory
WORKDIR /app

# Copy backend into the container at /app
COPY . .

# Build the Go app
RUN go build -o twitchtoken .

# Use a minimal image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built Go app from the builder stage
COPY --from=builder /app/twitchtoken .
COPY --from=builder /app/public ./public

# Set environment variables
ENV CLIENT_ID=""
ENV REDIRECT_URI=""

EXPOSE 8000

CMD ["./twitchtoken"]