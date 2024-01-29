package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"server04/controller"
	"server04/ent"
)

func TestUserCreateController(t *testing.T) {
	// Create a new user object
	newUser := &ent.UserSec{
		Email:    "test@example.com",
		Password: "password",
		// Add other fields as needed
	}

	// Marshal the user object to JSON
	userJSON, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the method "POST" and the URL "/user"
	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Set the content type of the request to "application/json"
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UserCreateController)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"success":true,"message":"","data":{}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestUserLoginController(t *testing.T) {
	// Create a new user object
	newUser := &ent.UserSec{
		Email:    "test@example.com",
		Password: "password",
		// Add other fields as needed
	}

	// Marshal the user object to JSON
	userJSON, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the method "POST" and the URL "/login"
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Set the content type of the request to "application/json"
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UserLoginController)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"success":true,"message":"","data":{}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUserValidationSendController(t *testing.T) {
	// Create a new email object
	emailToValidate := struct {
		Email string `json:"email,omitempty"`
	}{
		Email: "test@example.com",
	}

	// Marshal the email object to JSON
	emailJSON, err := json.Marshal(emailToValidate)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the method "POST" and the URL "/validation/send"
	req, err := http.NewRequest("POST", "/validation/send", bytes.NewBuffer(emailJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Set the content type of the request to "application/json"
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UserValidationSendController)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"success":true,"message":"","data":{"email":"test@example.com"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
