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

2. Get mysql container (skip if mysql container already exists):
```bash
docker pull mysql
```

3. Run mysql:
```bash
docker run --name mysqldb -v /my/own/datadir/.docker-runtime/mysqldb:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
```

4. Run golang container. Use appropriate env variables:
```bash
docker run --env PUBLIC_KEY=jwtRS256.key.pub MYSQL_SOURCE="root:root@tcp(172.17.0.2:3306)/test-db" -p 8002:8002 --name gl-mail-api-service --rm gl-mail-api
```

Dockerfile exposes `8002` port.

5. Go to http://localhost:8082

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

#### Create a new failed message:

Auth: *will be*

Create a new failed mail row
```bash
POST /failed-mails
Status: 201 Created
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
    "data": {
      "id": 1,
      "action": "register",
      "payload": {
        "to": "test@test.com"
      },
      "reason": "could not sent mail to non-existent mailbox"
    },
    "error": null
}
```
