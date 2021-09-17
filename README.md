:::info

:::spoiler **Import domain**
---


**URL** : `/auth/domain/default/import`

**Method** : `POST`

**Auth required** : `Yes`

**Permissions required** : `Admin`

**Header**

| Key          | value             |
| ------------ | ----------------- |
| Authozation  | Bearer your_token |
| Content-Type | application.json  |


**Params**


```json
{
    "Authozation":"Bearer your_token",
    "Content-Type":"application.json"
}
```

**Body**

```json
{
	"domain":"domain.com",
	"name":"domain name",
	"domain_type":1,
	"publisher_id":1,
	"news_type": 1
}
```

> ### Note:
> + domain: must contain `.vn` or `.com`
> + name: lenght must in [5,200]
> + domain type: `1` trang chính, `2` chuyên trang, `3` PB ngước ngoài, `4` trang tin điện tử có phép, `5` trang tin điện tử không phép
> + publiser_id: id of publishser
> + news_type: `1` Báo & tạp chí online, `2` Trang tin điện tử


**Data examples**

Partial data is allowed.

```json
{
	"domain":"domain.com",
	"name":"domain name",
	"domain_type":1,
	"publisher_id":1,
	"news_type": 1
}
```

## Success Responses

**Condition** : Data provided is valid and User is Authenticated.

**Code** : `200 OK`

**Response**

```json
{
    "code": 0,
    "message": success,
    "body":{
        "id": 1234,
        "first_name": "Joe",
        "last_name": "Bloggs",
        "email": "joe25@example.com",
        "uapp": "ios1_2"
    }
}
```

## Error Response

**Condition** : If provided data is invalid, e.g. a name field is too long.

**Code** : `400 BAD REQUEST`

**Response** :

```json
{
   "code": 1,
    "message": fail,
    "body": "fail to create domain, case by........"
}
```
