# URL Service
The goal of this service is to provide an interface to upload and query for URLs, returning some information to the user. One basic use case is white/blacklisting URLs.

Currently the service is implemented in Go, using Gorilla mux to provide a basic REST interface. The service makes use of a k/v store to store the 
URLs and whether or not it's blacklisted (true/false). Ideally the data store can replaced with any remote storage (MySQL, redis, whatever) and
the implementer just has to provide the functions from the data store interface, and the constraints around what information about a URL is stored 
and how to decide if a URL is white or blacklisted can be decided by the implementer. At the moment redis and go map stores are implemented in the dataStore package.

In the future, the "store info about this URL" endpoint and "get info about this URL" endpoint should probably be separate services, those 
workloads will scale differently.

# Dependencies
Working docker and golang installs

# Building
We assume you already have a box with a Golang install (I am using Go v1.8 on base Ubuntu).

A helper script is provided to build the service as a docker container, check there for manual
steps if you wish to build the binary separately from the doontainer..

# Running
A helper script is provided that runs the service and redis containers via docker, it maps hardcoded ports at the moment.  
The service listens on port 8080 by default. The endpoints `/urlinfo/v1/{url}`, `/urlinfo/v1/whitelist`, and `/urlinfo/v1/blacklist` are implemented.

As an example, you could whitelist urls via curl: `curl -X POST -d "aaa.com,bbbb.com" localhost:8080/urlinfo/v1/whitelist`