FROM golang:1.12 as build-env
# All these steps will be cached
RUN mkdir /tex-to-pdf-service
WORKDIR /tex-to-pdf-service
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN GOOS=linux go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o /go/bin/tex-to-pdf-service

FROM ubuntu

LABEL maintainer="Kristofer Jarl <kristofer@sootsafe.com>"

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y latexmk \
    texlive-science \
    texlive-latex-extra \
    texlive-fonts-recommended

EXPOSE 50051

RUN mkdir -p /temp

COPY --from=build-env /go/bin/tex-to-pdf-service /tex-to-pdf-service

ENTRYPOINT [ "/tex-to-pdf-service" ]