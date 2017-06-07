#!/bin/bash --noprofile

port="8080"

# in case some containers are lying around
docker rm -f /urlService
docker rm -f /urlService-redis

docker run --name urlService-redis -d redis
# remove -it if we don't need interactive tty
docker run --link urlService-redis:redis --name urlService -p 8080:8080 url-service