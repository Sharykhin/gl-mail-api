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
GET /failed-mails?limit={number}&offset={number}
Status: 200 OK
```

**Filters**:
- limit - optional | number. How much rows to query
- offset - optional | number. How much rows to skip

**Response:**

Headers:
```bash
Content-Type application/json
```

 Body:
```json
{
    "success": true,
    "data": [
        {
            "id": 32,
            "action": "",
            "payload": {
                "to": "chapal@inbox.ru"
            },
            "reason": "test reason",
            "created_at": "Mon, 26 Feb 2018 16:33:23 UTC",
            "deleted_at": null
        },
        {
            "id": 33,
            "action": "test action",
            "payload": {},
            "reason": "test reason",
            "created_at": "Mon, 26 Feb 2018 16:42:23 UTC",
            "deleted_at": null
        },
        {
            "id": 34,
            "action": "test action",
            "payload": {
                "to": "chapal@inbox.ru"
            },
            "reason": "test reason",
            "created_at": "Mon, 26 Feb 2018 16:42:33 UTC",
            "deleted_at": null
        }
    ],
    "error": null,
    "meta": {
        "count": 3,
        "limit": 3,
        "offset": 5,
        "total": 10
    }
}
```
