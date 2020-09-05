# Use of one the offical Golang images to create a build artifact.
FROM golang:1.14-alpine3.11 as builder

# Fetch certs to allow use of TLS.
RUN apk add -U --no-cache ca-certificates git

WORKDIR /build

# Fetch dependencies.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy all files.
COPY . .

# Build.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/server

FROM scratch

# Copy the binary to the production image from the builder stage.
COPY --from=builder /build/main /main

# Copy CA certificates (handy if we want to use TLS).
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run the service on container startup.
CMD ["/main"]
