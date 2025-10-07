FROM golang:1.24-alpine AS builder

WORKDIR /app
 
# Copy module files from the build context (assumes build context is the repo root)
# Use paths relative to the build context (no "..")
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source files from the build context
COPY . .

# Build the server located in ./cmd/server (adjust path if your main package is elsewhere)
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app ./cmd/server

EXPOSE 8080
CMD ["/go-app"]