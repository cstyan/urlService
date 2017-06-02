#!/bin/bash --noprofile

port="8080"

# if [[ $1 -eq 0 ]; then 
#   port=$1
# fi

# remove -it if we don't need interactive tty
docker run -p 8080:8080 url-service