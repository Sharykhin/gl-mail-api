API Documentation:
------------------

#### Ping:

*Auth*: No

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

**Request**

Headers:
```bash
Content-Type application/json
```

JSON-in payload:
```json
{
	"action":"register",
	"payload": {
		"to":"chapal@inbox.ru"
	},
	"reason": "failed to sent mail. Mailgun errorL 41245, box is invalid."
}
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
            "to": "chapal@inbox.ru"
        },
        "reason": "failed to sent mail. Mailgun errorL 41245, box is invalid.",
        "created_at": "2018-02-13T09:18:43.576916898+03:00"
    },
    "error": null
}
```
