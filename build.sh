#!/bin/sh

# Nothing special, just building the binary and the docker
GOOS=linux go build
docker build -t com.sootsafe/tex-to-pdf-service:latest .
