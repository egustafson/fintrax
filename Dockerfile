# Dockerfile 
#
# Build Stage
#
FROM golang:latest AS builder

WORKDIR /workspaces/fintrax

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./bin/fintrax .

#
# Package Stage
#
FROM debian:trixie-slim

WORKDIR /app
COPY --from=builder /workspaces/fintrax/bin/fintrax .
EXPOSE 8080
CMD ["./fintrax", "daemon"]
