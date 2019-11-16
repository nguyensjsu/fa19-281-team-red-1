# Top Domain API
This Go API will take a JSON request with a url, and track the number of times a certain domain was hit. It stores this data in a NoSQL MongoDB database. Also, it returns the top 5 domains that were hit. 

## Endpoints
- `/ping`: health check
- `/domains`: GET all domains stored in MongoDB collection `top_domains`
- `/url`: POST a url to the server 
    - Example Request: `{"url": "http://nytimes.com/f82jd0203k349dk"}`
- `/top`: GET top 5 most hit domains

## How to run
- Reference [Makefile](./Makefile)
- Local:
    - `make run`
- Docker-compose
    - `make run-all`
- Kubernetes
    - Important note: make sure to push your image to a repo, for `docker stack deploy` to work
    - Create a name space: `make create-namespace`
    - Deploy: `make stack-up-namespace`

## Tests
- Reference [Makefile](./Makefile)
- In the makefile, there are commands for running curL commands to test the API, after `make run-all` or `make stack-up-namespace` is successfully ran
    - `make test-top-5` is the probably the most important test
    - Individual tests are also included

## Kubernetes Notes
- Create a namespace (`topdomains`), note kubernetes doesn't accept underscores in the name

## Misc.
- **Go Modules**
    - I utilize Go modules in `go.mod` and `go.sum` for managing packages
