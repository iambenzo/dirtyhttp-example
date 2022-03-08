package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iambenzo/dirtyhttp"
)

var api dirtyhttp.Api = dirtyhttp.Api{}

// Handler/Controller struct
type httpHandler struct{}

// Implement http.Handler
//
// Logic goes here
func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// get the URL query parameters
		var queryParameters = r.URL.Query()

		if queryParameters.Get("id") != "" {
			v, ok := getUserById(queryParameters.Get("id"))
			if ok {
				dirtyhttp.EncodeResponseAsJSON(v, w)
				return
			} else {
				api.HttpErrorWriter.BadRequest(w, "User does not exist")
				return
			}
		} else {
			dirtyhttp.EncodeResponseAsJSON(getAllUsers(), w)
			return
		}

	case http.MethodPut:
		// get the URL query parameters
		var queryParameters = r.URL.Query()

		if queryParameters.Get("id") != "" {
			// Get user object from request body
			d := json.NewDecoder(r.Body)
			var user User
			err := d.Decode(&user)
			if err != nil {
				api.Logger.Error(fmt.Sprintf("%v", err))
				api.HttpErrorWriter.InternalServerError(w, "Unable to parse request body")
				return
			}

			// Update our DB
			u, err := updateUser(queryParameters.Get("id"), &user)
			if err != nil {
				api.HttpErrorWriter.InternalServerError(w, "User doesn't exist")
			}

			dirtyhttp.EncodeResponseAsJSON(u, w)
			return

		} else {
			api.HttpErrorWriter.BadParameters(w, "id")
			return
		}

	case http.MethodDelete:
		// get the URL query parameters
		var queryParameters = r.URL.Query()

		if queryParameters.Get("id") != "" {
			if deleteUser(queryParameters.Get("id")) {
				w.WriteHeader(http.StatusNoContent)
				return
			} else {
				api.HttpErrorWriter.BadRequest(w, "User does not exist")
				return
			}
		} else {
			api.HttpErrorWriter.WriteError(w, http.StatusBadRequest, "Please include an 'id' parameter")
			return
		}

	case http.MethodPost:
		// Get user object from request body
		d := json.NewDecoder(r.Body)
		var user User
		err := d.Decode(&user)
		if err != nil {
			api.Logger.Error(fmt.Sprintf("%v", err))
			api.HttpErrorWriter.InternalServerError(w, "Unable to parse request body")
			return
		}

		// save the data
		dirtyhttp.EncodeResponseAsJSON(createUser(user), w)

	default:
		// Write a timestamped log entry
		api.Logger.Error("A non-implemented method was attempted")

		// Return a pre-defined error with a custom message
		api.HttpErrorWriter.MethodNotAllowed(w, "Naughty, naughty.")
		return
	}
}

func main() {
	// Initialisation
	// Use custom config to remove auth
	config := dirtyhttp.EnvConfig{}
	config.ApiPort = "8080" // change port here
	api.InitWithConfig(&config)

	// Register a handler
	h := &httpHandler{}
	api.RegisterHandler("/", *h)

	// Go, baby, go!
	api.StartServiceNoAuth()
}
