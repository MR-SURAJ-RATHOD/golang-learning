# Build Stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
# COPY go.sum ./ # Uncomment if using dependencies
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Run Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
