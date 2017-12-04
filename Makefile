# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet
GOLINT=golint
GOFMT=gofmt
DEPCMD=dep
DEPENSURE=$(DEPCMD) ensure
BINARY_NAME=bin/gody
BINARY_UNIX=$(BINARY_NAME)_unix

all: test clean build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v
test:
		$(GOFMT) $(go list ./... | grep -v /vendor/)
		$(GOLINT) $(go list ./... | grep -v /vendor/)
		$(GOVET) $(go list ./... | grep -v /vendor/)
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		$(GOGET) -d -v .
		$(DEPENSURE) -update

## Cross compilation
#build-linux:
#		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
#docker-build:
#		docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v