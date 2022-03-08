package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var testUserHttp User = User{
	Email:          "test@testh.com",
	AccountEnabled: true,
	DisplayName:    "Test HttpTesting",
	GivenName:      "Test",
	Surname:        "HttpTesting",
	CompanyName:    "TestCo",
}

func TestMain(m *testing.M) {
	go main()
	os.Exit(m.Run())
}

func HttpDo(req *http.Request) ([]byte, error) {
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()

	out, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}

func TestCreateHttp(t *testing.T) {
	newUser, err := json.Marshal(testUserHttp)
	if err != nil {
		t.Fatalf("Failed to parse user into bytes")
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/", bytes.NewBuffer(newUser))
	if err != nil {
		t.Fatalf("Error in creating request: %v", err)
	}

	out, err := HttpDo(req)
	if err != nil {
		t.Fatalf("There was a problem with parsing the upstream response %s\n", err.Error())
	}

	var body User
	json.Unmarshal(out, &body)

	if body.Id == "" && body.Email != testUserHttp.Email {
		t.Errorf("Create Failed. New User ID is %s.\n New User Email is %s.\n Expecting Email %s\n", body.Id, body.Email, testUserHttp.Email)
	} else {
		testUserHttp.Id = body.Id
	}

}

func TestUpdateHttp(t *testing.T) {
	testUserHttp.Email = "test@testinghttp.com"

	updatedUser, err := json.Marshal(testUserHttp)
	if err != nil {
		t.Fatalf("Failed to parse user into bytes")
	}

	url := fmt.Sprintf("http://localhost:8080?id=%s", testUserHttp.Id)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(updatedUser))
	if err != nil {
		t.Fatalf("Error in creating request: %v", err)
	}

	out, err := HttpDo(req)
	if err != nil {
		t.Fatalf("There was a problem with parsing the upstream response %s\n", err.Error())
	}

	var body User
	json.Unmarshal(out, &body)

	if body.Email != testUserHttp.Email {
		fmt.Printf("%v\n", fakeDb)
		t.Errorf("Updated User Email is %s.\n Was expecting %s", body.Email, testUserHttp.Email)
	}

}

func TestGetByIdHttp(t *testing.T) {
	url := fmt.Sprintf("http://localhost:8080?id=%s", testUserHttp.Id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatalf("Error in creating request: %v", err)
	}

	out, err := HttpDo(req)
	if err != nil {
		t.Fatalf("There was a problem with parsing the upstream response %s\n", err.Error())
	}

	var body User
	json.Unmarshal(out, &body)

	if body.Email != testUserHttp.Email {
		fmt.Printf("%v\n", fakeDb)
		t.Errorf("Updated User Email is %s.\n Was expecting %s", body.Email, testUserHttp.Email)
	}
}

func TestGetAllHttp(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	if err != nil {
		t.Fatalf("Error in creating request: %v", err)
	}

	out, err := HttpDo(req)
	if err != nil {
		t.Fatalf("There was a problem with parsing the upstream response %s\n", err.Error())
	}

	var body []User
	json.Unmarshal(out, &body)

	if len(fakeDb) != len(body) {
		t.Errorf("Listing all users failed. Expecting %d, got %d", len(fakeDb), len(body))
	}

}

func TestDeleteHttp(t *testing.T) {
	currentDbSize := len(fakeDb)

	url := fmt.Sprintf("http://localhost:8080?id=%s", testUserHttp.Id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		t.Fatalf("Error in creating request: %v", err)
	}

	out, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not send delete request to API")
	}

	if out.StatusCode != http.StatusNoContent {
		t.Errorf("Bad status code. Expecting %d. Got %d", http.StatusNoContent, out.StatusCode)
	}

	if len(fakeDb) == currentDbSize {
		t.Errorf("User not deleted. Current DB size: %d. Was Expecting %d", currentDbSize, currentDbSize-1)
	}
}
