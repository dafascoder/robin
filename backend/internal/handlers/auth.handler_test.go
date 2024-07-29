package handlers

import (
	"backend/internal/forms"
	"backend/internal/services"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// create test for this function
func TestAuthHandler_HandleSignUp(t *testing.T) {
	// Mock AuthServices
	mockAuthServices := &services.AuthServices{}

	// Create AuthHandler with mock services
	authHandler := &AuthHandler{
		AuthServices: mockAuthServices,
	}

	// Create a mock request
	signUpForm := forms.SignUpForm{
		Email:    "test@example.com",
		Password: "password",
	}
	body, _ := json.Marshal(signUpForm)
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the HandleSignUp method
	authHandler.HandleSignUp(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `{"message":"User registered successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
