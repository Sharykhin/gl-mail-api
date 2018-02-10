Mail API Service:
================

Api service that provide various information regarding mails.

Requirements:
-------------

[docker](https://www.docker.com/)

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

1. Build an image from Dockerfile:
```bash
docker build -t gl-mail-api .
```

2. Run a container:
```bash
docker run -p 8082:8082 gl-mail-api --rm gl-mail-api
```

Dockerfile exposes `8002` port.

3. Go to http://localhost:8082

API:
----

#### Ping:

Some kind of health check endpoint

```bash
GET /ping
Status: 200 OK
```

**Response:**

Headers:
```bash
Content-Type: text/plain; charset=utf-8
```

Body:
```bash
OK
```

#### Get Failed Mails:

Auth: Role `admin` requires, uses JWT

Returns a list of failed mails

```bash
GET /failed-mails
Status: 200 OK
```

**Response:**

Headers:
```bash
Content-Type application/json
```

 Body:
```json
{
    "success": true,
    "data": null,
    "error": null
}
```
