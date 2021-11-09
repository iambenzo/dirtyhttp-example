package main

import (
	"errors"
	"strconv"

	"github.com/m1/go-generate-password/generator"
)

var newId = 0
var fakeDb = make(map[string]User)

type PasswordProfile struct {
	ForceChangePasswordNextSignIn bool   `json:"forceChangePasswordNextSignIn"`
	Password                      string `json:"password"`
}

type User struct {
	Id             string          `json:"id,omitempty"`
	Email          string          `json:"email"`
	AccountEnabled bool            `json:"accountEnabled"`
	DisplayName    string          `json:"displayName"`
	GivenName      string          `json:"givenName"`
	Surname        string          `json:"surname"`
	CompanyName    string          `json:"companyName"`
	PassProfile    PasswordProfile `json:"-"`
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

func createUser(u User) (User, error) {
	// Generate an ID for the user
	u.Id = strconv.Itoa(newId)
	newId++

	// Generate a password for them
	passwordGeneratorConfig := generator.Config{
		Length:                     16,
		IncludeSymbols:             false,
		IncludeNumbers:             true,
		IncludeLowercaseLetters:    true,
		IncludeUppercaseLetters:    true,
		ExcludeSimilarCharacters:   true,
		ExcludeAmbiguousCharacters: true,
	}

	g, err := generator.New(&passwordGeneratorConfig)
	if err != nil {
		return User{}, errors.New("Could not configure password generator")
	}
	pass, err := g.Generate()
	if err != nil {
		return User{}, errors.New("Could not generate a password")
	}
	newPass := PasswordProfile{
		ForceChangePasswordNextSignIn: true,
		Password:                      *pass,
	}
	u.PassProfile = newPass

	fakeDb[u.Id] = u
    return u, nil
}

func updateUser(u *User) *User {
	return u
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
