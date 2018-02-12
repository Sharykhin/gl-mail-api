GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=gl-mail-api
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	source .env
	go run main.go

lint:
	gometalinter ./...

docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w  /go/src/github.com/Sharykhin/gl-mail-api golang:1.9 go build -o "$(BINARY_UNIX)" -v