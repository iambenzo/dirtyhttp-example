package main

import (
	"errors"
	"strconv"
)

var newId = 0
var fakeDb = make(map[string]User)

type User struct {
	Id             string `json:"id,omitempty"`
	Email          string `json:"email"`
	AccountEnabled bool   `json:"accountEnabled"`
	DisplayName    string `json:"displayName"`
	GivenName      string `json:"givenName"`
	Surname        string `json:"surname"`
	CompanyName    string `json:"companyName"`
}

// Get a list of all users in our DB
func getAllUsers() []User {

	// Parse the values of our Map structure into a list
	out := make([]User, 0)
	for _, v := range fakeDb {
		out = append(out, v)
	}

	return out
}

// Returns a user if it's in the DB
func getUserById(id string) (User, bool) {
	// get the value if it's present
	v, ok := fakeDb[id]
	return v, ok
}

// Create a new user. Generates an ID and Password for the User
// before saving the User to our DB
func createUser(u User) User {
	// Generate an ID for the user
	u.Id = strconv.Itoa(newId)
	newId++

	fakeDb[u.Id] = u
	return u
}

// Update an existing user in the DB
func updateUser(id string, u *User) (*User, error) {
	// Check the user exists
	v, ok := fakeDb[id]

	if ok {
		v.Email = u.Email
		v.AccountEnabled = u.AccountEnabled
		v.DisplayName = u.DisplayName
		v.GivenName = u.GivenName
		v.Surname = u.Surname
		v.CompanyName = u.CompanyName

		fakeDb[id] = v

		return u, nil

	} else {
		return &User{}, errors.New("User doesn't exist")
	}

}

// Delete a user if it exists
// Returns true if a user was found and removes
// Returns false if the user wasn't found
func deleteUser(id string) bool {
	_, ok := fakeDb[id]
	if ok {
		delete(fakeDb, id)
		return true
	} else {
		return false
	}
}
