# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet
GOLINT=golint
GOX=gox
DEPCMD=dep
DEPENSURE=$(DEPCMD) ensure
BINARY_NAME=gody

all: test clean build
build:
		$(GOX) --osarch "darwin/amd64 linux/amd64 windows/amd64" -output="bin/{{.OS}}_{{.Arch}}/$(BINARY_NAME)"
install:
		$(GOINSTALL)
test:
		find . -type f -name "*.go" | grep -v "^./vendor" | xargs gofmt -d -e -s -w -l
		$(GOLINT) $(go list ./... | grep -v /vendor/)
		$(GOVET) $(go list ./... | grep -v /vendor/)
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -rf bin/
#run:
#		$(GOBUILD) -o $(BINARY_NAME) -v ./...
#		./$(BINARY_NAME)
deps:
		$(GOGET) -d -v .
		$(DEPENSURE) -update

## Cross compilation
#build-linux:
#		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
#docker-build:
#		docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v