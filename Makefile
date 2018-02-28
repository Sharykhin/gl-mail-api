GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=gl-mail-api
BINARY_UNIX=$(BINARY_NAME)_unix


dev:
	APP_ENV=test JWT_PUBLIC_KEY=jwtRS256.key.pub GRPC_PUBLIC_KEY=server.crt GRPC_SERVER_ADDRESS=localhost:50051 go run main.go

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

lint:
	gometalinter ./...

test:
	echo "test"

