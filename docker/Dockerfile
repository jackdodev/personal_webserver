FROM golang:1.24-alpine AS builder

WORKDIR /app
 
# Copy module files from the build context (assumes you run `docker build` from the repo root)
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the server located in ./cmd/server (adjust path if your main package is elsewhere)
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app ./cmd/server

EXPOSE 8080
CMD ["/go-app"]