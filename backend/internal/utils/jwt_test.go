package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"testing"
	"time"
)

type unSupportedClaims struct {
	super string
	jwt.Claims
}

func TestGenerateJwt(t *testing.T) {
	// Use a fake token secret for testing
	cfg := struct {
		AuthTokenSecret    string
		RefreshTokenSecret string
	}{
		AuthTokenSecret:    "TEST_AUTH_SECRET",
		RefreshTokenSecret: "TEST_REFRESH_SECRET",
	}

	// Test for AuthTokenClaims
	authClaims := &AuthTokenClaims{
		UserID: uuid.New(),
	}
	authToken, err := GenerateJwt(authClaims, time.Now().Add(time.Hour), cfg.AuthTokenSecret)
	if err != nil {
		t.Errorf("Error generating auth JWT: %v", err)
	}
	if authToken == "" {
		t.Error("Generated auth JWT is empty")
	}

	// Test for RefreshTokenClaims
	refreshClaims := &RefreshTokenClaims{
		UserID:              uuid.New(),
		RefreshTokenVersion: 1,
	}
	refreshToken, err := GenerateJwt(refreshClaims, time.Now().Add(24*time.Hour), cfg.RefreshTokenSecret)
	if err != nil {
		t.Errorf("Error generating refresh JWT: %v", err)
	}
	if refreshToken == "" {
		t.Error("Generated refresh JWT is empty")
	}

	// Test for invalid signing secret
	_, err = GenerateJwt(authClaims, time.Now().Add(time.Hour), "")
	if err == nil {
		t.Error("Expected error for invalid signing secret")
	}

	// Test for unsupported claims type
	_, err = GenerateJwt(unSupportedClaims{
		super: "super",
	}, time.Now().Add(time.Hour), cfg.AuthTokenSecret)
	if err == nil {
		t.Error("Expected error for unsupported claims type")
	}
}

func TestDecodeJwt(t *testing.T) {
	// Use a fake token secret for testing

	var UserID = "00000000-0000-0000-0000-000000000000"

	var parsedUserID, _ = uuid.Parse(UserID)

	cfg := struct {
		AuthTokenSecret    string
		RefreshTokenSecret string
	}{
		AuthTokenSecret:    "8199aa73b295dd8bab866dec57eb699bd9ba96469785029985a1c19acd95d479",
		RefreshTokenSecret: "8199aa73b295dd8bab866dec57eb699bd9ba96469785029985a1c19acd95d479",
	}

	// Test for AuthTokenClaims
	authClaims := &AuthTokenClaims{
		UserID: parsedUserID,
	}

	authToken, err := GenerateJwt(authClaims, time.Now().Add(time.Hour), cfg.AuthTokenSecret)
	if err != nil {
		t.Errorf("Error generating auth JWT: %v", err)
	}

	decodedAuth, err := DecodeJwt(authToken, &AuthTokenClaims{}, true)
	if err != nil {
		t.Errorf("Error decoding auth JWT: %v", err)
	}

	if decodedAuth.UserID != authClaims.UserID {
		t.Errorf("Expected user ID %v, got %v", authClaims.UserID, decodedAuth.UserID)
	}

	// Test for RefreshTokenClaims
	refreshClaims := &RefreshTokenClaims{
		UserID:              parsedUserID,
		RefreshTokenVersion: 1,
	}

	refreshToken, err := GenerateJwt(refreshClaims, time.Now().Add(24*time.Hour), cfg.RefreshTokenSecret)
	if err != nil {
		t.Errorf("Error generating refresh JWT: %v", err)
	}

	decodedRefresh, err := DecodeJwt(refreshToken, &RefreshTokenClaims{}, false)
	if err != nil {
		t.Errorf("Error decoding refresh JWT: %v", err)
	}

	if decodedRefresh.UserID != refreshClaims.UserID {
		t.Errorf("Expected user ID %v, got %v", refreshClaims.UserID, decodedRefresh.UserID)
	}

	// Test for invalid signing secret
	_, err = DecodeJwt(authToken, &AuthTokenClaims{}, true)
	if err == nil {
		t.Error("Expected error for invalid signing secret")
	}

}
