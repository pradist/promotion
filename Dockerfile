# Stage 1: Build the Go binary
FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Stage 2: Minimal image
FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080
ENTRYPOINT ["./app"]
