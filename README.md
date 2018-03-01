Mail API Service:
================

Api service that provide various information regarding mails.

Requirements:
-------------

[docker](https://www.docker.com/)

Technologies:
-------------

[golang 1.9](https://golang.org/)  
[mysql 5.7](https://dev.mysql.com)  

Testing:
-------
Generate key-pair:
```bash
ssh-keygen -t rsa -b 4096 -f jwtRS256.key
# Don't add passphrase
openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub
```

Go to [jwt.io](https://jwt.io/)  
Select RS256 ALG and put private and public keys. Increase exp if necessary and put `role` with `admin` value into a payload.
Copy token and paste it into a request.

Usage:
------

1. Go to .docker/golang and .docker/mysql and copy .env.example to .env
```bash
cd .docker/golang && cp .env.example .env
```

2. Build images:
```bash
docker-compose build
```

4. Run containers:
```bash
docker-compose up
```

5. Go to http://localhost:8082

For local env:
```bash
JWT_PUBLIC_KEY=jwtRS256.key.pub GRPC_PUBLIC_KEY=server.crt GRPC_SERVER_ADDRESS=localhost:50051 go run main.go
```

or use Makefile:
```bash
make dev
```

*Examples:*  

Health-check
```bash
curl -XGET http://localhost:8002/ping
```

Create a new failed mail message
```bash
curl -XGET -H "Content-Type: application/json" http://localhost:8002/failed-mails?limit=10&offset=5
```


[API Documentation](./doc/api.md)

