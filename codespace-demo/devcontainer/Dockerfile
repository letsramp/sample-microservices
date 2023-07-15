FROM golang:1.20.3 AS build-dlv

RUN apt-get update && apt-get install -y --no-install-recommends curl 
RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.2
