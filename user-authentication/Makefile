setup:
	go mod init .

# Just runs Go API
run-server:
	go run src/user-authentication.go

# Runs Go API and MongoDB
run-all:
	docker-compose up --build

build: 
	go build -o bin/userauth src/user-authentication.go

run-binary: build
	bin/userauth

mongodb-run-local:
	docker run --name mongodb -p 27017:27017 -d mongo:3.7
