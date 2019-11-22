# Yet Another URL Shortener
For the CMPE 281 Team Hackathon project, we created a URL shortener application with 3 core microservices. The project is implemented with React.js as the frontend, Golang as the backend API microservices, and MongoDB as the NoSQL database. In addition, we use Docker containers (containers are the future!) for packaging, deploying and running our microservices. All of these components are hosted on AWS. For more details, see the architecture diagram

## Microservice Architecture Diagram
![](docs/architecture.png)

## Kong Setup and Configuration

### Set up
 - sudo docker network create kong-net
 - sudo docker run -d --name kong-database --network=kong-net -p 5432:5432 -e "POSTGRES_USER=kong" -e "POSTGRES_DB=kong" postgres:9.6
 - sudo docker run --rm --network=kong-net -e "KONG_DATABASE=postgres" -e "KONG_PG_HOST=kong-database" kong:latest kong migrations bootstrap
 - sudo docker run -d --name kong --network=kong-net -e "KONG_DATABASE=postgres" -e "KONG_PG_HOST=kong-database" -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" -e "KONG_PROXY_ERROR_LOG=/dev/stderr" -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" -p 8000:8000 -p 8443:8443 -p 8001:8001 -p 8444:8444 kong:latest

### Configuration
 - curl -i -X POST --url http://localhost:8001/services/   -d 'name=serviceName' -d 'url=serviceUrl'
 - curl -i -X POST --url http://localhost:8001/routes/ -d 'paths[]=path1&paths[]=path2' -d 'strip_path=false' -d 'service.id=serviceId'

Replace the "serviceName", "serviceUrl", "path" and "serviceId" as you need.

## Core Features / Microservices
- URL shortener
- Most popular domains
- User authentication and profile with history of past URLs
