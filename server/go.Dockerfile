# syntax=docker/dockerfile:1
FROM golang:1.17-buster as builder

WORKDIR /server
RUN apt-get -y update 
RUN apt-get install -y make


# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.mod ./
COPY go.sum ./
# Copy local code to the container image.
COPY . ./

RUN go mod download

# Build the binary.
RUN go build -v -o /docker-server 

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /docker-server  /docker-server 

EXPOSE 8080
# Run the web service on container startup.
CMD ["/docker-server"]
