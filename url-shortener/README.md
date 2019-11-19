# URL Shortener Service
This service will provide the core functionality of this URL shortener project. It accepts url shorten requests from client and gives shortened url back. Also, it maintains some stats by calling APIs provded by other services, such as user service and top service. 

## Endpoints
- `/`: health check
- `/shorten`: POST on this endpoint will initiate a url shorten
    - Example Request: `{"url" : "https://www.google.com"}`
    - Example Response: `{"ShortUrl" : "http://shortenserver/unshorten/abcdefg"}`
- `/unshorten/XXX`: GET on this endpoint will redirect the client to the origin url
    - Example Request: `http://shortenserver/unshorten/abcdefg`

## How to run
- Reference [Makefile](./Makefile)
- Local:
    - `make run`
- Docker-compose
    - `make docker-run`

## Tests
- Reference [Makefile](./Makefile)
- In the makefile, there are commands for running curL commands to test the API, after `make run` or `make docker-run` is successfully ran
    - `make test-shorten` is the probably the most important test
    - Individual tests are also included

## MongoDB Cluster Setup
- A MongoDB replica set cluster is setup using 3 AWS EC2 instances (1 primary, 2 secondaries)
    - Setup notes:
        - A docker host EC2 instance is in Public Subnet (Same VPC)
        - 3 MongoDB EC2 instances in the Private Subnet  (Same VPC)
    - Reference is here: [MongoDB Cluster Setup](https://github.com/paulnguyen/cmpe281/blob/master/labs/lab5/aws-mongodb-replica-set.md)

## Misc.
- **Go Modules**
    - I utilize Go modules in `go.mod` and `go.sum` for managing packages