package main

import (
	"context"
	"database/sql"
)

const baseQuery string = `SELECT id,
                            email,
                            enabled,
                            display_name,
                            given_name,
                            surname,
                            company_name,
                            password,
                            force_pass_change
                        FROM users`

type PasswordProfile struct {
	ForceChangePasswordNextSignIn bool   `json:"forceChangePasswordNextSignIn"`
	Password                      string `json:"password"`
}

type User struct {
	Id             string           `json:"id,omitempty"`
	Email          string           `json:"email"`
	AccountEnabled bool             `json:"accountEnabled"`
	DisplayName    string           `json:"displayName"`
	GivenName      string           `json:"givenName"`
	Surname        string           `json:"surname"`
	CompanyName    string           `json:"companyName"`
	PassProfile    *PasswordProfile `json:"passwordProfile,omitempty"`
}

func getAllUsers(ctx context.Context) []User {
	return make([]User, 0)
}

func getUserById(id string, ctx context.Context) *User {
	return &User{}
}

func getUserByEmail(email string, ctx context.Context) *User {
	return &User{}
}

func updateUser(u *User, ctx context.Context) *User {
	return u
}

func deleteUser(id string, ctx context.Context) bool {
	return true
}

func parseQueryResult(r *sql.Rows) ([]User, error) {
	return nil, nil
}
