Mail API Service:
================

Api service provides information regarding failed mails.

Requirements:
-------------

[docker](https://www.docker.com/)

Technologies:
-------------

[golang 1.9](https://golang.org/)

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

For grpc copy server.crt file from grpc server.

Attention:
----------
By default app environment is set to **dev** and communication
with grpc server goes with insecure way. To switch to secure
change *APP_ENV* to **prod**

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

Get a list of failed mails
```bash
curl -XGET -H "Content-Type: application/json" http://localhost:8002/failed-mails?limit=10&offset=5
```


[API Documentation](./doc/api.md)

