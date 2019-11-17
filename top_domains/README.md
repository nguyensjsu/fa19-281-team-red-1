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
        - E.g. `topdomains`
    - Deploy to a namespace: `make stack-up-namespace`

## Tests
- Reference [Makefile](./Makefile)
- In the makefile, there are commands for running curL commands to test the API, after `make run-all` or `make stack-up-namespace` is successfully ran
    - `make test-top-5` is the probably the most important test
    - Individual tests are also included

## MongoDB Cluster Setup
- A MongoDB replica set cluster is setup using 3 AWS EC2 instances (1 primary, 2 secondaries)
    - Setup notes:
        - A docker host EC2 instance is in Public Subnet (Same VPC)
        - 3 MongoDB EC2 instances in the Private Subnet  (Same VPC)
    - Reference is here: [MongoDB Cluster Setup](https://github.com/paulnguyen/cmpe281/blob/master/labs/lab5/aws-mongodb-replica-set.md)


## Kubernetes Notes
- There are a couple ways (probably more) to deploy:
- `docker stack` in combination with the `docker-compose.yaml` file
    - `docker stack` doesn't support some of the native `docker-compose` commands, such as `build`, so for now, push the image to a repo before calling any `docker stack` command.
        - You can’t build new images using the stack commands. It need pre-built images to exist. So docker-compose is better suited for development scenarios. [Ref](https://vsupalov.com/difference-docker-compose-and-docker-stack/)
    - Create a namespace (`topdomains`), note kubernetes doesn't accept underscores in the name
- Generate Kubernetes yaml files using `kompose` and `docker-compose.yaml`
    - To add..
- Make persistent data volumes for mongodb pods, to ensure that restarts always have the right data
## Misc.
- **Go Modules**
    - I utilize Go modules in `go.mod` and `go.sum` for managing packages

## Resources
- https://www.melvinvivas.com/converting-a-mongodb-docker-compose-file-to-a-kubernetes-deployment/