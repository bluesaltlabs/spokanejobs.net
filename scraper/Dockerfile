FROM golang:1-alpine

# Set working directory for COPY
WORKDIR /

# Copy source code
COPY . .

# Install Go Dependencies
RUN go mod download

# Run
CMD ["go", "run", "./cmd/main.go"]
