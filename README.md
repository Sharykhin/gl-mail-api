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

1. Build images:
```bash
docker-compose build
```

2. Run containers:
```bash
docker-compose up
```

3. Go to http://localhost:8082

*Examples:*  

Health-check
```bash
curl -XGET http://localhost:8002/ping
```

Create a new failed mail message
```bash
curl -XPOST -H "Content-Type: application/json" -d '{"action":"register", "payload":{"to":"unknown@mail.com"}, "reason":"no such mailbox"}' http://localhost:8002/failed-mails
```


[API Documentation](./doc/api.md)

