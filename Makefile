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

lint:
	gometalinter ./...

test:
	pwd

docker-dev:
	docker run --rm --name gl-mail-api-go -v "$(pwd)":/go/src/github.com/Sharykhin/gl-mail-api -p 8002:8002 -d golang:1.9 tail -f /dev/null
	docker run --name gl-mail-api-mysql -v "$(pwd)"/.docker-runtime/mysqldb:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
	docker exec gl-mail-api-mysql mysql -uroot -proot -e "create database if not exists test character set utf8 collate utf8_general_ci;"
	docker exec gl-mail-api-mysql mysql -uroot -proot -e "use test; create table if not exists failed_messages(id int auto_increment primary key, action varchar(80) not null, payload json not null, reason text, created_at timestamp);"

docker-dev-stop:
	docker stop gl-mail-api-go
	docker stop gl-mail-api-mysql
