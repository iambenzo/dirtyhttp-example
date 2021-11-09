package main

import (
	"database/sql"
	"net/http"

	// "github.com/m1/go-generate-password/generator"
	"github.com/iambenzo/dirtyhttp"
)

var api dirtyhttp.Api = dirtyhttp.Api{}

// Handler/Controller struct
type helloHandler struct{}

// Implement http.Handler
//
// Your logic goes here
func (hey helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

        // get the URL query parameters
		var queryParameters = r.URL.Query()

		if queryParameters.Get("id") != "" {
            dirtyhttp.EncodeResponseAsJSON(getUserById(queryParameters.Get("id"), r.Context()), w)
            return
		} else if queryParameters.Get("email") != "" {
            dirtyhttp.EncodeResponseAsJSON(getUserByEmail(queryParameters.Get("email"), r.Context()), w)
            return
		} else {
            dirtyhttp.EncodeResponseAsJSON(getAllUsers(r.Context()), w)
            return
		}

	default:
		// Write a timestamped log entry
		api.Logger.Error("A non-implemented method was attempted")

		// Return a pre-defined error with a custom message
		api.HttpErrorWriter.MethodNotAllowed(w, "Naughty, naughty.")
		return
	}
}

func getDbConnection(cnf *dirtyhttp.EnvConfig) *sql.DB {
	var db *sql.DB
	var err error

	db, err = sql.Open("sqlite3", "./data.sqlite")

	if err != nil {
		api.Logger.Fatal("Couldn't connect to database")
	}

	return db
}

func main() {
	// Initialisation
	api.Init()

	// Register a handler
	hello := &helloHandler{}
	api.RegisterHandler("/", *hello)

	// Go, baby, go!
	api.StartService()
}
