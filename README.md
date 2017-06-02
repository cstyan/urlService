# Building
We assume you already have a box with a Golang install (I am using Go v1.8 on base Ubuntu).
	go get github.com/gorilla/mux
	go build

A helper script is provided to run the service as a docker container.

# Running
The service listens on port 8080 by default. The endpoints `/urlinfo/v1/{url}` and `/urlinfo/v1/upload` are implemented.
A helper script for running the service in docker is provided, it maps hardcoded ports at the moment.