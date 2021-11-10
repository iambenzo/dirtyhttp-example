The service will run on port 8080 - You can change this by setting the API_PORT environment variable to your desired port.

To use the sample application, you can use the `run.sh` script from the git repository. It sets two environment variables: API_USER and API_PASSWORD. These are used to configure basic authentication for communicating with the service.

If you send a GET request, the service will return all users in it's data structure. Including an `id` URL query parameter will return a single user with a matching ID.

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

If you send a PUT request with an `id` URL parameter and the above request body, you will update the user mapped to that ID.
