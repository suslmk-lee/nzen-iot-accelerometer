# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o accelerometer .

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/accelerometer .
COPY config.properties .

CMD ["./accelerometer"]
