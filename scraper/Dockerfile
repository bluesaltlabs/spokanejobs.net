FROM golang:1-alpine

# Set working directory for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . ./

# Build
RUN go build -o app .

CMD ["./app"]
