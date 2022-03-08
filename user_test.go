package main

import (
	"testing"
)

var testUser User = User{
	Email:          "test@testr.com",
	AccountEnabled: true,
	DisplayName:    "Test Testing",
	GivenName:      "Test",
	Surname:        "Testing",
	CompanyName:    "TestCo",
}

func TestCreate(t *testing.T) {
	newUser := createUser(testUser)

	if newUser.Id == "" && newUser.Email != testUser.Email {
		t.Errorf("Create Failed. New User ID is %s.\n New User Email is %s.\n Expecting Email %s\n", newUser.Id, newUser.Email, testUser.Email)
	} else {
		testUser.Id = newUser.Id
	}
}

func TestUpdate(t *testing.T) {
	testUser.Email = "test@testing.com"

	if updatedUser, err := updateUser(testUser.Id, &testUser); err != nil {
		t.Errorf("%v", err)
	} else if updatedUser.Email != testUser.Email {
		t.Errorf("Updated User Email is %s.\n Was expecting %s", updatedUser.Email, testUser.Email)
	}
}

func TestGetById(t *testing.T) {
	if response, ok := getUserById(testUser.Id); !ok {
		t.Errorf("Failed to find user ID %s", testUser.Id)
	} else if response.Email != testUser.Email {
		t.Errorf("Return User email was %s. Was expecting %s", response.Email, testUser.Email)
	}
}

func TestGetAll(t *testing.T) {
	users := getAllUsers()

	if len(users) != len(fakeDb) {
		t.Errorf("0 Users returned")
	}

	// if users[0].Email != testUser.Email {
	// 	t.Errorf("First user email was %s. Expecting %s", users[0].Email, testUser.Email)
	// }
}

func TestDelete(t *testing.T) {
	currentDbSize := len(fakeDb)

	if ok := deleteUser(testUser.Id); !ok {
		t.Errorf("Couldn't find user %s", testUser.Id)
	}

	if len(fakeDb) == currentDbSize {
		t.Errorf("User not deleted. Current DB size: %d. Was Expecting %d", currentDbSize, currentDbSize-1)
	}
}
