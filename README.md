# URL Service
The goal of this service is to provide an interface to upload and query for URLs, returning some information to the user. One basic use case is white/blacklisting URLs.

Currently the service is implemented in Go, using Gorilla mux to provide a basic REST interface. The service makes use of a map to store the 
URLs and whether or not it's blacklisted (true/false). Ideally the data store can replaced with any remote storage (MySQL, redis, whatever) and
the implementer just has to provide the functions from the data store struct, and the constraints around what information about a URL is stored 
and how to decide if a URL is white or blacklisted can be decided by the implementer.

In the future, the "store info about this URL" endpoint and "get info about this URL" endpoint should probably be separate services, those 
workloads will scale differently.

# Building
We assume you already have a box with a Golang install (I am using Go v1.8 on base Ubuntu).
	go get github.com/gorilla/mux
	go build

A helper script is provided to run the service as a docker container.

# Running
The service listens on port 8080 by default. The endpoints `/urlinfo/v1/{url}` and `/urlinfo/v1/upload` are implemented.
A helper script for running the service in docker is provided, it maps hardcoded ports at the moment.