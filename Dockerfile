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
COPY --from=builder /workspaces/fintrax/fintrax_example.yml /config/fintrax.yml
VOLUME [ "/config" ]
EXPOSE 8080/tcp
EXPOSE 8443/tcp
CMD ["/app/fintrax", "daemon"]
