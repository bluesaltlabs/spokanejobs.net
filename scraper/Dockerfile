FROM golang:1-alpine

# Set working directory for COPY
WORKDIR /scraper

# Copy source code
COPY . .

# Install Go Dependencies
RUN go mod download
RUN apk add --no-cache git openssh

# Configure git for automated commits
RUN git config --global user.email "scraper@bedrock.local" && \
    git config --global user.name "Bedrock Scraper"

# Run
CMD ["go", "run", "./cmd/main.go"]
