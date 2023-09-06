#!/bin/bash

# Build cartservice
echo 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cartservice -gcflags "all=-N -l" .'
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cartservice -gcflags "all=-N -l" .