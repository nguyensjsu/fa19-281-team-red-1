
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

run-binary:
	bin/top_domains

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

test-create-and-update:
	curl -X POST \
  	http://localhost:3000/url \
  	-H 'Content-Type: application/json' \
	-d '{ "url": "http://localhost.com/4f8df838fdfhj29fi" }' 
	curl -X POST \
  	http://localhost:3000/url \
  	-H 'Content-Type: application/json' \
	-d '{ "url": "http://localhost.com/v939dfh3jdfj93fkd" }' 
	curl localhost:3000/domains

test-setup:
	curl -X POST \
  	http://localhost:3000/url \
  	-H 'Content-Type: application/json' \
	-d '{ "url": "http://localhost.com/" }' 
	curl -X POST \
  	http://localhost:3000/url \
  	-H 'Content-Type: application/json' \
	-d '{ "url": "http://localhost.com/4f8df838fdfhj29fi" }' 
	curl -X POST \
  	http://localhost:3000/url \
  	-H 'Content-Type: application/json' \
	-d '{ "url": "http://localhost.com/v939dfh3jdfj93fkd" }' 
	curl -X POST \
  	http://localhost:3000/url \
  	-H 'Content-Type: application/json' \
	-d '{ "url": "http://nytimes.com/f9fj0f3k3fj3f0fkj" }' 
	curl -X POST \
  	http://localhost:3000/url \
  	-H 'Content-Type: application/json' \
	-d '{ "url": "http://nytimes.com/"}' 
	curl -X POST \
  	http://localhost:3000/url \
  	-H 'Content-Type: application/json' \
	-d '{ "url": "http://cbsnews.com" }' 
	curl localhost:3000/domains

test-top-5: test-setup
	curl localhost:3000/top

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

## Kubernetes stuff

version:
	kubectl version

cluster:
	kubectl cluster-info

config:
	kubectl config view

nodes:
	kubectl get nodes

list-pods:
	kubectl get pods

list-all-pods:
	kubectl get pods --all-namespaces

list-system-pods:
	kubectl get pods --namespace=kube-system

install-dashboard:
	kubectl create -f $(CURDIR)/infrastructure/kubernetes-dashboard.yaml

run-dashboard:
	kubectl port-forward $(pod) 8443:8443 --namespace=kube-system

start-api-proxy:
	kubectl proxy --port=8080 

list-deployments:
	kubectl get deployments

describe-pod:
	kubectl describe pod $(pod)

stack-up:
	docker stack deploy --compose-file docker-compose.yml topdomainsstack

stack-services:	
	docker stack services topdomainsstack

stack-down:
	docker stack rm topdomainsstack

create-namespace:
	kubectl create -f $(CURDIR)/infrastructure/kubernetes-namespace.yaml

stack-up-namespace:
	docker stack deploy --namespace topdomains --compose-file docker-compose.yml topdomainsstack

stack-down-namespace:
	docker stack rm --namespace topdomains topdomainsstack

kube-namespace-services:
	kubectl get services -n topdomains 

## Different way to deploy with Kubernetes


kompose-convert:
	kompose convert -f docker-compose.yml -o infrastructure/topdomains

deploy-kube:
	kubectl apply -f ./infrastructure/topdomains --namespace=topdomains

clean-kube:
	kubectl delete deploy/mongodb --namespace=topdomains
	kubectl delete services/mongodb services/topdomains --namespace=topdomains

get-all-kube:
	kubectl get all --namespace=topdomains