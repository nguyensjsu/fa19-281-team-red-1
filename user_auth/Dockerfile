FROM golang:latest 

# Wait for db connection
RUN curl https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh > /wait-for-it.sh && chmod +x /wait-for-it.sh

EXPOSE 3000
RUN mkdir /app 
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . .
# ENV GOPATH /app
# Build the Go app
RUN cd /app ; go build -o /app/bin/userauth src/user-authentication.go
# RUN cd /app ; go install top_domains
CMD ["/app/bin/userauth"]
