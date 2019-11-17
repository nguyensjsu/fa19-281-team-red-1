setup:
	go mod init .

run-server:
	go run src/user-authentication.go

run:
	docker-compose up --build

build: 
	go build src/user-authentication.go

mongodb-run-local:
	docker run --name mongodb -p 27017:27017 -d mongo:3.7