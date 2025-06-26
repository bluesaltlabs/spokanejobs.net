FROM golang:1-alpine

# Set working directory for COPY
WORKDIR /app

# Copy source code
COPY . .

# Install Go Dependencies
RUN go mod download


# Build
RUN go build -o scraper .

CMD ["./scraper"]
