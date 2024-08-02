package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type AuthTokenClaims struct {
	UserID uuid.UUID
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	UserID              uuid.UUID
	RefreshTokenVersion int32
	jwt.RegisteredClaims
}

type JWTUser struct {
	UserID              uuid.UUID
	RefreshTokenVersion int32
}

type ClaimsInterface interface {
	jwt.Claims
}

func GenerateJwt(claims ClaimsInterface, exp time.Time, signingSecret string) (string, error) {
	switch c := claims.(type) {
	case *AuthTokenClaims:
		c.UserID = claims.(*AuthTokenClaims).UserID
		c.ExpiresAt = &jwt.NumericDate{Time: exp}
		c.IssuedAt = &jwt.NumericDate{Time: time.Now()}
	case *RefreshTokenClaims:
		c.UserID = claims.(*RefreshTokenClaims).UserID
		c.RefreshTokenVersion = claims.(*RefreshTokenClaims).RefreshTokenVersion
		c.ExpiresAt = &jwt.NumericDate{Time: exp}
		c.IssuedAt = &jwt.NumericDate{Time: time.Now()}
	default:
		return "", fmt.Errorf("unsupported claims type")
	}

	if signingSecret == "" {
		return "", fmt.Errorf("invalid signing secret")
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := unsignedToken.SignedString([]byte(signingSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func DecodeJwt(token string, claims ClaimsInterface, signingSecret string) (JWTUser, error) {
	decoded, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(signingSecret), nil
	})

	if err != nil {
		return JWTUser{}, err
	}

	switch claims.(type) {
	case *AuthTokenClaims:
		claims, ok := decoded.Claims.(*AuthTokenClaims)
		if !ok {
			return JWTUser{}, fmt.Errorf("invalid claims type")
		}
		return JWTUser{
			UserID: claims.UserID,
		}, nil
	case *RefreshTokenClaims:
		claims, ok := decoded.Claims.(*RefreshTokenClaims)
		if !ok {
			return JWTUser{}, fmt.Errorf("invalid claims type")
		}
		return JWTUser{
			UserID:              claims.UserID,
			RefreshTokenVersion: claims.RefreshTokenVersion,
		}, nil
	default:
		return JWTUser{}, fmt.Errorf("unsupported claims type")
	}

}

func ValidateJwt(exp *time.Time) (bool, error) {
	if time.Now().After(*exp) {
		return false, fmt.Errorf("token expired")
	}

	return true, nil
}
