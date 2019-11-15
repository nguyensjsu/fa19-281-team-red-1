
all: clean

clean: 
	find . -name 'top_domains' -type f -exec rm -f {} \; 
	go clean

setup:
	go mod init src/top_domains

build:
	go build -o bin/top_domains src/top_domains/*

run: build
	go run src/top_domains/*.go

format:
	go fmt src/top_domains/*

install:
	go install src/top_domains/*

start:
	./top_domains 

test-ping:
	curl localhost:3000/ping

test-top:
	curl localhost:3000/top

test-domains:
	curl localhost:3000/domains

docker-build: 
	docker build -t top_domains .
	docker images

docker-run:
	docker run --network top_domains --name top_domains -td -p 3000:3000 top_domains
	docker ps

docker-network-prune:
	docker network prune

network-create:
	docker network create --driver bridge top_domains

rabbitmq-run:
	docker run --name rabbitmq --network top_domains \
			   -p 8080:15672 -p 4369:4369 -p 5672:5672 \
			   -d rabbitmq:3-management

mongodb-run:
	docker run --name mongodb --network top_domains -p 27017:27017 -d mongo:3.7

mongodb-run-local:
	docker run --name mongodb -p 27017:27017 -d mongo:3.7


docker-network:
	docker network ls

docker-network-inspect:
	docker network inspect host

docker-shell:
	docker exec -it top_domains bash 

docker-clean:
	docker stop mongodb
	docker stop rabbitmq
	docker rm mongodb
	docker rm rabbitmq
	docker stop top_domains
	docker rm top_domains
	docker rmi top_domains

docker-ip:
	docker-machine ip

docker-ps:
	 docker ps --all --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t"

docker-ps-ports:
	 docker ps --all --format "table {{.Names}}\t{{.Ports}}\t"

run-all:
	docker-compose up --build
