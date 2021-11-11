# dirtyhttp-example

The service will run on port 8080. This can be configured im `main.go:113`

If you send a GET request, the service will return all users in it's data structure. Including an `id` URL query parameter will return a single user with a matching ID.

```sh
curl -X GET http://localhost:8080
```

```sh
curl -X GET http://localhost:8080?id=0
```

You can create users via a POST request with the following body (obviously, play with the content):

```json
{
	"email": "dave@dave.com",
	"accountEnabled": true,
	"displayName": "Big Dave",
	"givenName": "Dave",
	"surname": "Davis",
	"companyName": "The Dave Ltd"
}
```

```sh
curl -X POST http://localhost:8080 \
-d '{"email": "dave@dave.com","accountEnabled": true,"displayName": "Big Dave","givenName": "Dave","surname": "Davis","companyName": "The Dave Ltd"}'
```

If you send a PUT request with an `id` URL parameter and the above request body, you will update the user mapped to that ID.

```sh
curl -X PUT http://localhost:8080?id=0 \
-d '{"email": "dave@dave.com","accountEnabled": true,"displayName": "Even Bigger Dave","givenName": "Dave","surname": "Davis","companyName": "The Dave Ltd"}'
```

You can also delete users by ID:

```sh
curl -X DELETE http://localhost:8080?id=0
```

