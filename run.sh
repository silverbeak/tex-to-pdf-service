#!/bin/sh

# Just a basic run command, nothing fancy
docker run -d --name ss-latex1 -p50051:50051 com.sootsafe/tex-to-pdf-service:latest
