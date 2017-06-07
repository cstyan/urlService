#!/bin/bash --noprofile

set -e

VERSION="0.1"


go get github.com/gorilla/mux
go get github.com/go-redis/redis

# since we're using scratch image we need to staticly compile in C libraries
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

./test.sh

docker build -t url-service:latest -t url-service:"$VERSION" .
docker pull redis:latest